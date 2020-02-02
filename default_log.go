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
	Logs *Log
)

func init(){
	l := newMyLog()
	Logs = NewLog("./",l)
	Logs.Run()
}