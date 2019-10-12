package main

import (
	"fmt"
	"github.com/lvxin0315/syrup-plum"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type DemoW struct {
	Name     string
	Parent   string
	Num      int
	Children []*DemoW
}

func main() {
	serve()
	op, err := syrup_plum.InitOption("db/")
	if err != nil {
		log.Fatal(err)
	}
	sp := syrup_plum.NewSyrupPlum(op)

	syrup_plum.SetDebug(true)

	//生成10万个child
	var children []*DemoW
	for i := 0; i < 100000; i++ {
		children = append(children, &DemoW{
			Name:   "yiyi",
			Parent: "lvxin",
			Num:    i,
		})
	}

	//将10万个child付给100个parent
	var parents []*DemoW
	for i := 0; i < 1000; i++ {
		parents = append(parents, &DemoW{
			Name:     "yiyi",
			Num:      i,
			Children: children,
		})
	}

	err = sp.Save("demo10W", parents)
	if err != nil {
		log.Fatal(err)
	}

	var dbParents []*DemoW

	err = sp.Find("demo10W", &dbParents)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range dbParents {
		fmt.Println(v.Num, ":", len(v.Children))
	}
}

func serve() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()
}
