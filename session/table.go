/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package session

import (
	"fmt"
	"reflect"
	"strings"
	"tinyorm/log"
	"tinyorm/schema"
)

//操作数据库表

func (s *Session) Model(value interface{}) *Session {
	//只有映射对象为空或者已被更新才映射
	if s.refTable == nil || reflect.ValueOf(s.refTable.Model) != reflect.ValueOf(value) {
		s.refTable = schema.Parse(value, s.dialect)
	}

	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("tinyorm: model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() (err error) {
	table := s.RefTable()

	columns := make([]string, 0)
	for _, field := range table.Fields {
		column := fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag)
		columns = append(columns, column)
	}

	desc := strings.Join(columns, ",")
	sql := fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)
	_, err = s.Raw(sql).Exec()
	return
}

func (s *Session) DropTable() (err error) {
	_, err = s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSql(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()

	var table string
	_ = row.Scan(&table)

	return table == s.RefTable().Name
}
