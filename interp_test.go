// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp_test

// This test runs the SSA interpreter over sample Go programs.
// Because the interpreter requires intrinsics for assembly
// functions and many low-level runtime routines, it is inherently
// not robust to evolutionary change in the standard library.
// Therefore the test cases are restricted to programs that
// use a fake standard library in testdata/src containing a tiny
// subset of simple functions useful for writing assertions.
//
// We no longer attempt to interpret any real standard packages such as
// fmt or testing, as it proved too fragile.

import (
	"bytes"
	"fmt"
	"go/build"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/goplus/interp"
	_ "github.com/goplus/interp/pkg"
	// _ "github.com/goplus/interp/pkg/errors"
	// _ "github.com/goplus/interp/pkg/fmt"
	// _ "github.com/goplus/interp/pkg/math"
	// _ "github.com/goplus/interp/pkg/os"
	// _ "github.com/goplus/interp/pkg/reflect"
	// _ "github.com/goplus/interp/pkg/runtime"
	// _ "github.com/goplus/interp/pkg/strings"
	// _ "github.com/goplus/interp/pkg/sync"
	// _ "github.com/goplus/interp/pkg/time"
)

// Each line contains a space-separated list of $GOROOT/test/
// filenames comprising the main package of a program.
// They are ordered quickest-first, roughly.
//
// If a test in this list fails spuriously, remove it.
var gorootTestTests = []string{
	"235.go",
	"alias1.go",
	"func5.go",
	"func6.go",
	"func7.go",
	"func8.go",
	"helloworld.go",
	"varinit.go",
	"escape3.go",
	"initcomma.go",
	"cmp.go", // import OS
	"compos.go",
	"turing.go",
	"indirect.go",
	"complit.go",
	"for.go",
	"struct0.go",
	"intcvt.go",
	"printbig.go",
	"deferprint.go",
	"escape.go",
	"range.go",
	"const4.go",
	"float_lit.go",
	"bigalg.go",
	"decl.go",
	"if.go",
	"named.go",
	"bigmap.go",
	"func.go",
	"reorder2.go",
	"gc.go", // import runtime
	"simassign.go",
	"iota.go",
	"nilptr2.go",
	"utf.go", // import unicode/utf8
	"method.go",
	"char_lit.go", // import os
	//"env.go",         // import runtime;os
	"int_lit.go",     //import os
	"string_lit.go",  //import os
	"defer.go",       //import fmt
	"typeswitch.go",  //import os
	"stringrange.go", //import os fmt unicode/utf8
	"reorder.go",     //import fmt
	"method3.go",
	"literal.go",
	"nul1.go", // doesn't actually assert anything (errorcheckoutput)
	"zerodivide.go",
	"convert.go",
	"convT2X.go",
	"switch.go",
	"ddd.go",
	"blank.go",      // partly disabled //import os
	"closedchan.go", //import os
	"divide.go",     //import fmt
	"rename.go",     //import runtime fmt
	"nil.go",
	"recover1.go",
	"recover2.go",
	//"recover3.go", //TODO fix panic info
	"typeswitch1.go",
	"floatcmp.go",
	"crlf.go", // doesn't actually assert anything (runoutput)
	"append.go",
	"chancap.go",
	"const.go",
	"deferfin.go",
}

// These are files in go.tools/go/ssa/interp/testdata/.
var testdataTests = []string{
	"boundmeth.go",
	"complit.go",
	"coverage.go",
	"defer.go",
	"fieldprom.go",
	"ifaceconv.go",
	"ifaceprom.go",
	"initorder.go",
	"methprom.go",
	"mrvchain.go",
	"range.go",
	"recover.go",
	"reflect.go",
	"static.go",
	"recover2.go",
}

var (
	gorootTestSkips = make(map[string]string)
)

func init() {
	if runtime.GOARCH == "386" {
		interp.UnsafeSizes = &types.StdSizes{WordSize: 4, MaxAlign: 4}
		gorootTestSkips["printbig.go"] = "load failed"
		gorootTestSkips["peano.go"] = "stack overflow"
	}
	if runtime.GOOS == "windows" {
		gorootTestSkips["env.go"] = "skip GOARCH"
	}
	gorootTestSkips["closure.go"] = "runtime.ReadMemStats"
	gorootTestSkips["divmod.go"] = "timeout"
	gorootTestSkips["copy.go"] = "slow"
	gorootTestSkips["gcstring.go"] = "timeout"
	gorootTestSkips["finprofiled.go"] = "slow"
	gorootTestSkips["gcgort.go"] = "slow"
	gorootTestSkips["inline_literal.go"] = "TODO, runtime.FuncForPC"
	gorootTestSkips["nilptr.go"] = "skip drawin"
	gorootTestSkips["recover.go"] = "TODO, fix test16"
	gorootTestSkips["heapsampling.go"] = "runtime.MemProfileRecord"
	gorootTestSkips["makeslice.go"] = "TODO, panic info, allocation size out of range"
	gorootTestSkips["stackobj.go"] = "skip gc"
	gorootTestSkips["stackobj3.go"] = "skip gc"
	gorootTestSkips["nilptr_aix.go"] = "slow"
	gorootTestSkips["init1.go"] = "skip gc"
	gorootTestSkips["string_lit.go"] = "call to os.Exit(0) during test"
	gorootTestSkips["switch.go"] = "call to os.Exit(0) during test"
	gorootTestSkips["ken/divconst.go"] = "slow"
	gorootTestSkips["ken/modconst.go"] = "slow"
	gorootTestSkips["fixedbugs/bug347.go"] = "TODO: runtime.Caller"
	gorootTestSkips["fixedbugs/bug348.go"] = "TODO: runtime.Caller"
}

var (
	igop string
)

func init() {
	var err error
	igop, err = exec.LookPath("igop")
	if err != nil {
		panic(fmt.Sprintf("not found igop: %v", err))
	}
}

func runInterp(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	err := interp.Run(0, input, nil)
	sec := time.Since(start).Seconds()
	if err != nil {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

func runIgop(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	cmd := exec.Command(igop, "run", input)
	data, err := cmd.CombinedOutput()
	if len(data) > 0 {
		fmt.Println(string(data))
	}
	sec := time.Since(start).Seconds()
	if err != nil || bytes.Contains(data, []byte("BUG")) {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

func printFailures(failures []string) {
	if failures != nil {
		fmt.Println("The following tests failed:")
		for _, f := range failures {
			fmt.Printf("\t%s\n", f)
		}
	}
}

// TestTestdataFiles runs the interpreter on testdata/*.go.
func TestTestdataFiles(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var failures []string
	for _, input := range testdataTests {
		if !runInterp(t, filepath.Join(cwd, "testdata", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

// $GOROOT/test/*.go runs
func getGorootTestRuns(t *testing.T) (dir string, files []string) {
	dir = filepath.Join(build.Default.GOROOT, "test")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return nil
			}
			_, n := filepath.Split(path)
			switch n {
			case "abi":
				return filepath.SkipDir
			case "fixedbugs":
				return filepath.SkipDir
			case "bench", "dwarf", "codegen":
				return filepath.SkipDir
			}
			return nil
			return filepath.SkipDir
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("read %v error: %v", path, err)
			return nil
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 {
			line := strings.TrimSpace(lines[0])
			if line == "// run" {
				files = append(files, path)
			}
		}
		return nil
	})
	return
}

// TestGorootTest runs the interpreter on $GOROOT/test/*.go.
func TestGorootTest(t *testing.T) {
	dir, files := getGorootTestRuns(t)
	var failures []string

	for _, input := range files {
		f := input[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input, info)
			continue
		}
		if !runIgop(t, input) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}
