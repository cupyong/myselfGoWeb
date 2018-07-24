package router

import (
    "../common"
)

type Result interface{}

type HandlerFunc func(common.ResponseWriter, *common.Request) (Result,error)

type Route struct {
    HttpMethod string
    PathExp    string
    httpOption HttpOption
    Func      HandlerFunc
    Children  ChildrenFunc
}

type ChildrenFunc func() ([]* Route)
