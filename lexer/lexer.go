package lexer

import "togo/token"

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.SUB, l.ch)
	case '*':
		tok = newToken(token.MUL, l.ch)
	case '/':
		tok = newToken(token.QUO, l.ch)
	case '%':
		tok = newToken(token.REM, l.ch)
	case '&':
		tok = newToken(token.AND, l.ch)
	case '|':
		tok = newToken(token.OR, l.ch)
	case '^':
		tok = newToken(token.XOR, l.ch)
	case '<':
		tok = newToken(token.LSS, l.ch)
	case '>':
		tok = newToken(token.GTR, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '!':
		tok = newToken(token.NOT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case '[':
		tok = newToken(token.LBRACK, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '.':
		tok = newToken(token.PERIOD, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ']':
		tok = newToken(token.RBRACK, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '~':
		tok = newToken(token.TILDE, l.ch)
	//case '!=':
	//	tok = newToken(token.NEQ, l.ch)
	//case '<=':
	//	tok = newToken(token.LEQ, l.ch)
	//case '>=':
	//	tok = newToken(token.GEQ, l.ch)
	//case ':=':
	//	tok = newToken(token.DEFINE, l.ch)
	//case '...':
	//	tok = newToken(token.ELLIPSIS, l.ch)
	//case 'ILLEGAL':
	//	tok = newToken(token.ILLEGAL, l.ch)
	//case 'EOF':
	//	tok = newToken(token.EOF, l.ch)
	//case 'COMMENT':
	//	tok = newToken(token.COMMENT, l.ch)
	//case 'IDENT':
	//	tok = newToken(token.IDENT, l.ch)
	//case 'INT':
	//	tok = newToken(token.INT, l.ch)
	//case 'FLOAT':
	//	tok = newToken(token.FLOAT, l.ch)
	//case 'IMAG':
	//	tok = newToken(token.IMAG, l.ch)
	//case 'CHAR':
	//	tok = newToken(token.CHAR, l.ch)
	//case 'STRING':
	//	tok = newToken(token.STRING, l.ch)
	//case '<<':
	//	tok = newToken(token.SHL, l.ch)
	//case '>>':
	//	tok = newToken(token.SHR, l.ch)
	//case '&&':
	//	tok = newToken(token.LAND, l.ch)
	//case '||':
	//	tok = newToken(token.LOR, l.ch)
	//case '<-':
	//	tok = newToken(token.ARROW, l.ch)
	//case '++':
	//	tok = newToken(token.INC, l.ch)
	//case '--':
	//	tok = newToken(token.DEC, l.ch)
	//case '==':
	//	tok = newToken(token.EQL, l.ch)
	//case ':':
	//	tok = newToken(token.COLON, l.ch)
	//case 'break':
	//	tok = newToken(token.BREAK, l.ch)
	//case 'case':
	//	tok = newToken(token.CASE, l.ch)
	//case 'chan':
	//	tok = newToken(token.CHAN, l.ch)
	//case 'const':
	//	tok = newToken(token.CONST, l.ch)
	//case 'continue':
	//	tok = newToken(token.CONTINUE, l.ch)
	//case 'default':
	//	tok = newToken(token.DEFAULT, l.ch)
	//case 'defer':
	//	tok = newToken(token.DEFER, l.ch)
	//case 'else':
	//	tok = newToken(token.ELSE, l.ch)
	//case 'fallthrough':
	//	tok = newToken(token.FALLTHROUGH, l.ch)
	//case 'for':
	//	tok = newToken(token.FOR, l.ch)
	//case 'func':
	//	tok = newToken(token.FUNC, l.ch)
	//case 'go':
	//	tok = newToken(token.GO, l.ch)
	//case 'goto':
	//	tok = newToken(token.GOTO, l.ch)
	//case 'if':
	//	tok = newToken(token.IF, l.ch)
	//case 'import':
	//	tok = newToken(token.IMPORT, l.ch)
	//case 'interface':
	//	tok = newToken(token.INTERFACE, l.ch)
	//case 'map':
	//	tok = newToken(token.MAP, l.ch)
	//case 'package':
	//	tok = newToken(token.PACKAGE, l.ch)
	//case 'range':
	//	tok = newToken(token.RANGE, l.ch)
	//case 'return':
	//	tok = newToken(token.RETURN, l.ch)
	//case 'select':
	//	tok = newToken(token.SELECT, l.ch)
	//case 'struct':
	//	tok = newToken(token.STRUCT, l.ch)
	//case 'switch':
	//	tok = newToken(token.SWITCH, l.ch)
	//case 'type':
	//	tok = newToken(token.TYPE, l.ch)
	//case 'var':
	//	tok = newToken(token.VAR, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[startPosition:l.position])
}

func newToken(tokenType token.Toke, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
