package sharedutil

import (
	"encoding/binary"
	"math/rand/v2"
	"path/filepath"
	"time"
)

var bigEndian = &binary.BigEndian

// fast random name generate
func RndStr() string {
	const hextable = "0123456789abcdef"

	buff12 := [12]byte{}
	{
		now := uint64(time.Now().UnixNano())

		rndNumb := uint16(rand.Uint32())

		bigEndian.PutUint64(buff12[:], now)

		var hashSum uint16

		for i := 0; i < 8; i++ {
			hashSum += uint16(buff12[i])
		}

		bigEndian.PutUint16(buff12[8:], hashSum)

		bigEndian.PutUint16(buff12[10:], rndNumb)
	}

	buff24 := [24]byte{}
	{
		var j int

		for _, v := range buff12 {
			buff24[j] = hextable[v>>4]
			buff24[j+1] = hextable[v&0x0f]
			j += 2
		}
	}

	return string(buff24[:])
}

func RndWithExt(filename string) string {

	ext := filepath.Ext(filename)

	return RndStr() + ext
}
