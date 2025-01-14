package gilder_data

import (
	"bytes"

	"realy.lol/context"
	"realy.lol/lol"
)

type (
	bo  = bool
	by  = []byte
	st  = string
	er  = error
	no  = int
	i8  = int8
	u8  = byte
	i16 = int16
	u16 = uint16
	i32 = int32
	u32 = uint32
	i64 = int64
	u64 = uint64
	f32 = float32
	f64 = float64
	cx  = context.T
)

var (
	log, chk, errorf = lol.Main.Log, lol.Main.Check, lol.Main.Errorf
	equals           = bytes.Equal
)
