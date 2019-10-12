package syrup_plum

import (
	"errors"
	"reflect"
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

//追加保存，会覆盖
func (sp *SyrupPlum) Append(index string, object interface{}, appendObject interface{}) error {
	if err := sp.Find(index, object); err != nil {
		return err
	}
	values := reflect.ValueOf(object).Elem()
	appendValues := reflect.ValueOf(appendObject)
	valArr := reflect.Append(values, appendValues)
	values.Set(valArr)
	return sp.getQuery().Save(index, object)
}
