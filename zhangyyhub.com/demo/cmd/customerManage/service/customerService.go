package service

import (
	"zhangyyhub.com/demo/cmd/customerManage/model"
)

// add, delete, correct and check customers.
type CustomerService struct {
	customers   []model.Customer // 客户切片
	customerNum int              // 可作为新客户编号
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 0
	return customerService
}

// list customer.
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// add customer.
func (cs *CustomerService) Add(customer model.Customer) bool {
	// 添加客户时自动添加ID
	cs.customerNum++
	customer.Id = cs.customerNum

	// 添加客户到客户列表
	cs.customers = append(cs.customers, customer)
	return true
}

// find customer.
func (cs *CustomerService) FindById(id int) (index int) {
	index = -1 // 默认没有该客户
	for i, c := range cs.customers {
		if c.Id == id {
			index = i
		}
	}
	return index
}

// delete customer.
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	} else {
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		return true
	}
}

// update customer.
func (cs *CustomerService) Update(index int, name string, gender string, age int, phone string, email string) bool {
	if index == -1 {
		return false
	} else {
		cs.customers[index].Name = name
		cs.customers[index].Gender = gender
		cs.customers[index].Age = age
		cs.customers[index].Phone = phone
		cs.customers[index].Email = email
		return true
	}
}
