package log

import (
	"os"
)

type MyLog struct {
	file *os.File
}

func newMyLog()*MyLog{
	return &MyLog{
		file:os.Stdout,
	}
}
func (l *MyLog)GetFile()*os.File{
	return l.file
}

func (l *MyLog)GetFileName()string{
	return "stdout"
}

func (l *MyLog)Flush(){

}

func (l *MyLog)Close(){
	l.file.Close()
}

var (
	logs *Log
)

func init(){
	l := newMyLog()
	logs = NewLog("./",l)
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

func Close(){
	logs.Close()
}