package service

import (
	"../model"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"../common"
)

func TestService() model.ModeTest {
	panic("bad end")
	return model.ModeTest{
		"a",
		"b",
	}
}

func TestAddPerson() string {
	p:= model.TestPerson{
		Name:"aaa",
		Phone:"15000000000",
	}
	p.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := common.WitchCollection("person", query)
	if err != nil {
		return "false"
	}
	return p.Id.Hex()
}

func TestGetOne(id string) *model.TestPerson {
	objid := bson.ObjectIdHex(id)
	person := new(model.TestPerson)
	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&person)
	}
	common.WitchCollection("person", query)
	return person
}

func TestGetAll() []model.TestPerson {
	var persons []model.TestPerson
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"tname": "aaa"}).Skip(1).Limit(1).All(&persons)
	}
	err := common.WitchCollection("person", query)
	if err != nil {
		return persons
	}
	return persons
}