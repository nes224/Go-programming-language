package main

// Struct 
// มันคือ Type ที่เก็บรวบรวม field ต่างๆ เพื่อที่จะจัดกลุ่ม data เข้าไว้ด้วยกัน ง่ายๆคือเหมือนเป็นกล่องใบหนึ่ง ที่เราจะมัดของที่่เกี่ยวข้องกัน นัดใส่ลงไป
// // ประกาศตัวแปรแบบเต็ม
// var bookTitle string
// var price float32
// var stock int
// var avaliable bool
// bookTitle = "Clean Code"
// price = 22.53
// stock = 33
// avaliable = true

// แบบ short หรือสายย่อ
// bookTitle := "Clean Code"
// price := 22.53
// stock := 33
// available := true

// Composite data เพื่อ group เข้าด้วยกัน ในภาษา Go เราจะใช้ Declare Struct เพื่อให้สามารถ group fileds เข้าด้วยกัน
type Book struct {
	bookTitle string
	price float32
	stock int
	available bool
}

// pass object
// func isAvailable(b *Book) bool {
// 		return b.available
// }
// เราไม่ควรที่จะส่ง object ของ Book มาแต่ isAvailable ควรจะรู้จักกับ Book ได้เองเลย คือ function isAvailable เป็นส่วนหนึ่งของ Book เสมอ
// เลยต้องทำให้ isAvailable มันสามารถเรียกผ่าน object ของ Book ได้เอง โดยวิธีการทำง่ายมาก คือ เราจะใช้ receiver func (b *Book)isAvailable() bool { return b.available}

func (b *Book) isAvailable() bool {
	return b.available
}

func structFunc() {
	// ซึ่งจะทำให้เราสามารถเรียกใช้ Available ผ่าน book ได้
	b := Book{}
	b.bookTitle = "Clean Code"
	b.price = 22.53
	b.stock = 33
	b.available = true
	// เราสามารถเรียกใช้ isAvailable ผ่าน object Book ได้แล้ว จากเดินแทนที่เราเองจะต้องเรียก isAvailable(b) --> b.isAvailable()
	b.isAvailable()

}