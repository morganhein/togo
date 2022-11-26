package token

type Toke string

var keywords = map[string]Toke{
	"break":       BREAK,
	"case":        CASE,
	"chan":        CHAN,
	"const":       CONST,
	"continue":    CONTINUE,
	"default":     DEFAULT,
	"defer":       DEFER,
	"else":        ELSE,
	"fallthrough": FALLTHROUGH,
	"for":         FOR,
	"func":        FUNC,
	"go":          GO,
	"goto":        GOTO,
	"if":          IF,
	"import":      IMPORT,
	"interface":   INTERFACE,
	"map":         MAP,
	"package":     PACKAGE,
	"range":       RANGE,
	"return":      RETURN,
	"select":      SELECT,
	"struct":      STRUCT,
	"switch":      SWITCH,
	"type":        TYPE,
	"var":         VAR,
}

const (
	// Special tokens
	ILLEGAL Toke = "ILLEGAL"
	EOF     Toke = "EOF"
	COMMENT Toke = "COMMENT"

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  Toke = "IDENT"  // main
	INT    Toke = "INT"    // 12345
	FLOAT  Toke = "FLOAT"  // 123.45
	IMAG   Toke = "IMAG"   // 123.45i
	CHAR   Toke = "CHAR"   // 'a'
	STRING Toke = "STRING" // "abc"

	// Operators and delimiters
	ADD            Toke = "ADD"            // +
	SUB            Toke = "SUB"            // -
	MUL            Toke = "MUL"            // *
	QUO            Toke = "QUO"            // /
	REM            Toke = "REM"            // %
	AND            Toke = "AND"            // &
	OR             Toke = "OR"             // |
	XOR            Toke = "XOR"            // ^
	SHL            Toke = "SHL"            // <<
	SHR            Toke = "SHR"            // >>
	AND_NOT        Toke = "AND_NOT"        // &^
	ADD_ASSIGN     Toke = "ADD_ASSIGN"     // +=
	SUB_ASSIGN     Toke = "SUB_ASSIGN"     // -=
	MUL_ASSIGN     Toke = "MUL_ASSIGN"     // *=
	QUO_ASSIGN     Toke = "QUO_ASSIGN"     // /=
	REM_ASSIGN     Toke = "REM_ASSIGN"     // %=
	AND_ASSIGN     Toke = "AND_ASSIGN"     // &=
	OR_ASSIGN      Toke = "OR_ASSIGN"      // |=
	XOR_ASSIGN     Toke = "XOR_ASSIGN"     // ^=
	SHL_ASSIGN     Toke = "SHL_ASSIGN"     // <<=
	SHR_ASSIGN     Toke = "SHR_ASSIGN"     // >>=
	AND_NOT_ASSIGN Toke = "AND_NOT_ASSIGN" // &^=
	LAND           Toke = "LAND"           // &&
	LOR            Toke = "LOR"            // ||
	ARROW          Toke = "ARROW"          // <-
	INC            Toke = "INC"            // ++
	DEC            Toke = "DEC"            // --
	EQL            Toke = "EQL"            // ==
	LSS            Toke = "LSS"            // <
	GTR            Toke = "GTR"            // >
	ASSIGN         Toke = "ASSIGN"         // =
	NOT            Toke = "NOT"            // !
	NEQ            Toke = "NEQ"            // !=
	LEQ            Toke = "LEQ"            // <=
	GEQ            Toke = "GEQ"            // >=
	DEFINE         Toke = "DEFINE"         // :=
	ELLIPSIS       Toke = "ELLIPSIS"       // ...

	LPAREN    Toke = "LPAREN"    // (
	LBRACK    Toke = "LBRACK"    // [
	LBRACE    Toke = "LBRACE"    // {
	COMMA     Toke = "COMMA"     // ,
	PERIOD    Toke = "PERIOD"    // .
	RPAREN    Toke = "RPAREN"    // )
	RBRACK    Toke = "RBRACK"    // ]
	RBRACE    Toke = "RBRACE"    // }
	SEMICOLON Toke = "SEMICOLON" // ;
	COLON     Toke = "COLON"     // :
	QUESTION  Toke = "QUESTION"  // ?

	// Keywords
	BREAK    Toke = "BREAK"
	CASE     Toke = "CASE"
	CHAN     Toke = "CHAN"
	CONST    Toke = "CONST"
	CONTINUE Toke = "CONTINUE"

	DEFAULT     Toke = "DEFAULT"
	DEFER       Toke = "DEFER"
	ELSE        Toke = "ELSE "
	FALLTHROUGH Toke = "FALLTHROUGH"
	FOR         Toke = "FOR"

	FUNC   Toke = "FUNC"
	GO     Toke = "GO"
	GOTO   Toke = "GOTO"
	IF     Toke = "IF"
	IMPORT Toke = "IMPORT"

	INTERFACE Toke = "INTERFACE"
	MAP       Toke = "MAP"
	PACKAGE   Toke = "PACKAGE"
	RANGE     Toke = "RANGE"
	RETURN    Toke = "RETURN"

	SELECT Toke = "SELECT"
	STRUCT Toke = "STRUCT"
	SWITCH Toke = "SWITCH"
	TYPE   Toke = "TYPE"
	VAR    Toke = "VAR"

	TILDE Toke = "TILDE"
)

type Token struct {
	Type     Toke
	Literal  string
	Position Position
	Mut      bool // if a literal, is it mutable?
}

type Position struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (rune count)
}

func LookupIdent(ident string) Toke {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
