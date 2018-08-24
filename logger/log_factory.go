package logger

import (
	"github.com/cool/coollog/config"
	"sync"
)
var InitCompleted bool
var Lock sync.Mutex
var LogMap map[string]*Log = make(map[string]*Log)
var MsgChan = make(chan Msg)

type Msg struct {
	data string
}

type Log struct {
	Name string
}

func NewLog(name string)*Log{
	var log *Log = new(Log)
	log.Name = name
	return log
}

func (l*Log)Info(data string){
	for _,wt := range config.Writers{
		wt.Writef(data)
	}
}
func (l*Log)Error(data string){

}
func (l*Log)Warn(data string){

}

func GetLog(name string)*Log{
	Init()
	log,ok := LogMap[name]
	if ok{
		return log
	}
	return initLogLock(name)
}

func initLogLock(name string)*Log{
	Lock.Lock()
	log,ok := LogMap[name]
	if ok{
		return log
	}
	log = NewLog(name)
	LogMap[name] = log
	defer Lock.Unlock()
	return log
}

func Init(){
	if !IsLoadConfig(){
		config.ReadDefault()
	}
}

func IsLoadConfig()bool{
	if InitCompleted{
		return true
	}
	Lock.Lock()
	if InitCompleted{
		return true
	}else{
		InitCompleted = true
	}
	defer Lock.Unlock()
	return false
}