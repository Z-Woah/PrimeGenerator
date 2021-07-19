package main

import "fmt"

func worker(startNum int, stopNum int, checkNum int, c chan bool) {
	sentinel := true
	for i := startNum; i <= stopNum; i++ {
		if checkNum%i == 0 && (i != 1 && i != checkNum) {
			sentinel = false
			break
		}
	}
	c <- sentinel
}

func isPrime(num int) bool {
	//make sure num % 2 != 0
	if num%2 == 0 {
		return false
	} // Remove else statement, it is not really needed and helps save indentation space

	var numHalfed int = num / 2
	var firstHalf int = numHalfed / 2
	var secondHalf int = numHalfed - firstHalf

	c1 := make(chan bool)
	c2 := make(chan bool)

	// Worker for first half
	go worker(1, firstHalf, num, c1)
	// Worker for second half
	go worker(secondHalf, numHalfed, num, c2)

	resp1 := <-c1
	resp2 := <-c2

	return resp1 && resp2

}

func main() {
	counter := 1

	for {

		switch counter {
		case 1:
			fallthrough
		case 2:
			fmt.Println(counter)
		default:

			//checking if the number is prime; print if true
			if isPrime(counter) {
				fmt.Println(counter)
			}
		}
		counter++
	}
}
