/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package session

import (
	"database/sql"
	"strings"
	"tinyorm/clause"
	"tinyorm/dialect"
	"tinyorm/log"
	"tinyorm/schema"

	_ "github.com/mattn/go-sqlite3"
)

type Session struct {
	db       *sql.DB
	tx       *sql.Tx
	refTable *schema.Schema
	dialect  dialect.Dialect
	clause   clause.Clause
	sql      strings.Builder
	sqlVars  []interface{}
}

//检查是否实现了CommonDB
var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func NewSession(db *sql.DB, d dialect.Dialect) (*Session, error) {
	return &Session{
		db:      db,
		dialect: d,
	}, nil
}

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info("tinyorm: ", s.sql.String(), s.sqlVars)
	result, err = s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error("tinyorm: ", err.Error())
	}
	return
}

func (s *Session) QueryRow() (row *sql.Row) {
	defer s.Clear()
	log.Info("tinyorm: ", s.sql.String(), s.sqlVars)
	row = s.DB().QueryRow(s.sql.String(), s.sqlVars...)
	return
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info("tinyorm: ", s.sql.String(), s.sqlVars)
	rows, err = s.DB().Query(s.sql.String(), s.sqlVars...)
	return
}
