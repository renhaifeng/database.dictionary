package main

import (
	"database.dictionary/dao"
	"database.dictionary/routes"
	"database.dictionary/setting"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：xxx conf/config.ini")
		return
	}

	// Load Config
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	err := dao.InitMySQL(setting.Conf.MySQLConfig)

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

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
	}
}
