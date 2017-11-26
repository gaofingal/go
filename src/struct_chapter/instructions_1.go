package main

import "fmt"

//  结构体中的属性为结构体类型

type Log struct {
	msg string
}
type Customer struct {
	name string
	log  *Log
}

func main() {
	c := &Customer{"gao", &Log{"you can write your message"}}
	c.getLog().addMessage("this is my message")
	fmt.Println(c.getLog().getMsg())

}

func (l *Log) addMessage(s string) {
	l.msg += "\n" + s
}

func (l *Log) getMsg() string {
	return l.msg
}

func (c *Customer)getLog() *Log{
	return c.log
}