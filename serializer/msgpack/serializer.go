package msgpack

import (
	"github.com/marcelocquadros/shortener/shortener"
	"github.com/pkg/errors"

	"github.com/vmihailenco/msgpack"
)

type MsgPackSerializer struct{}

func (s *MsgPackSerializer) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (s *MsgPackSerializer) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
