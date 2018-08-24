package config

import "encoding/xml"

type ConfigData struct {
	XMLName    xml.Name  `xml:"Config"`
	Appenders  Appenders
	Loggers	   Loggers
}

type Appenders struct {
	XMLName    xml.Name  `xml:"Appenders"`
	Console Console `xml:"Console"`
}

type Console struct {
	Name string `xml:"name,attr"`
	PatternLayout PatternLayout
}

type PatternLayout struct {
	Pattern string	`xml:"pattern,attr"`
}

type Loggers struct{
	XMLName    xml.Name  `xml:"Loggers"`
	Root Root
}
type Root struct {
	XMLName    xml.Name  `xml:"Root"`
	Level string	`xml:"level,attr"`
	AppenderRef	AppenderRef
}

type AppenderRef struct {
	XMLName    xml.Name  `xml:"AppenderRef"`
	Ref		   string	 `xml:"ref,attr"`
}