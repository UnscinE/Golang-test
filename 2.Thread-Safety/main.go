package main

import (
	"fmt"
	"sync"
)

// สร้าง struct ชื่อ SafeCounter ที่ภายในเก็บค่า count (int)
type SafeCounter struct {
	// ป้องกันการเกิด Race Condition ด้วยการใช้ sync.Mutex
	mu    sync.Mutex
	count int

	//มีMethods 2 ตัวดังนี้:
	//Inc(): สําหรับเพิ่มค่า count
	Inc func()
	//Value(): สำหรับคืนค่า count ปัจจุบันออกมา
	Value func() int
}

func main() {
	//สร้าง instance ของ SafeCounter ให้เกิดการ define function Inc และ Value ตามที่ได้อธิบายไว้ใน struct
	counter := NewSafeCounter()

	var wg sync.WaitGroup
	//โค้ดนี้ต้องรองรับกรณีที่มี Goroutines จำนวนมาก
	// เช่น 1,000 routines
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	//ค่าสุดท้ายต้องถูกต้องแม่นยำ ห้ามเกิด Race Condition
	wg.Wait()
	fmt.Println("Final count:", counter.Value())
}

func NewSafeCounter() *SafeCounter {
	sc := &SafeCounter{}

	sc.Inc = func() {
		//ล็อกก่อนที่จะเพิ่มค่า count
		sc.mu.Lock()
		//ปลดล็อกหลังจากเพิ่มค่า count เสร็จ
		defer sc.mu.Unlock()
		//เพิ่มค่า count ทีละ 1
		sc.count++
	}

	sc.Value = func() int {
		//ล็อกก่อนที่จะอ่านค่า count
		sc.mu.Lock()
		//ปลดล็อกหลังจากอ่านค่า count เสร็จ
		defer sc.mu.Unlock()
		return sc.count
	}
	return sc
}
