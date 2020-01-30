package exmaple

import (
	"time"
	"work_space/log"
)

type Log struct {

}

func (l *Log)GetFileName()string{
	t := time.Now()
	return "runtime_"+t.Format("2006-01-02-15")+".log"
}

var (
	Logger *log.Log
)

func init(){
	l := Log{}
	Logger = log.NewLog("./",&l)
}