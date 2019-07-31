package syrup_plum

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"github.com/pkg/errors"
	"os"
)

type query struct {
	dbPath string   //db目录
	sha1   string   //sha1对应保存文件名称
	file   *os.File //文件操作
	index  string   //使用key
}

func NewQuery(dbPath string) *query {
	q := new(query)
	q.dbPath = dbPath
	return q
}

func (q *query) open() error {
	if !q.isExist() {
		return errors.New("open: file is not exist")
	}
	file, err := os.Open(q.dbPath + q.sha1)
	if err != nil {
		SPError("open:", err)
		return err
	}
	q.file = file
	return nil
}

func (q *query) create() error {
	file, err := os.Create(q.dbPath + q.sha1)
	if err != nil {
		SPError("create:", err)
		return err
	}
	q.file = file
	return nil
}

func (q *query) isExist() bool {
	_, err := PathExists(q.dbPath + q.sha1)
	if err != nil {
		return false
	}
	return true
}

func (q *query) close() error {
	if q.file == nil {
		err := errors.New("query:file close error")
		SPError("close:", err)
		return err
	}
	return q.file.Close()
}

func (q *query) rm() error {
	err := os.Remove(q.dbPath + q.sha1)
	if err != nil {
		SPError(err)
		return err
	}
	return nil
}

func (q *query) setIndex(index string) error {
	h := sha1.New()
	_, err := h.Write([]byte(index))
	if err != nil {
		SPError("setIndex:", err)
		return err
	}
	q.index = index
	q.sha1 = hex.EncodeToString(h.Sum(nil))
	return nil
}

//内容保存
func (q *query) Save(index string, object interface{}) error {
	if err := q.setIndex(index); err != nil {
		return err
	}
	if err := q.create(); err != nil {
		return err
	}
	defer q.close()
	encoder := gob.NewEncoder(q.file)
	if err := encoder.Encode(object); err != nil {
		SPError("encoder.Encode:", err)
		return err
	}
	return nil
}

//读取内容
func (q *query) Load(index string, object interface{}) error {
	if err := q.setIndex(index); err != nil {
		return err
	}
	if err := q.open(); err != nil {
		return err
	}
	defer q.close()
	decoder := gob.NewDecoder(q.file)
	if err := decoder.Decode(object); err != nil {
		SPError("decoder.Decode:", err)
		return err
	}
	return nil
}

func (q *query) Remove(index string) error {
	if !q.isExist() {
		return errors.New("Remove:file is not exist")
	}
	if err := q.setIndex(index); err != nil {
		return err
	}
	return q.rm()
}
