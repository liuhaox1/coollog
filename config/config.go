package config

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"github.com/cool/coollog/io"
)
var Writers  []io.Writer
func ReadDefault(){
	file, err := os.Open("default.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	ParseConfig(file)
}

func ParseConfig(file *os.File){
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := ConfigData{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	if v.Appenders.Console.Name != ""{
		lx := ParsePattern(v.Appenders.Console.PatternLayout.Pattern)
		cw := io.NewConsoleWriter(os.Stdout)
		cw.Lexs = lx
		//Writers = append(Writers,cw)
	}
	if v.Loggers.Root.AppenderRef.Ref != ""{

	}
}

func ParsePattern(format string)*[]*io.Lex{
	lexs := &[]*io.Lex{}
	runs := []rune(format)
	index :=0
	rlen := len(runs)
	for index < rlen-1{
		lx := io.Scan(runs,rlen,&index)
		if lx!=nil{
			*lexs= append(*lexs,lx)
		}
	}
	return lexs
}



