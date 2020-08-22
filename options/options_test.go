/**
 *
 * @author liangjf
 * @create on 2020/8/22
 * @version 1.0
 */
package options

import (
	"testing"
)

func TestOptions(t *testing.T) {
	opts := Options{
		DialectName:      "sqlite3",
		DBDriverName:     "sqlite3",
		DBDataSourceName: "gee.db",
	}

	if opts.DialectName != "sqlite3" && opts.DBDriverName != "sqlite3" && opts.DBDataSourceName != "gee.db" {
		t.Fatal("options edit fail")
	}
}
