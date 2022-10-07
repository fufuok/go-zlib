package zlib

import "github.com/fufuok/go-zlib/native"

func checkClosed(c native.StreamCloser) error {
	if c.IsClosed() {
		return errIsClosed
	}
	return nil
}
