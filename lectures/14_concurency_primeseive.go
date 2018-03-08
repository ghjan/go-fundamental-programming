// A concurrent prime sieve
package main

import "fmt"

// The prime sieve: Daisy-chain FilterPrime processes.
func main() {
	primeSieveDemo()
}
func primeSieveDemo() {
	ch := make(chan int)
	// Create a new channel.
	go GeneratePrime(ch)
	// Launch GeneratePrime goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go FilterPrime(ch, ch1, prime)
		ch = ch1
	}
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func GeneratePrime(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func FilterPrime(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
