package log

import (
	"fmt"
	"github.com/fatih/color"
	"errors"
)

const (  
	Basic = iota 
	Blue = iota  
	Red = iota  
)

type Logger interface {
	Log(message string)
}

var loggers map[int]Logger

func init(){
	loggers = make(map[int]Logger)
	loggers[Blue] = &ColorLogger{Color: color.FgBlue}
	loggers[Red] 	= &ColorLogger{Color: color.FgRed}
	loggers[Basic] 	= &ConsoleLogger{}
}
func NewLogger(loggerTypes ...int) (Logger,error){
	var loggerType int
	if len(loggerTypes) == 0 {
		return loggers[Basic], nil
	}	else if len(loggerTypes) == 1{
		loggerType = loggerTypes[0]
		return loggers[loggerType],nil
	} else {
		return nil, errors.New("NewLogger called with more than 1 parameter")
	}
}

type ConsoleLogger struct{}
func (l ConsoleLogger) Log(message string){
	fmt.Println(message)
}

type ColorLogger struct{
	Color color.Attribute
}
func (cl ColorLogger) Log(message string){
	red := color.New(cl.Color)
	boldRed := red.Add(color.Bold)
	boldRed.Println(message)
}

