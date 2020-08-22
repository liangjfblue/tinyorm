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
	"tinyorm/options"
)

//Dialect go type to sql type
type Dialect interface {
	DataTypeOf(t reflect.Value) string
	TableExistSql(tableName string) (string, []interface{})
}

var (
	_dialectM       = map[string]Dialect{}
	_defaultDialect = &sqlite3{}
)

func init() {
	RegisterDialect("sqlite3", _defaultDialect)
}

func RegisterDialect(name string, dialect Dialect) {
	_dialectM[name] = dialect
}

func ProviderDialect(Opts options.Options) (Dialect, error) {
	d, ok := _dialectM[Opts.DialectName]
	if !ok {
		return nil, fmt.Errorf("not exist %s", Opts.DialectName)
	}
	return d, nil
}
