package syrup_plum

import "errors"

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
