package syrup_plum

import (
	"errors"
)

var Debug = false

type SyrupPlum struct {
	option *Option
}

func NewSyrupPlum(option *Option) *SyrupPlum {
	if option == nil {
		panic(errors.New("option is error"))
	}
	if _, err := option.CheckHealthy(); err != nil {
		panic(err)
	}
	sp := new(SyrupPlum)
	sp.option = option
	return sp
}

func (sp *SyrupPlum) getQuery() *query {
	q := NewQuery(sp.option.SavePath)
	return q
}

//保存，如果存在会覆盖
func (sp *SyrupPlum) Save(index string, object interface{}) error {
	return sp.getQuery().Save(index, object)
}

//获取
func (sp *SyrupPlum) Find(index string, object interface{}) error {
	return sp.getQuery().Load(index, object)
}

//删除
func (sp *SyrupPlum) Delete(index string) error {
	return sp.getQuery().Remove(index)
}
