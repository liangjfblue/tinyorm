/**
 *
 * @author liangjf
 * @create on 2020/8/24
 * @version 1.0
 */
package session

import "tinyorm/log"

func (s *Session) Begin() (err error) {
	log.Info("tinyorm: transaction begin")
	s.tx, err = s.db.Begin()
	if err != nil {
		log.Error("tinyorm: err:", err.Error())
		return err
	}
	return err
}

func (s *Session) RollBack() (err error) {
	log.Info("tinyorm: transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error("tinyorm: err:", err.Error())
		return err
	}
	return err
}

func (s *Session) Commit() (err error) {
	log.Info("tinyorm: transaction rollback")
	if err = s.tx.Commit(); err != nil {
		log.Error("tinyorm: err:", err.Error())
		return err
	}
	return err
}
