/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package tinyorm

import (
	"context"
	"database/sql"
	"time"
	"tinyorm/dialect"
	"tinyorm/log"
	"tinyorm/options"
	"tinyorm/session"
)

type Engine struct {
	Opts    options.Options
	db      *sql.DB
	dialect dialect.Dialect
	session *session.Session
}

func (o *Engine) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err := o.db.PingContext(ctx)
	if err != nil {
		log.Error("tinyorm: ", err.Error())
		return err
	}

	log.Info("tinyorm: ping db ok")
	return nil
}

func (o *Engine) Close() (err error) {
	log.Info("tinyorm: tinyorm close")
	if err = o.db.Close(); err != nil {
		log.Error("tinyorm: ", err.Error())
		return
	}
	return
}

func (o *Engine) NewSession() (*session.Session, error) {
	return session.NewSession(o.db, o.dialect)
}
