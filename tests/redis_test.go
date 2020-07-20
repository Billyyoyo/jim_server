package tests

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"jim/logic/cache"
	"jim/logic/model"
	"testing"
)

func TestSequence(t *testing.T) {
	seq, err := cache.GetUserMsgSequence(1)
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println(seq)
}

func TestSaveUserConn(t *testing.T) {
	conn1 := &model.UserConn{
		Server:   "localhost:5000",
		Addr:     "127.0.0.1:43334",
		DeviceId: 4,
	}
	cache.SaveUserConn(4, conn1)
}

func TestGetUserConn(t *testing.T) {
	conn := &model.UserConn{}
	err := cache.GetUserConn(1, 1, conn)
	if err != nil {
		log.Error(err.Error())
		return
	}
	print(conn)
}

func TestGetAllUserConn(t *testing.T) {
	conns, err := cache.ListUserConn(1)
	if err != nil {
		log.Error(err.Error())
		return
	}
	print(conns)
}

