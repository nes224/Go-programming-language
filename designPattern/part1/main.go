package main

import (
	"encoding/json"
	"fmt"
)

// Factory Method

// Design Patterns
/*
	1. Creational Design Patterns ใช้ในการกำหนดกลไกลในการสร้าง Object ต่างๆ ให้มันมีความยืดหยุ่น
	และสามารถกลับมาแก้ไขหรือใช้ใหม่ได้ง่าย
*/
/*
	2. Structural Design Patterns ใช้ในการรวม Class หลายๆ Class หรือหลายๆ Object เข้าด้วยกัน
	โดยยังคงความยืดหยุ่นเอาไว้ และสามารถแก้ไขได้ง่าย
*/
/*
	3. Behavioral Design Patterns ใช้แก้ปัญหาในการที่จะเชื่อมความสัมพันธ์ระหว่าง Object กับ Object ด้วยกัน
*/
/*
	----------------------								 -----------------------
	|		Book		 |								 |		Textbook	   |
	----------------------								 |---------------------|
	|	Title: string	 |								 |	Title: string	   |
	|	Author: string	 |								 |	Author: string	   |
	|	Page: uint		 | <---------------------------> |	Page: uint		   |
	|	Price: float64	 |								 |	Price: float64	   |
	----------------------								 -----------------------
	|.........			 |								 | .....			   |
	|newBook(type string)|								 | newTextbook()	   |
	----------------------								 -----------------------
			  ^
			  |
			  |
	----------------------
	|		Comics		 |
	|--------------------|
	|	title: string	 |
	|	Author: string	 |
	|	Page: uint		 |
	|	Price: float64	 |
	|	IsPostCard: bool |
	|--------------------|
	| ...........        |
	| new Comics()		 |
	----------------------
*/

type IBook interface {
	setTitle(string)
	setAuthor(string)
	setPage(uint)
	setPrice(float64)
	getTitle() string
	getAuthor() string
	getPage() uint
	getPrice() float64
	paint()
}

type Book struct {
	Title  string
	Author string
	Page   uint
	Price  float64
}

type Comics struct {
	PostCard IPostCard
	Book
}

type IPostCard interface {
	setPostCard(isPostCard bool)
	getPostCard() bool
}

type PostCard struct {
	IsPostCard bool
}

func (b *Comics) setTitle(title string)         { b.Title = title }
func (b *Comics) setAuthor(author string)       { b.Author = author }
func (b *Comics) setPage(page uint)             { b.Page = page }
func (b *Comics) setPrice(price float64)        { b.Price = price }
func (p *PostCard) setPostCard(isPostCard bool) { p.IsPostCard = isPostCard }

func (b *Comics) getTitle() string  { return b.Title }
func (b *Comics) getAuthor() string { return b.Author }
func (b *Comics) getPage() uint     { return b.Page }
func (b *Comics) getPrice() float64 { return b.Price }

func (b *Comics) print() {
	obj, _ := json.MarshalIndent(&b, "", " ")
	fmt.Println(string(obj))
}

func main() {

}
