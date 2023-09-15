package codeanalysis

func NewLexer(t string) Lexer {
	return Lexer{
		Text: t,
	}
}
