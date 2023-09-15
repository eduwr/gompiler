package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eduwr/gompiler/codeanalysis"
)

func main() {
	println("Hello gompiler")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		print(">")
		scanner.Scan()
		line := scanner.Text()

		println(line)

		lexer := codeanalysis.NewLexer(line)

		for {
			token := lexer.NextToken()

			fmt.Printf("%v \n", token)

			if token.Kind == codeanalysis.EnfOfFileToken {
				break
			}
		}
	}
}
