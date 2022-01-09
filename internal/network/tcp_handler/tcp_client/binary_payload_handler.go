package tcp_client

import (
	"encoding/binary"
	"io"
)

type Binary []byte

func (b *Binary) Bytes() []byte {
	return *b
}

func (b *Binary) String() string {
	return string(*b)
}

func (b *Binary) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, BinaryType) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	err = binary.Write(w, binary.BigEndian, uint32(len(*b))) // 4-byte len
	if err != nil {
		return n, err
	}
	n += 4
	o, err := w.Write(*b) // value
	return n + int64(o), err
}

func (b *Binary) ReadFrom(r io.Reader) (int64, error) {
	var n int64 = 0
	// var typ byte
	// err := binary.Read(r, binary.BigEndian, &typ)	// 1-byte type
	// if err != nil {
	// 	return n, err
	// }
	// n += 1
	// if typ != BinaryType {
	// 	return n, errors.New("invalid binary")
	// }

	var len uint32
	err := binary.Read(r, binary.BigEndian, &len) // 4-byte len
	if err != nil {
		return n, err
	}
	n += 4
	if len > MaxPacketSize {
		return n, ErrMaxPacketSize
	}

	*b = make([]byte, len)
	o, err := r.Read(*b) // value
	return n + int64(o), err
}
