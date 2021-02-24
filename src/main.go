package main

import (
	"fmt"
	"time"
	"mos"
)

func throwPanic() {
	defer func() {
		fmt.Println(" This is useless")
		if r := recover(); r != nil {
			fmt.Printf(" Caught Panic: %v\n", r)
		}
	}()
	panic(fmt.Errorf("%s", "Panic has struck...."))
}

func throwError() (string, error) {
	return "", fmt.Errorf("I am an simple error")
}

func dummy(i int) {
	fmt.Println("I will throw panic after few seconds.")
	for j := 0; j <= i; j++ {
		fmt.Printf(" Inside the dummy: %d\n", j)
		time.Sleep(1 * time.Second)
	}
}

func dirtyPanic() {
	fmt.Println("main function")

	_, err := throwError()

	if err != nil {
		fmt.Println(" catch error...." + err.Error())
	}

	go dummy(1)
	time.Sleep(1 * time.Second)
	throwPanic()
	fmt.Println("I will come no matter what happenes...")
}

func main() {
	// br.Cmd()
	// enc.Cmd()
	// dts.Cmd()
	// cry.Cmd()
	// dirtyPanic()
	mos.Cmd()
}
