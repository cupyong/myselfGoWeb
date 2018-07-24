package controller

import (
	"../common"
	"../router"
	"../service"
)

func Test(w common.ResponseWriter, r *common.Request) (router.Result, error) {
    data := service.TestService()
	return  data ,nil
}

func TestAddPerson(w common.ResponseWriter, r *common.Request) (router.Result, error) {
	data := service.TestAddPerson()
	return  data ,nil
}

func TestGetOneMgo(w common.ResponseWriter, r *common.Request) (router.Result, error) {
	data := service.TestGetOne("5a964fdd43707a2c58ef7efd")
	return  data ,nil
}

func TestGetAllMgo(w common.ResponseWriter, r *common.Request) (router.Result, error) {
	data := service.TestGetAll()
	return  data ,nil
}