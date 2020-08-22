/**
 *
 * @author liangjf
 * @create on 2020/8/22
 * @version 1.0
 */
package db

import (
	"testing"
	"tinyorm/options"

	_ "github.com/mattn/go-sqlite3"
)

func TestProviderDB(t *testing.T) {
	_, err := ProviderDB(options.Options{
		DialectName:      "sqlite3",
		DBDriverName:     "sqlite3",
		DBDataSourceName: "gee.db",
	})
	if err != nil {
		t.Fatal(err)
	}
}
