package main

func main() {
	// START1 OMIT
	for i := 0; i < 1000; i++ {
		go func() {
			for {
			}
		}()
	}

	for {
	}
	// END1 OMIT
}
