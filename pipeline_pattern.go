package main

import "fmt"

func main() {
	LaunchPipeline(4)
}

func LaunchPipeline(amount int) int {
	firstCh := generator(amount)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)

	result := <-thirdCh
	return result
}

func generator(max int) <-chan int {
	outChint := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			fmt.Println("genrator", i)
			outChint <- i
		}
		close(outChint)
	}()

	return outChint
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int

		for v := range in {
			fmt.Println("sum", v)
			sum += v
		}

		out <- sum
		close(out)
	}()

	return out
}
