package message

import (
	"encoding/gob"
	"io"
)

type data struct {
	Data interface{}
}

// Send encodes a message as a gob, then sends
// it directly to a Writer.
func Send(msg Message, w io.Writer) error {
	enc := gob.NewEncoder(w)

	return enc.Encode(data{Data: msg})
}

// Receive decodes a message from a reader and
// returns the value of it.
func Receive(r io.Reader) (Message, error) {
	var v data
	dec := gob.NewDecoder(r)

	if err := dec.Decode(&v); err != nil {
		return nil, err
	}

	return v.Data, nil
}
