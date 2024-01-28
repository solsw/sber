package main

func main() {
	if err := mainprim(); err != nil {
		// if err := imageprim(); err != nil {
		panic(err)
	}
}
