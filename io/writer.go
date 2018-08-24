package io

import (
	"bufio"
	"io"
	"github.com/cool/coollog/common"
	"time"
)
type Writer interface{
	Writef(msg string)
}
type ConsoleWriter struct {
	In io.Writer
	Writer *bufio.Writer
	Lexs *[]*Lex
}

func(c*ConsoleWriter)Writef(msg string){
	for _,lx:= range *c.Lexs {
		if lx.tok == common.DATE{
			c.Writer.WriteString(time.Now().Format(lx.lit))
		}else if lx.tok == common.SPACE{
			c.Writer.WriteString(lx.lit)
		}else if lx.tok == common.MSG{
			c.Writer.WriteString(msg)
		}else if lx.tok == common.NEXT{
			c.Writer.WriteString("\n")
		}
	}
	c.Writer.Flush()
}

func NewConsoleWriter(in io.Writer)*ConsoleWriter{
	cw := new(ConsoleWriter)
	cw.Writer = bufio.NewWriterSize(in, common.DefaultBufferSize)
	cw.Lexs = &[]*Lex{}
	return cw
}
