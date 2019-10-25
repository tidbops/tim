package models

import (
	"database/sql"
	"fmt"

	"xorm.io/core"
	"xorm.io/xorm"

	_ "github.com/mattn/go-sqlite3"
)

// Engine represents a xorm engine or session.
type Engine interface {
	Table(tableNameOrBean interface{}) *xorm.Session
	Count(...interface{}) (int64, error)
	Decr(column string, arg ...interface{}) *xorm.Session
	Delete(interface{}) (int64, error)
	Exec(...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Incr(column string, arg ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *xorm.Session
	SQL(interface{}, ...interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
	Asc(colNames ...string) *xorm.Session
}

var (
	x      *xorm.Engine
	tables []interface{}
	// HasEngine specifies if we have a xorm.Engine
	// HasEngine bool
)

func init() {
	tables = append(tables,
		new(TiDBCluster))
}

func getEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", "./tim.db")
}

// SetEngine sets the xorm.Engine
func SetEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	x.ShowExecTime(true)
	x.SetMapper(core.GonicMapper{})
	x.SetMaxOpenConns(32)
	x.SetMaxIdleConns(8)
	return nil
}

// NewEngine initializes a new xorm.Engine
func NewEngine() (err error) {
	if err = SetEngine(); err != nil {
		return err
	}

	// if err = x.Ping(); err != nil {
	// 	return err
	// }

	if err = x.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("sync database struct error: %v", err)
	}

	return nil
}
