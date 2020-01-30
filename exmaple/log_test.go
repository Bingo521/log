package exmaple

import "testing"

func TestExample(t *testing.T){
	Logger.Run()
	Logger.Info("xx-")
	Logger.Close()
}
