/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package db

import (
	"database/sql"
	"tinyorm/options"
)

func ProviderDB(Opts options.Options) (*sql.DB, error) {
	return sql.Open(Opts.DBDriverName, Opts.DBDataSourceName)
}
