package closezeug

import (
	"io"
)

// CloseOrPanic causes a panic if the resource cannot be closed.
func CloseOrPanic(c io.Closer) {
	CloseOrHandleError(c, func(closer io.Closer, e error) { panic(e) })
}

// CloseOrSwallowError swallows the error.
func CloseOrSwallowError(c io.Closer) {
	CloseOrHandleError(c, func(closer io.Closer, e error) {})
}

// CloseOrHandleError allows the consumer to pass a function that can attempt some other action if the Close call
// fails.
func CloseOrHandleError(c io.Closer, handler func(io.Closer, error)) {
	err := c.Close()
	if err != nil {
		handler(c, err)
	}
}
