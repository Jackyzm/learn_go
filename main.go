package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	bbb "learn_go/pro"
	// bookType "myType"
)

// Books 书
type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var a = true
	var b = "21312"
	fmt.Println("Hello, World!")
	fmt.Println(a)
	fmt.Println("--------------------")
	fmt.Println(b)
	fmt.Println("--------------------")
	fmt.Printf(b)
	fmt.Println("--------------------")
	fmt.Println(b)
	fmt.Println("The time is", time.Now())
	fmt.Println("--------------------")
	rand.Seed(100)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("--------------------")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("--------------------")
	fmt.Println(math.Pi)
	fmt.Println("--------------------")
	fmt.Println(add(123, 456))
	fmt.Println("--------------------")
	fmt.Printf("aaa %d dsd.\n", 666)
	fmt.Println("--------------------")
	fmt.Println(&a)

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	changeBook(&Book1)

	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	bbb.Bool()
}

/**
 * @description:
 * @param x {int}
 * @param y {int}
 * @return: int
 */
func add(x int, y int) int {
	return x + y
}

/**
 * @description:
 * @param book {struct}
 */
func changeBook(book *Books) {
	book.title = "book1_change"
}

/**
 * @description:
 * @param book {struct}
 */
func printBook(book Books) {
	fmt.Println("--------------------")
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
	fmt.Println("--------------------")
}
