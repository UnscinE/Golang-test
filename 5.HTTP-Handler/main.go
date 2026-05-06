package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//สร้าง struct สำหรับรับข้อมูลจาก request และส่งข้อมูลกลับใน response
type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	//มี endpoint ชื่อ/hello
	http.HandleFunc("/hello", HelloHandler)
	//พิมพ์ข้อความแสดงการทำงานของ server
	fmt.Println("Server is running on http://localhost:8080")

	//รันบน port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		//ถ้าเกิด error ในการรัน server ให้พิมพ์ข้อความรายละเอียดของ error นั้นออกมา
		fmt.Printf("Failed to start server: %s\n", err)
	}

}

// กรณีที่Method ไม่ใช่POST ให้ส่งกลับข้อความ "Method not allowed" และสถานะ HTTP 405
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//รับ Request Method เป็น POST เท่านั้น
	if r.Method != http.MethodPost {
		//ถ้าเป็นกรณีอื่นจะส่ง error กลับ
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}

	var req HelloRequest

	//อ่านข้อมูลจาก body ของ request และแปลง JSON เป็น struct HelloRequest
	decoder := json.NewDecoder(r.Body)
	//ป้องกันการรับ JSON ที่มีฟิลด์ที่ไม่รู้จัก
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	//ถ้าเกิด error ในการแปลง JSON ให้ส่งกลับข้อความ "Invalid JSON format" และสถานะ HTTP 422
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusUnprocessableEntity)
		return
	}

	//แกะค่า name ออกมาแล้วตอบกลับ (Response) เป็น JSON: {"message": "Hello Somchai"}
	res := HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}

	//ตั้ง header ของ response เป็น application/json และส่งสถานะ HTTP 200 OK พร้อมกับข้อมูล JSON ที่แปลงมาจาก struct HelloResponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
