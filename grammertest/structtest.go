package main

import "fmt"
/*
struct 初始化方式
1.声明之后使用.方式，给其内部设值
2.直接声明时初始化
book3:=Books{"go3","xiao","idsh",23}
3.struct是值引用
函数体需要指针才能修改地址上的值
 */
type Books struct {
	title string
	author string
	subject string
	book_id int
}

func changeBook(book Books){
	book.title="改变了"
}
func changeBookPoint(book *Books){
	book.title="改变了"
}


func main() {
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407
	Book1.title+="sdu"

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700


	book3:=Books{"go3","xiao","idsh",23}

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)

	fmt.Printf( "Book 2 book_id : %d\n", book3.book_id)


	changeBook(book3)
	fmt.Println( "Book 3 book_name : ", book3.title)
	changeBookPoint(&book3)
	fmt.Println( "Book 3 book_name : ", book3.title)
}