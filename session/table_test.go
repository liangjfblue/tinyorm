/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package session

import (
	"log"
	"testing"
)

type TBUser struct {
	Id   string `tinyorm:"PRIMARY KEY"`
	Name string
	Addr string
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	s, err := NewSessionRaw()
	if err != nil {
		t.Fatal(err)
	}

	s.Model(&TBUser{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("create fail")
	}

	_, err = s.Raw("INSERT INTO TBUser(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	row := s.Raw("SELECT Name FROM TBUser LIMIT 1").QueryRow()

	var name string
	err = row.Scan(&name)
	if err != nil {
		log.Fatal(err)
		return
	}

	t.Log(name)
	t.Log("create ok")
}
