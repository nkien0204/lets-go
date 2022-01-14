package tcp_server

import (
	"encoding/binary"
	"io"
)

type String string

func (s *String) Bytes() []byte {
	return []byte(*s)
}

func (s *String) String() string {
	return string(*s)
}

func (s *String) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, StringType) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	err = binary.Write(w, binary.BigEndian, uint32(len(*s))) // 4-byte len
	if err != nil {
		return n, err
	}
	n += 4
	o, err := w.Write([]byte(*s)) // value
	return n + int64(o), err
}

func (s *String) ReadFrom(r io.Reader) (int64, error) {
	var n int64 = 0
	// var typ byte
	// err := binary.Read(r, binary.BigEndian, &typ)	// 1-byte type
	// if err != nil {
	// 	return n, err
	// }
	// n += 1
	// if typ != StringType {
	// 	return n, errors.New("invalid string")
	// }

	var len uint32
	err := binary.Read(r, binary.BigEndian, &len) // 4-byte len
	if err != nil {
		return n, err
	}
	n += 4

	buf := make([]byte, len)
	o, err := r.Read(buf) // Value
	if err != nil {
		return n, err
	}
	*s = String(buf)
	return n + int64(o), err
}
