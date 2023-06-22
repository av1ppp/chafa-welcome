package main

import "errors"

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
	return errors.New("not implemented")
}
