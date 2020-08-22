/**
 *
 * @author liangjf
 * @create on 2020/8/22
 * @version 1.0
 */
package dialect

import (
	"reflect"
	"testing"
)

func Test_sqlite3_DataTypeOf(t *testing.T) {
	s := sqlite3{}

	str := "123"
	if s.DataTypeOf(reflect.ValueOf(str)) != "text" {
		t.Fatal("go get sql data type error")
	}
}
