package _1

import (
	"bytes"
	"encoding/binary"
)

func uintToByte(num uint64)[]byte{
	var buff bytes.Buffer

	binary.Write(&buff,binary.LittleEndian,&num)

	return buff.Bytes()
}