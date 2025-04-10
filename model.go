package main

import "fmt"

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

func Model() {
	GL_DB.AutoMigrate(&Blog{}, &Author{})
	GL_DB.Create(&Blog{Author: Author{Name: "zhangsan123", Email: "123@163.com"}, Upvotes: 100})
	fmt.Println("插入数据成功")
}
