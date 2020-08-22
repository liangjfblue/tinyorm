/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package schema

import (
	"testing"
	"tinyorm/dialect"
	"tinyorm/options"
)

type TBUser struct {
	Username string `tinyorm:"PRIMARY KEY"`
	Password string
}

func (T *TBUser) TableName() string {
	return "tb_user"
}

func TestParse(t *testing.T) {
	user := TBUser{
		Username: "liangjf",
		Password: "123",
	}

	d, err := dialect.ProviderDialect(options.Options{
		DialectName: "sqlite3",
	})
	if err != nil {
		t.Fatal("not exist sqlite3 dialect")
		return
	}

	s := Parse(&user, d)
	t.Log(s)

	if s.Name != "TBUser" || len(s.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}

	if s.GetField("Username").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
