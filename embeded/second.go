package embeded

func Now() *ReadWriter {

	var one ReadWriter = new(One).
		setThree(1).
		setTwo("go to")

	one.same()

	return &one
}

func (r *One) same() {
	r.Embed()
}
