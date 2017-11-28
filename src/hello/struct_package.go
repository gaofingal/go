package hello

import "fmt"

type OutsideType struct {
	a int     //  不能被其他包使用
	b string  //  不能被外部包使用
	C float64 // 可以被外部包使用
}

//  获取这个类型中属性的值，可以被外部使用
func (ot *OutsideType) Get() string {
	var res string
	res = fmt.Sprintf("a=%d,b='%s',C=%f", ot.a, ot.b, ot.C)
	return res
}

//  对属性进行赋值，不能被外部使用
func (ot *OutsideType) set() {

}
