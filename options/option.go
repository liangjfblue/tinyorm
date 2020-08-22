/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package options

type Option func(*Options)

var (
	DefaultOptions = Options{
		DialectName:      "sqlite3",
		DBDriverName:     "sqlite3",
		DBDataSourceName: "gee.db",
	}
)
