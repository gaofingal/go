package main

import "fmt"

func main() {
	/*
		【求1-100之间的素数】
	*/
	var i, j int
	i = 1

A:
	for i < 100 {
		i++
		for j = 2; j < i; j++ {
			//有被除数
			if i%j == 0 {
				goto A
			}
		}
		fmt.Printf("%d is prime \n", i)
	}
	/*
		for i = 2; i < 100; i++ {
			for j = 2; j <= (i / j); j++ {
				if (i%j == 0) {
					break; // 如果发现因子，则不是素数
				}
			}
			if (j > (i / j)) {
				fmt.Printf("%d  是素数\n", i);
			}
		}
	*/
}
