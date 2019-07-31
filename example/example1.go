package main

import (
	"fmt"
	"github.com/lvxin0315/syrup-plum"
	"log"
)

type Demo struct {
	Name     string
	Parent   string
	Children []*Demo
}

func main() {
	op, err := syrup_plum.InitOption("./example/db/")
	if err != nil {
		log.Fatal(err)
	}
	sp := syrup_plum.NewSyrupPlum(op)

	syrup_plum.SetDebug(true)

	children := &Demo{
		Name:   "yiyi",
		Parent: "lvxin",
	}

	my := &Demo{
		Name:     "lvxin",
		Children: []*Demo{children},
	}

	err = sp.Save("abc", my)
	if err != nil {
		log.Fatal(err)
	}

	newMy := new(Demo)

	err = sp.Find("abc", newMy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newMy)
	fmt.Println(newMy.Children[0].Name)
}
