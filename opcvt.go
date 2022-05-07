package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
)

func cvtInt(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtInt8(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int8
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtInt16(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int16
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtInt32(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int32
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtInt64(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = int64
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUint(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uint
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUint8(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uint8
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUint16(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uint16
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUint32(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uint32
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUint64(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uint64
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtUintptr(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = uintptr
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtFloat32(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = float32
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}

func cvtFloat64(ir, ix register, xkind reflect.Kind, xtyp reflect.Type, typ reflect.Type) func(fr *frame) {
	type T = float64
	t := basic.TypeOfType(typ)
	isBasic := typ.PkgPath() == ""
	if xtyp.PkgPath() == "" {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(fr.reg(ix).(int))
			case reflect.Int8:
				v = T(fr.reg(ix).(int8))
			case reflect.Int16:
				v = T(fr.reg(ix).(int16))
			case reflect.Int32:
				v = T(fr.reg(ix).(int32))
			case reflect.Int64:
				v = T(fr.reg(ix).(int64))
			case reflect.Uint:
				v = T(fr.reg(ix).(uint))
			case reflect.Uint8:
				v = T(fr.reg(ix).(uint8))
			case reflect.Uint16:
				v = T(fr.reg(ix).(uint16))
			case reflect.Uint32:
				v = T(fr.reg(ix).(uint32))
			case reflect.Uint64:
				v = T(fr.reg(ix).(uint64))
			case reflect.Uintptr:
				v = T(fr.reg(ix).(uintptr))
			case reflect.Float32:
				v = T(fr.reg(ix).(float32))
			case reflect.Float64:
				v = T(fr.reg(ix).(float64))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	} else {
		return func(fr *frame) {
			var v T
			switch xkind {
			case reflect.Int:
				v = T(basic.Int(fr.reg(ix)))
			case reflect.Int8:
				v = T(basic.Int8(fr.reg(ix)))
			case reflect.Int16:
				v = T(basic.Int16(fr.reg(ix)))
			case reflect.Int32:
				v = T(basic.Int32(fr.reg(ix)))
			case reflect.Int64:
				v = T(basic.Int64(fr.reg(ix)))
			case reflect.Uint:
				v = T(basic.Uint(fr.reg(ix)))
			case reflect.Uint8:
				v = T(basic.Uint8(fr.reg(ix)))
			case reflect.Uint16:
				v = T(basic.Uint16(fr.reg(ix)))
			case reflect.Uint32:
				v = T(basic.Uint32(fr.reg(ix)))
			case reflect.Uint64:
				v = T(basic.Uint64(fr.reg(ix)))
			case reflect.Uintptr:
				v = T(basic.Uintptr(fr.reg(ix)))
			case reflect.Float32:
				v = T(basic.Float32(fr.reg(ix)))
			case reflect.Float64:
				v = T(basic.Float64(fr.reg(ix)))
			}
			if isBasic {
				fr.setReg(ir, v)
			} else {
				fr.setReg(ir, basic.Make(t, v))
			}
		}
	}
}
