package main

import "fmt"

func main() {
	naturals := []int{}
	const noOfGoroutines = 10
	var chans = [noOfGoroutines]chan int{}
	nlen := 10000

	bachlen := (nlen / noOfGoroutines)

	for i := 0; i < noOfGoroutines; i++ {
		chans[i] = make(chan int)
	}
	for i := 0; i < nlen; i++ {
		naturals = append(naturals, i)
	}

	for i := 1; i <= noOfGoroutines; i++ {
		batch := i * bachlen
		//fmt.Println(naturals[(batch - bachlen):batch])
		chans[i-1] = isPrime(naturals[(batch-bachlen):batch], i)

	}
	sum := 0
	for i := 0; i < noOfGoroutines; i++ {
		for _ = range chans[i] {

			sum++
		}
	}

	fmt.Println("count:", sum)

}

func isPrime(num []int, threadno int) chan int {
	primes := make(chan int)
	fmt.Println("batch", num)
	go func() {
		//fmt.Println("executing go routine")
		for _, n := range num {
			fmt.Println("testing", n)
			temp := true

			if n == 1 || n == 0 {
				temp = false
			}

			for i := 2; i < n; i++ {
				if n%i == 0 {
					//fmt.Println("false")
					temp = false
					break
				}
			}

			//fmt.Println(temp)
			if temp {
				//fmt.Println("thread", threadno, ":", n)
				//fmt.Println("prime", n)
				primes <- n
			}

		}
		close(primes)
	}()

	return primes

}
