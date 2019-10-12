package test

import (
	"fmt"
	"syrup-plum"
	"testing"
	"time"
)

type SpData struct {
	Who   string
	When  time.Time
	Which int64
}

func Test_NewSyrupPlum(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
}

func Test_Save(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	//保存
	spData := new(SpData)
	spData.When = time.Now()
	spData.Which = 999
	spData.Who = "lvxin"
	err := sp.Save("demo", spData)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Find(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	spData := new(SpData)
	err := sp.Find("demo", spData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(spData)
	if spData.Who == "" {
		t.Fatal()
	}
}

func Test_SaveList(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	//保存
	var spDataList []*SpData
	spDataList = append(spDataList, &SpData{
		Who:   "lvxin1",
		Which: 2,
		When:  time.Now(),
	}, &SpData{
		Who:   "lvxin2",
		Which: 3,
		When:  time.Now(),
	})
	err := sp.Save("demoList", &spDataList)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetList(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	var spDataList []*SpData

	err := sp.Find("demoList", &spDataList)
	if err != nil {
		t.Fatal(err)
	}
	if len(spDataList) == 0 {
		t.Fatal("list is empty")
	}
	fmt.Println(len(spDataList))
	fmt.Println(spDataList[0])
}

func Test_Delete(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	//保存
	var spDataList []*SpData
	spDataList = append(spDataList, &SpData{
		Who:   "lvxin1",
		Which: 2,
		When:  time.Now(),
	}, &SpData{
		Who:   "lvxin2",
		Which: 3,
		When:  time.Now(),
	})
	err := sp.Save("demoDelList", &spDataList)
	if err != nil {
		t.Fatal(err)
	}
	err = sp.Delete("demoDelList")
	if err != nil {
		t.Fatal(err)
	}
	var spDataListNew []*SpData
	err = sp.Find("demoDelList", &spDataListNew)
	if err == nil {
		t.Fatal("文件应该被删除")
	}
}

func Test_AppendList(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
	//保存
	var spDataList []*SpData
	spDataList = append(spDataList, &SpData{
		Who:   "lvxin1",
		Which: 2,
		When:  time.Now(),
	}, &SpData{
		Who:   "lvxin2",
		Which: 3,
		When:  time.Now(),
	})
	err := sp.Save("demoList", &spDataList)
	if err != nil {
		t.Fatal(err)
	}
	//追加
	var spAppendDataList []*SpData
	err = sp.Append("demoList", &spAppendDataList, &SpData{
		Who:   "lvxin3",
		Which: 4,
		When:  time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(spAppendDataList) != 3 {
		t.Fatal("追加失败---", len(spAppendDataList))
	}

	//查询持久化内容
	var spAppendEdDataList []*SpData
	err = sp.Find("demoList", &spAppendEdDataList)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(spAppendEdDataList[2])

	if len(spAppendEdDataList) != 3 {
		t.Fatal("追加保存失败")
	}
}
