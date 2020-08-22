/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package options

type Options struct {
	DialectName      string
	DBDriverName     string
	DBDataSourceName string
}

func DialectName(dialectName string) Option {
	return func(o *Options) {
		o.DialectName = dialectName
	}
}

func DBDriverName(dBDriverName string) Option {
	return func(o *Options) {
		o.DBDriverName = dBDriverName
	}
}

func DBDataSourceName(dBDataSourceName string) Option {
	return func(o *Options) {
		o.DBDataSourceName = dBDataSourceName
	}
}
