package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	/*
		【结构体定义】
		type struct_variable_type struct{
			member definition;
			member definition;
			member definition;
		}
		一旦定义了结构体类型，它就能用于变量的声明，语法格式
		variable_name := structure_variable_type {value1,value2,value3}
		使用结构体成员，使用(.)操作符
	*/

	var book1 Books

	book1.title = "Go is god"
	book1.author = "fingal"
	book1.subject = "Go's practice"
	book1.book_id = 100

	printBook(book1)
	fmt.Println()
	printBook1(&book1)

}

func printBook(books Books) {
	fmt.Printf("Book's title : %s\n", books.title)
	fmt.Printf("Book's author : %s\n", books.author)
	fmt.Printf("Book's subject : %s\n", books.subject)
	fmt.Printf("Book's book_id : %d\n", books.book_id)
}

func printBook1(books *Books) {
	fmt.Printf("Book's title : %s\n", books.title)
	fmt.Printf("Book's author : %s\n", books.author)
	fmt.Printf("Book's subject : %s\n", books.subject)
	fmt.Printf("Book's book_id : %d\n", books.book_id)
}
