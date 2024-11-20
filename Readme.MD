# Test Back-End Sandhya

A simple REST API built with Go to demonstrate CRUD operations, filtering, and data handling in memory.

## **Getting Started**

### **Run the Project**
To run the project, simply execute the following command in your terminal:

```bash
go run main.go

Get All Data : http://localhost:8080/data
Get Data by Code : http://localhost:8080/data/detail?code=tfx
Get Data by Filter : http://localhost:8080/data/filter?model=humanoid&tech=cyborg
