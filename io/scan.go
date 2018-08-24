package io

import (
	"github.com/cool/coollog/common"

)

type Lex struct{
	tok common.Token
	lit  string
}

func NewLex(tok common.Token,lit string)*Lex{
	lx := new(Lex)
	lx.tok = tok
	lx.lit = lit
	return lx
}

func Scan(str []rune,len int,index *int)*Lex{
	switch rn := str[*index]; {
		case isWhitespace(rn):
			return ParseWhitespace(str,len,index)
		case isBegin(rn):
			tok := ParseToken(str,index,len)
			if tok == common.DATE{
				datePt := ParseTimeFormat(str,len,index)
				return NewLex(common.DATE,datePt)
			}else if tok == common.MSG{
				*index++
				return NewLex(common.MSG,"")
			}else if tok == common.NEXT{
				*index++
				return NewLex(common.NEXT,"")
			}
	default:
		*index++
	}
	return nil
}

func ParseTimeFormat(str []rune,len int,index *int)string{
	*index++
	if IsEnd(index,len){
		return common.DefTimeFormat
	}
	i:=*index
	symbol := str[*index]
	if symbol == '{'{
		for *index < len && symbol != '}'{
			symbol = str[*index]
			*index++
		}
		return string(str[i+1:*index-1])
	}
	return common.DefTimeFormat
}

func ParseWhitespace(str []rune,len int,index *int)*Lex{
	i:= *index
	for *index < len{
		if !isWhitespace(str[*index]){
			break
		}
		*index++
	}
	lit :=  string(str[i:*index])
	return NewLex(common.SPACE,lit)
}

func ParseToken(str []rune,index *int,len int)common.Token{
	*index++
	if IsEnd(index,len){
		return common.EOF
	}
	skipWhitespace(str,index,len)
	if IsEnd(index,len){
		return common.EOF
	}
	switch str[*index]{
		case 'd':
			return common.DATE
		case 'm':
			return common.MSG
		case 'n':
			return common.NEXT
		}
	return common.UNKOWN
}


func IsEnd(index *int,len int)bool{
	if *index >= len {
		return true
	}
	return false
}

func isBegin(ch rune)bool{
	return '%' == ch
}


func isWhitespace(ch rune)bool{
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func skipWhitespace(rns []rune,index *int,len int){
	for *index < len{
		if !isWhitespace(rns[*index]){
			return
		}
		*index++
	}
}