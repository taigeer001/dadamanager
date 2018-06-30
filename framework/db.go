package xf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"errors"
)

type connnectinf struct {
	t string
	str string
	db *gorm.DB
}

type _Connector struct{
	list map[string]connnectinf
}

func (this *_Connector) GetConnect(name string) *connnectinf {
	if v, ok := this.list[name]; ok {
		return &v
	}
	return nil
}

var tmp *_Connector

func NewConnnect(inf connnectinf) *gorm.DB {
	db, err := gorm.Open(inf.t, inf.str)
	if err != nil {
		return nil
	}
	return db
}

func BindConnector(name string, t string, str string)  {
	if tmp == nil {
		tmp = &_Connector{list:make(map[string]connnectinf)}
	}
	tmp.list[name] = connnectinf{t:t, str:str}
}

func GetConnect(name string) (*gorm.DB, error) {
	if tmp == nil {
		tmp = &_Connector{list:make(map[string]connnectinf)}
	}
	iinf := tmp.GetConnect(name)
	if iinf == nil {
		return nil, errors.New("connect error")
	}
	inf := *iinf
	if inf.db == nil {
		dbtmp := NewConnnect(inf)
		if dbtmp == nil {
			return nil, errors.New("connect error")
		}
		return dbtmp, nil
	}
	return inf.db, nil
}
