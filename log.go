package log

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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
}

type Log struct {
	path string
	buf chan string
	handler *bufio.Writer
	file *os.File
	log Logger
	fileName string
}

func NewLog(path string,slog Logger)*Log{
	buf := make(chan string,1024)
	fileName:=slog.GetFileName()
	npath := path+"/"+fileName
	option := os.O_WRONLY |os.O_APPEND | os.O_CREATE
	file,err:=os.OpenFile(npath,option,755)
	if err !=nil{
		panic(err)
	}
	w := bufio.NewWriter(file)
	return &Log{
		path:path,
		buf:buf,
		log:slog,
		fileName:fileName,
		handler:w,
		file:file,
	}
}

func (l *Log)Run(){
	go func() {
		for{
			select {
			case buf:=<-l.buf:
				fileName := l.log.GetFileName()
				if fileName != l.fileName{
					l.createFile(fileName)
				}
				fmt.Println("=====>"+buf)
				l.handler.WriteString(buf+"\n")
				l.handler.Flush()
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
	l.Write(WARNING,format,args)
}

func (l*Log)createFile(fileName string){
	if l.file != nil{
		l.file.Close()
	}
	l.fileName = fileName
	npath := l.path+"/"+l.fileName
	option := os.O_RDWR|os.O_APPEND | os.O_CREATE
	file,err:=os.OpenFile(npath,option,755)
	if err != nil{
		panic(err)
	}
	l.file = file
	l.handler = bufio.NewWriter(file)
}

func (l *Log)Flush(){
	l.handler.Flush()
}

func (l *Log)Close(){
	l.handler.Flush()
	l.file.Close()
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
