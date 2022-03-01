package main

import (
	"database.dictionary/dao"
	"database.dictionary/routes"
	"fmt"
)

func main() {
	err := dao.InitMySQL()

	if err != nil {
		fmt.Printf("mysql init failed, err:%v\n", err)
		return
	}

	defer func() {
		err := dao.Close()
		if err != nil {
			fmt.Printf("mysql close failed, err:%v\n", err)
		}
	}()

	r := routers.SetupRouter()

	if err := r.Run(":9000"); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
	}
}
