package main

import "fmt"

type Phone interface {
	call()
	message()
}

type vivoPhone struct {
}

type huaweiPhone struct {
}

func main() {
	/*
		【接口定义】
		type interface_name interface{
			method_name1 [return type]
			method_name2[return type]
			method_name3 [return type]
		}
		// 定义结构体
		type struct_name struct{
		}
		// 实现结构方法
		func (struct_name_variable struct_name) method_name1()[return_type]{
		}
	*/

	var vivo vivoPhone
	var huawei huaweiPhone

	vivo.call()
	huawei.call()

}

func (vp vivoPhone) call() {
	fmt.Println("this is vivoPhne call")
}

func (hwPhone huaweiPhone) call() {
	fmt.Println("this huaweiPhone call")
}
