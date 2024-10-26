package helper

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func WriteOnTopTCP(c net.Conn, data []byte) error {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(len(data)))

	_, err := c.Write(bs)
	if err != nil {
		return fmt.Errorf("failed to write length to destination, %w", err)
	}

	_, err = c.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write message to destination, %w", err)
	}

	return nil
}

func ReadOnTopTCP(reader *bufio.Reader) ([]byte, error) {
	var length uint16
	err := binary.Read(reader, binary.BigEndian, &length)
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return nil, err
	} else if length == 0 {
		return nil, errors.New("received zero length ISO8583 message")
	} else if err != nil {
		return nil, fmt.Errorf("failed to read ISO8583 length, %w", err)
	}

	data := make([]byte, length)
	var content []byte

	bytesRead, err := io.ReadAtLeast(reader, data, int(length))
	if bytesRead > 0 {
		content = data[:bytesRead]
	}

	if err != nil {
		// Read if message contains QUIT. QUIT was sent from NLB (network load balancer).
		// With content of QUIT, usually it will causing a lot of UnexpectedEOF error
		if errors.Is(err, io.ErrUnexpectedEOF) && strings.Contains(string(content), "QUIT") {
			return nil, err
		}

		return nil, fmt.Errorf("failed to read iso8583 data, %w", err)
	}

	return data, nil
}

func UnpackBytesFromString(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	// log.Println(len(src), len(dst), n)
	return dst[:n], nil
}
