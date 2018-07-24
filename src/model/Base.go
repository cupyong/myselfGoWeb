package model

//import "fmt"

type Base struct {
	Code    int `json:"code"`
	Data    interface{} `json:"data"`
	Message string `json:"message"`
}

func (a *Base) SetCode(m int) {
	a.Code=m;
}

func test()  {
	a:=&Base{}
	a.SetCode(1)
}


