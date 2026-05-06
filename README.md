# 🧪 Golang Test

แบบทดสอบสำหรับการสมัครงานตำแหน่ง **Backend Developer (Golang)**
โปรเจกต์นี้ประกอบไปด้วยโจทย์ทั้งหมด 5 ข้อ โดยเน้นทักษะสำคัญในการพัฒนา Backend ด้วยภาษา Go

---

## 📌 Overview

โจทย์ถูกออกแบบมาเพื่อทดสอบความเข้าใจในหัวข้อหลักดังนี้:

* Concurrency และ Goroutines
* Thread Safety และการจัดการข้อมูลร่วมกัน
* การออกแบบด้วย Interfaces
* การแก้ปัญหาเชิง Logic และการใช้ Map
* การสร้าง HTTP Handler

---

## 📂 Project Structure

```bash
golang-test/
│
├── 1.basic-concurrency/
├── 2.thread-safety/
├── 3.interfaces/
├── 4.logic-map/
├── 5.http-handler/
├── go.mod
└── README.md
```

แต่ละโฟลเดอร์จะประกอบไปด้วย:

* ไฟล์คำตอบของแต่ละข้อ (Answer source code)

---

## 🧩 Problem List

### 1️⃣ Basic Concurrency

ทดสอบความเข้าใจเกี่ยวกับ:

* Goroutines
* Channels
* การทำงานแบบ Concurrent

---

### 2️⃣ Thread Safety

ทดสอบการจัดการข้อมูลที่ถูกเข้าถึงจากหลาย Goroutine ได้แก่:

* Mutex (`sync.Mutex`)
* Race Condition

---

### 3️⃣ Interfaces

ทดสอบการออกแบบโค้ดแบบ:

* Interface-based design
* Abstraction และ Polymorphism

---

### 4️⃣ Logic & Map

ทดสอบทักษะ:

* การคิดเชิงตรรกะ
* การใช้ Map ใน Go อย่างมีประสิทธิภาพ

---

### 5️⃣ HTTP Handler

ทดสอบการสร้าง API ด้วย Go:

* `net/http`
* Routing เบื้องต้น
* Request / Response handling

---

## 🚀 How to Run

```bash
# เข้าไปในโฟลเดอร์ของแต่ละข้อ
cd basic-concurrency

# รันโปรแกรม
go run main.go
```

---

## 🛠 Requirements

* Go version 1.26 ขึ้นไป

---

## 💡 Notes

* แต่ละข้อสามารถรันแยกกันได้
* โค้ดเน้นความเข้าใจมากกว่าความซับซ้อน
* สามารถปรับปรุงหรือ refactor เพิ่มเติมได้

---

## 👤 Author

**Tanawat Tanatka**
Computer Science Student @ Khon Kaen University
