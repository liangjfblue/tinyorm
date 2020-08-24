/**
 *
 * @author liangjf
 * @create on 2020/8/24
 * @version 1.0
 */
package session

import (
	"reflect"
	"tinyorm/log"
)

type HooksBeforeQuery interface {
	BeforeQuery(s *Session) error
}

type HooksAfterQuery interface {
	AfterQuery(s *Session) error
}
type HooksBeforeUpdate interface {
	BeforeUpdate(s *Session) error
}
type HooksAfterUpdate interface {
	AfterUpdate(s *Session) error
}
type HooksBeforeDelete interface {
	BeforeDelete(s *Session) error
}
type HooksAfterDelete interface {
	AfterDelete(s *Session) error
}

type HooksBeforeInsert interface {
	BeforeInsert(s *Session) error
}
type HooksAfterInsert interface {
	AfterInsert(s *Session) error
}

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.refTable.Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		ret := fm.Call(param)
		if len(ret) > 0 {
			if err, ok := ret[0].Interface().(error); ok {
				log.Error("tinyorm: call method err: ", err)
			}
		}
	}

}
