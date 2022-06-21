package service

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

type serializer[T any] struct {
	b bytes.Buffer
	t T
}

func Serializer[T any]() *serializer[T] {
	return &serializer[T]{b: bytes.Buffer{}}
}

func (s serializer[T]) Encode(data any) (value string, err error) {
	if err = gob.NewEncoder(&s.b).Encode(data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(s.b.Bytes()), nil
}

func (s serializer[T]) EncodeList(data []T) (values []string, err error) {
	for _, d := range data {
		item, err := s.Encode(d)
		if err != nil {
			return nil, err
		}
		values = append(values, item)
	}
	return values, nil
}

func (s serializer[T]) Decode(data string) (value T, err error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return value, err
	}
	s.b.Write(b)
	if err = gob.NewDecoder(&s.b).Decode(&value); err != nil {
		return value, err
	}
	return value, nil
}

func (s serializer[T]) DecodeList(data []string) (values []T, err error) {
	for _, d := range data {
		item, err := s.Decode(d)
		if err != nil {
			return values, err
		}
		values = append(values, item)
	}
	return values, nil
}
