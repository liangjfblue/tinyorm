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
	"fmt"
	"strings"
	"time"
	"tinyorm/dialect"
	"tinyorm/internal/utils"
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

type TxFunc func(s *session.Session) (interface{}, error)

func (o *Engine) Transaction(f TxFunc) (interface{}, error) {
	var err error

	s, err := o.NewSession()
	if err != nil {
		return nil, err
	}

	if err = s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = s.RollBack()
			panic(r)
		} else if err != nil {
			_ = s.RollBack()
			log.Error("tinyorm: rollback err:", err.Error())
		} else {
			err = s.Commit()
		}
	}()

	return f(s)
}

func (o *Engine) Migrate(value interface{}) error {
	if _, err := o.Transaction(func(s *session.Session) (interface{}, error) {
		if !s.Model(value).HasTable() {
			log.Info("tinyorm: create new table now")
			return nil, s.CreateTable()
		}

		//差集
		table := s.Model(value).RefTable()
		rows, err := s.Raw(fmt.Sprintf("SELECT * FROM %s", table.Name)).QueryRows()
		if err != nil {
			log.Error("tinyorm: select rows err:", err.Error())
			return nil, err
		}

		oldColumns, err := rows.Columns()
		if err != nil {
			log.Error("tinyorm: get columns err:", err.Error())
			return nil, err
		}

		addColumns := utils.DiffSlice(oldColumns, table.FieldNames)
		delColumns := utils.DiffSlice(table.FieldNames, oldColumns)

		//新增字段
		/**
		alter table add COLUMN xxx
		*/
		for _, column := range addColumns {
			field := table.GetField(column)
			if _, err = s.Raw(fmt.Sprintf(
				"ALTER TABLE %s add COLUMN %s %s;",
				table.Name,
				field.Name,
				field.Tag,
			)).Exec(); err != nil {
				log.Error("tinyorm: alter table add column err:", err.Error())
				return nil, err
			}
		}

		if len(delColumns) == 0 {
			return nil, err
		}

		tmpTableName := "tmp_" + table.Name
		fieldsStr := strings.Join(table.FieldNames, ",")
		s.Raw(fmt.Sprintf(
			"DROP TABLE IF EXISTS %s;",
			tmpTableName,
		))

		s.Raw(fmt.Sprintf(
			"CREATE TABLE %s AS SELECT %s FROM %s;",
			tmpTableName,
			fieldsStr,
			table.Name,
		))

		s.Raw(fmt.Sprintf(
			"DROP TABLE %s;",
			table.Name,
		))

		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmpTableName, table.Name))

		_, err = s.Exec()

		return nil, err
	}); err != nil {
		return err
	}
	return nil
}

func (o *Engine) NewSession() (*session.Session, error) {
	return session.NewSession(o.db, o.dialect)
}
