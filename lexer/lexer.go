package lexer

type Lexer struct {
	input        string
	position     int // 입력에서 현재 위치 (현재 문자를 가리킴)
	readPosition int // 입력에서 현재 읽는 위치 (현재 문자의 다음을 가리킴)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
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
