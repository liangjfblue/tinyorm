/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package tinyorm

import (
	"fmt"
	"log"
	"testing"
	"tinyorm/options"
)

func TestNewTinyOrm(t *testing.T) {
	o, err := NewTinyOrm(options.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	defer o.Close()

	_, err = o.session.Raw("DROP TABLE IF EXISTS User;").Exec()
	if err != nil {
		log.Fatal(err)
	}

	_, err = o.session.Raw("CREATE TABLE IF NOT EXISTS User(Name text);").Exec()
	if err != nil {
		log.Fatal(err)
	}

	//test create table exists
	_, err = o.session.Raw("CREATE TABLE IF NOT EXISTS User(Name text);").Exec()
	if err != nil {
		log.Fatal(err)
	}

	_, err = o.session.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	row := o.session.Raw("SELECT Name FROM User LIMIT 1").QueryRow()

	var name string
	err = row.Scan(&name)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(name)
}

type User struct {
	Name string
	Age  int
}

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func TestInsert(t *testing.T) {
	o, err := NewTinyOrm(options.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	defer o.Close()

	s, err := o.NewSession()
	if err != nil {
		t.Fatal(err)
	}

	s.Model(&User{})

	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
}
