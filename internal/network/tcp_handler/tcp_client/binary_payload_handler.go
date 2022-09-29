package tcp_client

import (
	"encoding/binary"
	"io"

	"github.com/nkien0204/lets-go/internal/log"
	"go.uber.org/zap"
)

type Binary []byte

func (b *Binary) Bytes() []byte {
	return *b
}

func (b *Binary) String() string {
	return string(*b)
}

func (b *Binary) WriteTo(w io.Writer) (int64, error) {
	logger := log.Logger()
	err := binary.Write(w, binary.BigEndian, BinaryType) // 1-byte type
	if err != nil {
		logger.Error("binary write type failed", zap.Error(err))
		return 0, err
	}
	var n int64 = 1
	err = binary.Write(w, binary.BigEndian, uint32(len(*b))) // 4-byte len
	if err != nil {
		logger.Error("binary write len failed", zap.Error(err))
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

	logger := log.Logger()
	var len uint32
	err := binary.Read(r, binary.BigEndian, &len) // 4-byte len
	if err != nil {
		logger.Error("binary read len failed", zap.Error(err))
		return n, err
	}
	n += 4
	if len > MaxPacketSize {
		logger.Error("exceed max packet size")
		return n, ErrMaxPacketSize
	}

	*b = make([]byte, len)
	o, err := r.Read(*b) // value
	return n + int64(o), err
}
