package log

import "time"

type MyLog struct {

}

func (l *Log)GetFileName()string{
	t := time.Now()
	return "runtime_"+t.Format("2006-01-02-15")+".log"
}

var (
	Logs *Log
)

func init(){
	l := Log{}
	Logs = NewLog("./",&l)
}

