package model

import "fmt"

// definition customer struct.
type Customer struct {
	Id     int    // 客户编号
	Name   string // 客户姓名
	Gender string // 客户性别
	Age    int    // 客户年龄
	Phone  string // 客户电话
	Email  string // 客户邮箱
}

func NewCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// show customers.
func (customer *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", customer.Id, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
}
