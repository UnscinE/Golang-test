package main

import "fmt"

func main() {
	//กำหนดให้มี Slice ของตัวเลข nums := []int{2, 7, 11, 15}
	nums := []int{2, 7, 11, 15}
	//ค่าเป้าหมาย target := 9
	target := 9

	//เรียกใช้ฟังก์ชัน TwoSum(nums, target) เพื่อหาคู่ของตัวเลขที่บวกกันแล้วได้เท่ากับ target
	result := TwoSum(nums, target)
	//แล้วพิมพ์ผลลัพธ์ออกมา
	fmt.Println("Return Index: ", result)

}

//ฟังก์ชันที่รับ nums และ target
func TwoSum(nums []int, target int) []int {
	//ห้ามใช้ Loop ซ้อน Loop (Double for-loop)
	//ให้ใช้ Map ในการเก็บค่าที่เคยเจอมาแล้วแทน
	indexMap := make(map[int]int)
	
	for currentIndex, val := range nums {
		//คำนวณค่าที่ต้องการหา complement = target - val
		//รอบแรก complement = 9 - 2 = 7
		//รอบสอง complement = 9 - 7 = 2
		complement := target - val
		if prevIndex, ok := indexMap[complement]; ok {
			//คืนค่า index ของตัวเลข 2 ตัวใน slice ที่บวกกันแล้วได้เท่ากับ target พอดี
			return []int{prevIndex, currentIndex}
		}
		//ถ้ายังไม่เจอคู่ที่บวกกันได้เท่ากับ target ให้เก็บค่าปัจจุบันลงใน Map เพื่อใช้ในการตรวจสอบในรอบถัดไป
		indexMap[val] = currentIndex
	}
	return nil
}
