/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package session

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"tinyorm/dialect"
	"tinyorm/options"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.ProviderDialect(options.Options{
		DialectName: "sqlite3",
	})
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSessionRaw() (*Session, error) {
	return NewSession(TestDB, TestDial)
}

func TestSession_Raw(t *testing.T) {
	s, err := NewSession(TestDB, TestDial)
	if err != nil {
		t.Fatal(err)
	}
	//create table and insert record
	_, err = s.Raw("CREATE TABLE IF NOT EXISTS User(Name text);").Exec()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if err != nil {
		log.Fatal(err)
		return
	}

	//select
	row := s.Raw("SELECT Name FROM User LIMIT 1").QueryRow()
	if err != nil {
		log.Fatal(err)
		return
	}

	var name string
	err = row.Scan(&name)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(name)
}
