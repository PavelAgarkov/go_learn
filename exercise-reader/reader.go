package exercise_reader

import "golang.org/x/tour/reader"

type MyReader struct{}

type MyReaderError bool

func (e MyReaderError) Error() string {
	return "too short b capacity"
}

func (r MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

func Reader() {
	reader.Validate(MyReader{})
}
