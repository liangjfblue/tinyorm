/**
 *
 * @author liangjf
 * @create on 2020/8/22
 * @version 1.0
 */
package clause

import (
	"reflect"
	"testing"
)

func Test__select(t *testing.T) {
	var clause Clause
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Age ASC")
	clause.Set(LIMIT, 3)
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)

	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build sql")
	}

	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build sqlVars")
	}
}
