package common
type Token int
const(
	ILLEGAL Token = iota
	DATE
	MSG
	NEXT
	SPACE
	UNKOWN
	EOF
	END
)