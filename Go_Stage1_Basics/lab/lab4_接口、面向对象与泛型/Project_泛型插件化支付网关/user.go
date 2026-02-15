package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Credit   int    `json:"Credit"`
}

// 进行登录校验，校验的用户数据单独存放在user.json中
func (u *User) Login() bool {
	fmt.Println("请输入姓名与密码：")
	_, err := fmt.Scanln(&u.Name, &u.Password)
	if err != nil {
		fmt.Println("输入错误：", err)
		return false
	}

	// 1. 打开 JSON 文件中的用户数据
	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println("读取用户数据失败：", err)
		return false
	}
	defer file.Close() // 确保文件最后关闭

	// 2. 解析 JSON 到用户数组
	var userList []User // 定义切片存储所有用户
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&userList)
	if err != nil {
		fmt.Println("解析用户数据失败：", err)
		return false
	}

	// 校验输入的姓名和密码
	loginSuccess := false
	for _, user := range userList {
		if user.Name == u.Name && user.Password == u.Password {
			// 校验通过，把该用户的余额赋值给当前 User 实例
			u.Credit = user.Credit
			loginSuccess = true
			break
		}
	}

	// 输出校验结果
	if loginSuccess {
		fmt.Printf("登录成功！%s 的账户余额为：%d\n", u.Name, u.Credit)
		return true
	} else {
		fmt.Println("登录失败：用户名或密码错误")
		return false
	}
}
