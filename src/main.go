package main

import (
    "./common"
    "./router"
    "./controller"
    "fmt"
    "os"
)

func main() {
    var configPath string
    if len(os.Args) != 2 {
        fmt.Println("use default config")
        configPath = "./conf/conf.ini"
    } else {
        configPath = os.Args[1]
    }
    common.InitConfig(configPath)

    trie := router.MakeRouter(
        router.Get("/test/:id", controller.Test),//router.HttpOption{} 为缺省项
        router.Get("/test/mgo/add", controller.TestAddPerson, router.HttpOption{}),
        router.Get("/test/mgo/one", controller.TestGetOneMgo, router.HttpOption{}),
        router.Get("/test/mgo/all", controller.TestGetAllMgo, router.HttpOption{Cache:true}),
    )
    router.HttpService(trie)
    common.Service(common.All_port)
}
