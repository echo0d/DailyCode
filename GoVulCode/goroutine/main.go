package main

import (
	"fmt"
	"time"
)

func print(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
func mod(min int, max int, div int, signal chan int) {
	for n := min; n <= max; n++ {
		if n%div == 0 {
			signal <- n
		}
	}
}

func userChoice() string {
	time.Sleep(5 * time.Second)
	return "right choice"
}

func someAction() string {
	ch := make(chan string)
	timeout := make(chan bool)

	go func() {
		res := userChoice()
		ch <- res
	}()
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	select {
	case <-timeout:
		return "Timeout occured"
	case userchoice := <-ch:
		fmt.Println("User made a choice: ", userchoice)
		return ""
	}
}

func main() {
	// strings := []string{"one", "two", "three", "four", "five"}
	// fmt.Println("结果如下：")
	// go print(strings)

	// time.Sleep(3 * time.Second)
	// fmt.Println("Exiting")

	// fsignal := make(chan int)
	// ssignal := make(chan int)
	// go mod(15, 132, 14, fsignal)
	// go mod(18, 132, 17, ssignal)

	// fourteen, seventeen := <-fsignal, <-ssignal
	// fmt.Println("Divisible by 14: ", fourteen)
	// fmt.Println("Divisible by 17: ", seventeen)

	fmt.Println(someAction())
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting")
}
