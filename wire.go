//+build wireinject

/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package tinyorm

import (
	"tinyorm/dialect"
	"tinyorm/internal/db"
	"tinyorm/options"
	"tinyorm/session"

	"github.com/google/wire"
)

//func InitializeTinyOrm(Opts *options.Options) (*Engine, error) {
//	panic(wire.Build(
//		db.ProviderDB,
//		dialect.ProviderDialect,
//
//		session.NewSession,
//
//		wire.Struct(new(Engine), "*"),
//	))
//}

func NewTinyOrm(Opts options.Options) (*Engine, error) {
	panic(wire.Build(
		db.ProviderDB,
		dialect.ProviderDialect,

		session.NewSession,

		wire.Struct(new(Engine), "*"),
	))
}
