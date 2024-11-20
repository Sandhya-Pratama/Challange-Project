# Test Back-End Sandhya

Sebuah proyek yang dibuat untuk melaksanakan test coding pada **PT Maxchat Inovasi Indonesia**. Proyek ini merupakan REST API sederhana yang berfungsi untuk mengelola data secara in-memory, termasuk fitur CRUD dan filtering data.

---
### **Run the Project**
Untuk menjalankan proyek, jalankan perintah berikut di terminal:

```bash
go run main.go
```
---

### **Endpoint of the Project**
Get All Data : 
```http://localhost:8080/data```
Get Data by Code : 
```http://localhost:8080/data/detail?code=tfx```
Get Data by Filter : 
```http://localhost:8080/data/filter?model=humanoid&tech=cyborg```
