package main

import (
	"fmt"

	"zhangyyhub.com/demo/cmd/customerManage/model"
	"zhangyyhub.com/demo/cmd/customerManage/service"
	"zhangyyhub.com/demo/pkg/print"
)

type customerView struct {
	input           string                   // 用户输入
	eflag           bool                     // 退出标记
	customerService *service.CustomerService // 操作客户列表
}

// list > customerService.List()
func (cv *customerView) list() {
	customers := cv.customerService.List()
	// 显示用户列表
	print.PrintColuor("34", "--------------------------客户列表--------------------------", "\n")
	print.PrintColuor("34", "编号\t姓名\t性别\t年龄\t电话\t\t邮箱", "\n")
	for _, customer := range customers {
		print.PrintColuor("34", customer.GetInfo(), "\n")
	}
	print.PrintColuor("34", "--------------------------客户列表--------------------------", "\n\n")
}

// add > customerService.Add()
func (cv *customerView) add() {
	// 添加用户
	print.PrintColuor("34", "--------------------------添加客户--------------------------", "\n")
	print.PrintColuor("34", "姓名: ", "")
	name := ""
	fmt.Scanln(&name)
	print.PrintColuor("34", "性别: ", "")
	gender := ""
	fmt.Scanln(&gender)
	print.PrintColuor("34", "年龄: ", "")
	age := 0
	fmt.Scanln(&age)
	print.PrintColuor("34", "电话: ", "")
	phone := ""
	fmt.Scanln(&phone)
	print.PrintColuor("34", "邮箱: ", "")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		print.PrintColuor("34", "--------------------------添加完成--------------------------", "\n\n")
	} else {
		print.PrintColuor("31", "--------------------------添加失败--------------------------", "\n\n")
	}
}

// delete > customerService.Delete()
func (cv *customerView) delete() {
	// 删除客户
	print.PrintColuor("34", "--------------------------删除客户--------------------------", "\n")
	print.PrintColuor("34", "请输入待删除客户编号(-1退出): ", "")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	print.PrintColuor("34", "请确认是否删除(y/n): ", "")
	for {
		choice := ""
		fmt.Scanln(&choice)

		switch choice {
		case "Y", "y":
			if cv.customerService.Delete(id) {
				print.PrintColuor("34", "--------------------------删除完成--------------------------", "\n\n")
				return
			} else {
				print.PrintColuor("31", "--------------------------删除失败--------------------------", "\n\n")
				return
			}
		case "N", "n":
			print.PrintColuor("31", "取消删除!", "\n")
			return
		default:
			print.PrintColuor("31", "输入有误，请重新输入(y/n): ", "")
		}
	}
}

// update > customerService.Update()
func (cv *customerView) update() {
	// 更新客户
	print.PrintColuor("34", "--------------------------修改客户--------------------------", "\n")
	print.PrintColuor("34", "请输入待修改客户编号(-1退出): ", "")
	id := -1
	fmt.Scanln(&id)
	index := cv.customerService.FindById(id)
	if index == -1 {
		print.PrintColuor("31", "待修改客户不存在!", "\n")
		return
	}
	print.PrintColuor("34", fmt.Sprintf("姓名(%v): ", cv.customerService.List()[index].Name), "")
	name := ""
	fmt.Scanln(&name)
	cv.customerService.List()[index].Name = name
	print.PrintColuor("34", fmt.Sprintf("性别(%v): ", cv.customerService.List()[index].Gender), "")
	gender := ""
	fmt.Scanln(&gender)
	cv.customerService.List()[index].Gender = gender
	print.PrintColuor("34", fmt.Sprintf("年龄(%v): ", cv.customerService.List()[index].Age), "")
	age := 0
	fmt.Scanln(&age)
	cv.customerService.List()[index].Age = age
	print.PrintColuor("34", fmt.Sprintf("手机(%v): ", cv.customerService.List()[index].Phone), "")
	phone := ""
	fmt.Scanln(&phone)
	cv.customerService.List()[index].Phone = phone
	print.PrintColuor("34", fmt.Sprintf("邮箱(%v): ", cv.customerService.List()[index].Email), "")
	email := ""
	fmt.Scanln(&email)
	cv.customerService.List()[index].Email = email

	if cv.customerService.Update(index, name, gender, age, phone, email) {
		print.PrintColuor("34", "--------------------------修改完成--------------------------", "\n\n")
	} else {
		print.PrintColuor("34", "--------------------------修改失败--------------------------", "\n\n")
	}
}

func (cv *customerView) mainView() {
	for {
		print.PrintColuor("35", "-----------------------客户信息管理软件-----------------------", "\n")
		print.PrintColuor("35", "                       1. 添加客户", "\n")
		print.PrintColuor("35", "                       2. 修改客户", "\n")
		print.PrintColuor("35", "                       3. 删除客户", "\n")
		print.PrintColuor("35", "                       4. 客户列表", "\n")
		print.PrintColuor("35", "                       5. 结束退出", "\n")
		print.PrintColuor("35", "请选择(1~5 ): ", "")
		fmt.Scanln(&cv.input)

		switch cv.input {
		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.eflag = true
		default:
			print.PrintColuor("31", "输入有误，请重新输入!", "\n")
		}

		if cv.eflag {
			break
		}
	}

	print.PrintColuor("31", "已退出客户信息管理系统!", "\n")
}

func main() {
	// customerView instance
	customerView := customerView{
		input:           "",
		eflag:           false,
		customerService: service.NewCustomerService(),
	}
	customerView.mainView()
}
