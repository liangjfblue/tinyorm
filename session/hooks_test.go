/**
 *
 * @author liangjf
 * @create on 2020/8/24
 * @version 1.0
 */
package session

import (
	"testing"
	"tinyorm/log"
)

type Account struct {
	ID       int `tinyorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before inert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s, err := NewSessionRaw()
	if err != nil {
		t.Fatal(err)
	}

	s.Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}

	err = s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Log(err)
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
