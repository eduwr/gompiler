package main

import (
	"bufio"
	"os"
)

func main() {
	println("Hello gompiler")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		print(">")
		scanner.Scan()
		line := scanner.Text()

		println(line)
	}
}
