package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type TiDBStatus string

const (
	TiDBInited    TiDBStatus = "Inited"
	TiDBRunning              = "Runing"
	TiDBStoped               = "Stoped"
	TiDBUpgrading            = "Upgrading"
)

type TiDBCluster struct {
	ID          int64     `json:"id" xorm:"id"`
	Name        string    `json:"name" xorm:"VARCHAR(200) UNIQUE NOT NULL"`
	Version     string    `json:"version" xorm:"VARCHAR(200)"`
	Path        string    `json:"path" xorm:"VARCHAR(200)"`
	Host        string    `json:"host" xorm:"VARCHAR(200)"`
	Status      string    `json:"status" xorm:"VARCHAR(200)"`
	Description string    `json:"description" xorm:"VARCHAR(512)"`
	InitTime    time.Time `json:"init_time" xorm:"init_time"`
}

func CreateTiDBCluster(tc *TiDBCluster) error {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}

	isExist, err := isTiDBClusterExist(sess, 0, tc.Name)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("%s tidb cluster already exists", tc.Name)
	}

	tc.Host = strings.ToLower(tc.Host)
	tc.Path = strings.ToLower(tc.Path)
	isExist, err = sess.
		Where("host=? and path=?", tc.Host, tc.Path).
		Get(new(TiDBCluster))
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("%s:%s tidb cluster already exists", tc.Host, tc.Path)
	}

	if _, err := sess.Insert(tc); err != nil {
		return err
	}

	return sess.Commit()
}

func GetTiDBCluster(tc *TiDBCluster) (bool, error) {
	return x.Get(tc)
}

func GetTiDBClusterByName(name string) (*TiDBCluster, error) {
	return getTiDBClusterByName(x, name)
}

func getTiDBClusterByName(e Engine, name string) (*TiDBCluster, error) {
	if name == "" {
		return nil, errors.New("")
	}

	tc := &TiDBCluster{Name: name}
	has, err := e.Get(tc)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("tidb cluster %s not exist", name)
	}

	return tc, nil
}

func LoadTiDBClusters() ([]*TiDBCluster, error) {
	return loadTiDBClusters(x)
}

func loadTiDBClusters(e Engine) ([]*TiDBCluster, error) {
	tcs := make([]*TiDBCluster, 0, 10)
	if err := e.
		Where("1=1").
		OrderBy("init_time").
		Find(&tcs); err != nil {
		return nil, err
	}

	return tcs, nil
}

func GetTiDBClusterByHost(host string) ([]*TiDBCluster, error) {
	return getTiDBClusterByHost(x, host)
}

func getTiDBClusterByHost(e Engine, host string) ([]*TiDBCluster, error) {
	tcs := make([]*TiDBCluster, 0, 10)

	if err := x.
		Where("host=?", host).
		OrderBy("init_time").
		Find(&tcs); err != nil {
		return nil, err
	}

	return tcs, nil
}

func isTiDBClusterExist(e Engine, uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return e.
		Where("id!=?", uid).
		Get(&TiDBCluster{Name: strings.ToLower(name)})
}

func UpdateTiDBCluster(tc *TiDBCluster) error {
	return updateUser(x, tc)
}

func updateUser(e Engine, tc *TiDBCluster) error {
	_, err := e.ID(tc.ID).Update(tc)
	return err
}
