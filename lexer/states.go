package lexer

func PackageState(l *L) StateFunc {
	// for now just handle the package token and identifier
	l.Take()
}

func NumberState(l *L) StateFunc {
	l.Take("0123456789")
	l.Emit(INT)
	if l.Peek() == '.' {
		l.Next()
		l.Emit(PERIOD)
		return IdentState
	}

	return nil
}

func IdentState(l *L) StateFunc {
	return nil
}
