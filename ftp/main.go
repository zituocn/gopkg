package main

import (
	"fmt"
	ftp "github.com/jlaffaye/ftp"
)

func main() {
	c, err := ftp.Connect("42.96.206.158:21")
	if err != nil {
		fmt.Println(err)
	}

	err = c.Login("samsong", "tco99312^")
	if err != nil {
		fmt.Println(err)
	}

	err = c.ChangeDir("/www")
	if err != nil {
		fmt.Println(err)
	}

	dir, err := c.CurrentDir()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("当前目录：" + dir)
	}

	list, err := c.List("/www")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("目录列表：")
	fmt.Println("=====================================================================")
	for _, info := range list {
		fmt.Println(info.Name)
	}
	c.Quit()

}
