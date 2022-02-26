package embeded

import (
	"fmt"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type One struct {
	Three int
	Two   string
}

func (r *One) setThree(value int) *One {
	r.Three = value
	return r
}

func (r *One) setTwo(value string) *One {
	r.Two = value
	return r
}

func NewOne() *One {
	return new(One).
		setThree(1).
		setTwo("go to")
}

func (r *One) Read(p []byte) (n int, err error) {
	return r.Three, nil
}

func (r *One) Write(p []byte) (n int, err error) {
	return 2, nil
}

func (r *One) Embed() {
	var p []byte
	read, ok := r.Read(p)
	write, ok2 := r.Write(p)
	fmt.Println(read, write, ok, ok2)
}
