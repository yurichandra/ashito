package message

import "github.com/moov-io/iso8583"

const RrnLength = 12

// UnpackMessage allows us unpack byte message and customize the value inside.
func UnpackMessage(spec iso8583.MessageSpec, data []byte) (*iso8583.Message, error) {
	m := iso8583.NewMessage(&spec)
	if err := m.Unpack(data); err != nil {
		return nil, err
	}

	return m, nil
}
