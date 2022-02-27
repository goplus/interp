// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package filepath

import (
	q "path/filepath"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "filepath",
		Path: "path/filepath",
		Deps: map[string]string{
			"errors":       "errors",
			"os":           "os",
			"runtime":      "runtime",
			"sort":         "sort",
			"strings":      "strings",
			"syscall":      "syscall",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"WalkFunc": {reflect.TypeOf((*q.WalkFunc)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrBadPattern": reflect.ValueOf(&q.ErrBadPattern),
			"SkipDir":       reflect.ValueOf(&q.SkipDir),
		},
		Funcs: map[string]reflect.Value{
			"Abs":          reflect.ValueOf(q.Abs),
			"Base":         reflect.ValueOf(q.Base),
			"Clean":        reflect.ValueOf(q.Clean),
			"Dir":          reflect.ValueOf(q.Dir),
			"EvalSymlinks": reflect.ValueOf(q.EvalSymlinks),
			"Ext":          reflect.ValueOf(q.Ext),
			"FromSlash":    reflect.ValueOf(q.FromSlash),
			"Glob":         reflect.ValueOf(q.Glob),
			"HasPrefix":    reflect.ValueOf(q.HasPrefix),
			"IsAbs":        reflect.ValueOf(q.IsAbs),
			"Join":         reflect.ValueOf(q.Join),
			"Match":        reflect.ValueOf(q.Match),
			"Rel":          reflect.ValueOf(q.Rel),
			"Split":        reflect.ValueOf(q.Split),
			"SplitList":    reflect.ValueOf(q.SplitList),
			"ToSlash":      reflect.ValueOf(q.ToSlash),
			"VolumeName":   reflect.ValueOf(q.VolumeName),
			"Walk":         reflect.ValueOf(q.Walk),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"ListSeparator": {"untyped rune", constant.MakeInt64(int64(q.ListSeparator))},
			"Separator":     {"untyped rune", constant.MakeInt64(int64(q.Separator))},
		},
	})
}