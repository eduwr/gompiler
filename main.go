package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/eduwr/gompiler/syntaxkind"
)

func main() {
	println("Hello gompiler")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		print(">")
		scanner.Scan()
		line := scanner.Text()

		println(line)

		lexer := Lexer{Text: line}

		for {
			token := lexer.NextToken()

			fmt.Printf("%v \n", token)

			if token.Kind == syntaxkind.EnfOfFileToken {
				break
			}
		}
	}
}

type SyntaxToken struct {
	Kind     string
	Position int
	Text     string
	Value    int
}

type Lexer struct {
	Text     string
	Position int
}

func (l *Lexer) Next() {
	l.Position++
}

func (l *Lexer) Current() string {
	if l.Position >= len(l.Text) {
		return "\x00"
	}

	return string(l.Text[l.Position])
}

func (l *Lexer) NextToken() *SyntaxToken {
	if l.Position >= len(l.Text) {
		return &SyntaxToken{
			Kind:     syntaxkind.EnfOfFileToken,
			Position: l.Position,
			Text:     "\x00",
		}
	}

	if isDigit(l.Current()) {
		start := l.Position
		for isDigit(l.Current()) {
			l.Next()
		}

		s := l.Text[start:l.Position]

		v, err := strconv.Atoi(s)
		if err != nil {
			println("invalid number token")
		}

		return &SyntaxToken{
			Kind:     syntaxkind.NumberToken,
			Position: start,
			Text:     s,
			Value:    v,
		}
	}

	if isWhitespace(l.Current()) {
		start := l.Position

		for isWhitespace(l.Current()) {
			l.Next()
		}

		s := l.Text[start:l.Position]

		return &SyntaxToken{
			Kind:     syntaxkind.WhiteSpaceToken,
			Position: start,
			Text:     s,
		}
	}

	t := l.Current()
	p := l.Position
	l.Next()
	switch t {
	case "+":
		return &SyntaxToken{
			Kind:     syntaxkind.PlusToken,
			Position: p,
			Text:     t,
		}
	case "-":
		return &SyntaxToken{
			Kind:     syntaxkind.MinusToken,
			Position: p,
			Text:     t,
		}
	case "/":
		return &SyntaxToken{
			Kind:     syntaxkind.SlashToken,
			Position: p,
			Text:     t,
		}

	case "*":
		return &SyntaxToken{
			Kind:     syntaxkind.SlashToken,
			Position: p,
			Text:     t,
		}
	default:
		return &SyntaxToken{
			Kind:     syntaxkind.EnfOfFileToken,
			Position: p,
			Text:     "\x00",
		}
	}
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func isWhitespace(s string) bool {
	return s == " "
}
