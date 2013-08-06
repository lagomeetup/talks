package main

type MyWriter struct {
}

func (w *MyWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func main() {
	return
}
