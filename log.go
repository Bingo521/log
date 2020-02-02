package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	INFO = "INFO"
	DEBUG = "DEBUG"
	ERROR = "ERROR"
	WARNING = "WARNING"
)

const (
	FORMAT = "[%s] [%s:%d] "
)
type Logger interface {
	GetFileName()string
	GetFile()*os.File
	Flush()
	Close()
}

type Log struct {
	path string
	buf chan string
	file *os.File
	log Logger
}

func NewLog(path string,slog Logger)*Log{
	buf := make(chan string,1024)
	return &Log{
		path:path,
		buf:buf,
		log:slog,
	}
}

func (l *Log)Run(){
	go func() {
		for{
			select {
			case buf:=<-l.buf:
				file := l.log.GetFile()
				_,err:=file.Write([]byte(buf+"\n"))
				if err != nil{
					panic(err)
				}
			}
		}
	}()
}

func (l *Log)Write(level string,format string,args ...interface{}){
	_, filename, line, _ := runtime.Caller(2)
	prefix := fmt.Sprintf(FORMAT,level,filename,line)
	buf := fmt.Sprintf(format,args...)
	l.buf<-(prefix+buf)
}

func (l *Log)Info(format string,args ...interface{}){
	l.Write(INFO,format,args...)
}

func (l *Log)Debug(format string,args ...interface{}){
	l.Write(DEBUG,format,args...)
}

func (l *Log)Error(format string,args ...interface{}){
	l.Write(ERROR,format,args...)
}

func (l *Log)Warning(format string,args ...interface{}){
	l.Write(WARNING,format,args...)
}

func (l *Log)Flush(){
	l.log.Flush()
}

func (l *Log)Close(){
	time.Sleep(time.Second)
	l.log.Flush()
	l.log.Close()
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
