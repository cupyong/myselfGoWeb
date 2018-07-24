package model

import "gopkg.in/mgo.v2/bson"

type ModeTest struct {
    Code string `json:"code"`
    Name string `json:"name"`
}


type ExportModel struct {
    HttpType string `json:"httptype"`
    Data string `json:"data"`
}


type TestPerson struct {
    Id    bson.ObjectId `json:"_id" bson:"_id"`
    Name  string        `json:"tname" bson:"tname"` //bson:"name" 表示mongodb数据库中对应的字段名称
    Phone string        `bson:"tphone" bson:"tphone"`
}
