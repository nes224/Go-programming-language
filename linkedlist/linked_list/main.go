package main
import "fmt"

type node struct {
	data int
	next *node
}
//   a -> b -> c -> d -> e
// head

// double linked list
//  2 <-> 7 <-> 11 <-> 16 <-> 18

type linkedList struct {
	head *node
	length int
}

// (l *linkedList) the receiver is a pointer tihs means you want to actually make modifications to this receiver
// if you just put it as a value (l linkedList) receiver then we're we just going to be working on a copy of it 
// (l *linkedList) we want to actually work on the actual receiver like change the values and stuff we need just point as a pointer.
func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListdata() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// Golang string literals
// Interpreted String Literals ""
// escape sequences -> \n, \t, \
// \n ขึ้นบรรทัดใหม่
// Raw String Literals ``` จะถูกตีความตามที่เห็น
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		// assume that we reached the end and no input value that.
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
} 

func main() {
	mylist := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:18}
	node3 := &node{data:16}
	node4 := &node{data:11}
	node5 := &node{data:7}
	node6 := &node{data:2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListdata()
	mylist.deleteWithValue(100)
	mylist.deleteWithValue(2)
	mylist.printListdata()
	mylist.deleteWithValue(10)
}

// Golang Pointers
// Go pointers are those variables that hold the address of any variables.
// * operator -> called a dereferencing operator and is used for accessing the value in the address stored in the pointer.
// & operator -> called the address operator and is used for returning the address of the variable stored in the pointer.
// unmarshalling JSON data into a struct

// slice in Go?
// Go Interfaces?
// Go map contains a key?
// Go channels and how are channels used in Golang?


// Waitgroups
// เหมือนกับพนักงานที่ทำงานในร้านอาหาร
// ในร้านอาหารเนี่ย จะมีลูกค้าเข้ามาใช้บริการ
// และก่อนที่พนักงานจะปิดร้านอาหารได้
// ต้องทำให้แน่ใจว่าลูกค้าทุกคนออกจากร้านอาหารหมดแล้ว
// Mutexs
// เหมือนกับตู้ล็อคเกอร์ตามที่สาธารณะ
// ถ้าตู้ล็อคเกอร์นั้นมีคนล็อคอยู่
// คนที่จะใช้ตู้ล็อคเกอร์นั้นต่อ
// ต้องรอให้คนที่ล็อคก่อนหน้ามาปลดล็อคก่อน
// เค้าถึงมาล็อคตู้ล็อคเกอร์นั้นได้
// ในเชิงของการเขียนโปรแกรม
// Waitgroups ช่วยในการจัดกลุ่มสิ่งที่ต้องรอให้เสร็จ
// Mutexs ทำให้ Goroutine หลาย ๆ อัน
// ผลัดกันทำอะไรบางอย่าง
// ที่มีแค่ Goroutine ตัวเดียวทำได้ ณ​ ขณะนั้น