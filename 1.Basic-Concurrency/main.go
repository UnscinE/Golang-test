package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// รัน Worker 5 คน เพื่อจัดการงาน 10 อย่าง
	RunWorkers(10, 10)
}

// เขียนฟังก์ชันชื่อ RunWorkers โดยรับ parameter 2 ตัวคือ numWorkers (จำนวนคนงาน) และ numJobs (จํานวนงานทั้งหมด)
func RunWorkers(numWorkers int, numJobs int) {
	//สร้าง Channel สำหรับส่ง Job ID
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	//สร้าง Workers จำนวน numWorkers ตัว ที่ทำงานพร้อมๆ กัน (Concurrent)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(ID int) {
			defer wg.Done()
			//แต่ละ Worker จะคอยรับงานจาก Channel
			for jobID := range jobs {
				//แล้วพิมพ์ข้อความว่ํา "Worker [ID] processing job [JobID]"
				fmt.Printf("Worker [%d] processing job [%d]\n", ID, jobID)
				//จำลองกํารทำงานด้วย time.Sleep เล็กน้อย(1 วินาที)
				time.Sleep(time.Second * 1)
			}
		}(i)
	}

	//ส่งงานทั้งหมดเข้าไปใน Channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	//ปิด channel หลังจากส่งงานทั้งหมดแล้ว
	close(jobs)

	//รอให้ทุกงานถูกทํา จนเสร็จสิ้นจริงๆ ถึงจะจบโปรแกรม
	wg.Wait()
	fmt.Println("All jobs processed.")
}
