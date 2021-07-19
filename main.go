package main

import "fmt"

func worker(startNum int, stopNum int, checkNum int, c chan bool) {
	for i := startNum; i <= stopNum; i++ {
		if checkNum%i == 0 {
			if i != 1 {
				if i != checkNum {
					c <- false
					break
				}
			}
		}
	}
	c <- true
}

func isPrime(num int) bool {
	//make sure num % 2 != 0
	if num%2 == 0 {
		return false
	} else {

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

		if resp1 == true {
			if resp2 == true {
				return true
			} else {
				return false
			}
		} else {
			return false
		}

	}
}

func main() {
	counter := 1

	for {

		switch counter {
		case 1:
			fmt.Println(counter)
			counter++
		case 2:
			fmt.Println(counter)
			counter++
		default:

			//checking if the number is prime; print if true
			if isPrime(counter) == true {
				fmt.Println(counter)
				counter++
			} else {
				counter++
			}
		}
	}
}
