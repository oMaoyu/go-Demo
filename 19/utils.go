package _9

import (
	"bytes"
	"encoding/binary"
)

func uintToByte(num uint64)[]byte{
	var buff bytes.Buffer

	err:= binary.Write(&buff,binary.LittleEndian,&num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}