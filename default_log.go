package log

import "time"

type MyLog struct {

}

func (l *Log)GetFileName()string{
	t := time.Now()
	return "runtime_"+t.Format("2006-01-02-15")+".log"
}

var (
	logs *Log
)

func init(){
	l := Log{}
	logs = NewLog("./",&l)
	logs.Run()
}


func Info(format string,args ...interface{}){
	logs.Info(format,args...)
}

func Debug(format string,args ...interface{}){
	logs.Debug(format,args...)
}

func Error(format string,args ...interface{}){
	logs.Error(format,args...)
}

func Warning(format string,args ...interface{}){
	logs.Warning(format,args...)
}