package main
 
type Employee struct {
    EmpName   string  `json: "empname"`
    EmpSalary float64 `json: "salary"`
    Email     string  `json: "email"`
}
