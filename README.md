# syrupPlum
_使用gob对struct等golang数据的存储与管理_

- InitOption 通过参数生成option
- InitOptionWithConfigFile 通过配置文件生成option
- NewSyrupPlum
- Save 保存，如果存在会覆盖
- Append 追加保存，如果存在会覆盖
- Find 获取内容
- Delete 删除

```go
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
```

```go
package main

import (
	"fmt"
	"github.com/lvxin0315/syrupPlum"
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
	op, err := syrupPlum.InitOption("db/")
	if err != nil {
		log.Fatal(err)
	}
	sp := syrupPlum.NewSyrupPlum(op)

	syrupPlum.SetDebug(true)

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

```

