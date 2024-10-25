package main

import (
	"fmt"
	"sort"
	"strings"
)

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	text := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnoabcdefjklmnocdefghijklmnoabcghijklmnopqrstuvwxyzabcdefghijklmnopqrsxyzabcdefghijkopqrstuvwxyzabcdefklmnopqrstuvwxyzabcdefghijklmnoabcdefghijklmnrstuv"

	// สร้าง map สำหรับเก็บข้อมูล string และจำนวนสระ
	vowelCountMap := make(map[string]int)

	// แบ่งข้อความออกทุกๆ 3 ตัวอักษร
	for i := 0; i < len(text); i += 3 {
		end := i + 3
		if end > len(text) {
			end = len(text)
		}

		part := text[i:end]
		vowelCountMap[part] += countVowels(part) // ถ้า part ซ้ำ จะเพิ่ม count เข้าไป
	}

	// เก็บผลลัพธ์ใน slice เพื่อจัดเรียง
	type VowelGroup struct {
		Text  string
		Count int
	}

	var sortedGroups []VowelGroup
	for text, count := range vowelCountMap {
		sortedGroups = append(sortedGroups, VowelGroup{Text: text, Count: count})
	}

	// จัดเรียงตามจำนวนสระจากมากไปน้อย
	sort.Slice(sortedGroups, func(i, j int) bool {
		return sortedGroups[i].Count > sortedGroups[j].Count
	})

	// แสดงผลลัพธ์
	fmt.Println("ส่วนของข้อความที่จัดเรียงตามจำนวนสระจากมากไปน้อย (รวมกลุ่มซ้ำ):")
	for i, group := range sortedGroups {
		fmt.Printf("Group %d: %s (vowels: %d)\n", i+1, group.Text, group.Count)
	}
}
