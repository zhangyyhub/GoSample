package utils

import (
	"fmt"

	"zhangyyhub.com/demo/pkg/print"
)

type Bookkeeping struct {
	balance float64 // 账户余额
	money   float64 // 收支金额
	note    string  // 收支说明
	detail  string  // 收支详情

	key      string // 菜单选项
	flag     bool   // 退出标记
	behavior bool   //是否存在收支明细
}

// 方法：子菜单显示(收支明细)
func (bookkeeping *Bookkeeping) detailsSubmenu() {
	print.PrintColuor("34", "-----------------------当前收支明细记录-----------------------", "\n")
	if bookkeeping.behavior {
		print.PrintColuor("34", bookkeeping.detail, "\n")
	} else {
		print.PrintColuor("31", "当前没有收支明细，请添加一笔收支明细吧!", "\n")
	}
	fmt.Println()
	fmt.Println()
}

// 方法：子菜单显示(登记收入)
func (bookkeeping *Bookkeeping) incomeSubmenu() {
	print.PrintColuor("34", "-----------------------登记收入明细信息-----------------------", "\n")
	print.PrintColuor("34", "本次收入金额: ", "")
	fmt.Scanln(&bookkeeping.money)           // 收入金额
	bookkeeping.balance += bookkeeping.money // 余额变化
	print.PrintColuor("34", "本次收入说明: ", "")
	fmt.Scanln(&bookkeeping.note)
	// 收支详情更新
	bookkeeping.detail += fmt.Sprintf("收入\t\t%v\t\t%v\t\t%v\n", bookkeeping.money, bookkeeping.balance, bookkeeping.note)

	// 收支行为标记
	bookkeeping.behavior = true
	fmt.Println()
	fmt.Println()
}

// 方法：子菜单显示(登记支出)
func (bookkeeping *Bookkeeping) expenditureSubmenu() {
	print.PrintColuor("34", "-----------------------登记支出明细信息-----------------------", "\n")
	print.PrintColuor("34", "本次支出金额: ", "")
	fmt.Scanln(&bookkeeping.money)           // 支出金额
	bookkeeping.balance -= bookkeeping.money // 余额变化
	print.PrintColuor("34", "本次支出说明: ", "")
	fmt.Scanln(&bookkeeping.note)
	// 收支详情更新
	bookkeeping.detail += fmt.Sprintf("支出\t\t%v\t\t%v\t\t%v\n", bookkeeping.money, bookkeeping.balance, bookkeeping.note)

	// 收支行为标记
	bookkeeping.behavior = true
	fmt.Println()
	fmt.Println()
}

// 方法：子菜单显示(结束退出)
func (bookkeeping *Bookkeeping) exitSubmenu() {
	print.PrintColuor("31", "你确定要退出吗?(y/n): ", "")
	quit := ""

	for {
		fmt.Scanln(&quit)
		if quit == "y" || quit == "n" {
			break
		}
		print.PrintColuor("31", "你的输出有误，请重新输入!(y/n): ", "")
	}

	if quit == "y" {
		bookkeeping.flag = false
	}
}

// 方法：主菜单显示
func (bookkeeping *Bookkeeping) MainMenu() {
	// 初始化赋值

	// 循环打印菜单
	for {
		print.PrintColuor("35", "-----------------------家庭收支记账软件-----------------------", "\n")
		print.PrintColuor("35", "                       1. 收支明细", "\n")
		print.PrintColuor("35", "                       2. 登记收入", "\n")
		print.PrintColuor("35", "                       3. 登记支出", "\n")
		print.PrintColuor("35", "                       4. 结束退出", "\n")
		print.PrintColuor("35", "请选择(1~4): ", "")
		fmt.Scanln(&bookkeeping.key)

		switch bookkeeping.key {
		case "1":
			bookkeeping.detailsSubmenu()
		case "2":
			bookkeeping.incomeSubmenu()
		case "3":
			bookkeeping.expenditureSubmenu()
		case "4":
			bookkeeping.exitSubmenu()
		default:
			print.PrintColuor("31", "请输出正确的选项!!!", "\n")
		}

		if !bookkeeping.flag {
			break
		}
	}
}

// 工厂方法
func NewBookkeeping() *Bookkeeping {
	return &Bookkeeping{
		// 方法：初始化赋值
		balance:  0.0,
		money:    0.0,
		note:     "",
		detail:   "收支情况\t收支金额\t账户余额\t收支说明\n",
		key:      "",
		flag:     true,
		behavior: false,
	}
}
