package main

import (
	"fmt"
	"github.com/lvxin0315/syrupPlum"
	"log"
)

type Demo struct {
	Name     string
	Parent   string
	Children []*Demo
}

func main() {
	op, err := syrupPlum.InitOption("db/")
	if err != nil {
		log.Fatal(err)
	}
	sp := syrupPlum.NewSyrupPlum(op)

	syrupPlum.SetDebug(true)

	child := &Demo{
		Name:   "yiyi",
		Parent: "lvxin",
	}

	my := &Demo{
		Name:     "lvxin",
		Children: []*Demo{child},
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
