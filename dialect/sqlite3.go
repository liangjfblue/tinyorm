/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

//利用类型实现接口, 编译期检验类型必须实现接口
var _ Dialect = (*sqlite3)(nil)

func (s sqlite3) DataTypeOf(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Bool:
		return "bool"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := t.Interface().(time.Time); ok {
			return "datetime"
		}
	}

	panic(fmt.Sprintf("tinyorm: invalid sql type %s (%s)", t.Type().Name(), t.Kind()))
}

func (s sqlite3) TableExistSql(tableName string) (string, []interface{}) {
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", []interface{}{tableName}
}
