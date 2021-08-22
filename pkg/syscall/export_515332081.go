// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_515332081, typList_515332081)
}

var extMap_515332081 = map[string]interface{}{
	"syscall.Getdtablesize": syscall.Getdtablesize,
	"syscall.Undelete":      syscall.Undelete,
}

var typList_515332081 = []interface{}{
	(*syscall.IfmaMsghdr)(nil),
	(*syscall.InterfaceMulticastAddrMessage)(nil),
}