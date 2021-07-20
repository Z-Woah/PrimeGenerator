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

	switch num {
	case 1:
		return false
	case 2:
		return true
	}

	//make sure num % 2 != 0
	if num%2 == 0 || num%5 == 0 {
		return false
	} // Remove else statement, it is not really needed and helps save indentation space

	var numHalfed int = num / 2
	var firstHalf int = numHalfed / 2
	var secondHalf int = numHalfed - firstHalf

	channel := make(chan bool)

	// Worker for first half
	go worker(1, firstHalf, num, channel)
	// Worker for second half
	go worker(secondHalf, numHalfed, num, channel)

	//resp1 := <-c1
	//resp2 := <-c2

	return <-channel && <-channel

}

func main() {
	counter := 1

	for {

		//checking if the number is prime; print if true
		if isPrime(counter) {
			fmt.Println(counter)
		}
		counter++
	}
}
