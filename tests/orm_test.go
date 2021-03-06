package tests

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"jim/common/rpc"
	"jim/common/utils"
	"jim/logic/dao"
	"jim/logic/model"
	"testing"
)

func TestGetUser(t *testing.T) {
	user, err := dao.GetUser(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printj(user)
}

func TestGetSessions(t *testing.T) {
	sessions, err := dao.GetSessionByUser(1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printj(sessions)
}

func TestGetMembers(t *testing.T) {
	members, err := dao.GetMemberInSession(10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printj(members)
}

func TestAddMember(t *testing.T) {
	var member = model.Member{
		SessionId:  1,
		UserId:     2,
		CreateTime: utils.GetCurrentMS(),
	}
	err := dao.AddMemberV2(dao.DB(), member)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetDevice(t *testing.T) {
	_, device, err := dao.GetDevice(1, "123adskf23ek2jrh")
	if err != nil {
		fmt.Println(err)
		return
	}
	printj(device)
}

func TestRecordDevice(t *testing.T) {
	device := model.Device{
		UserId:         1,
		SerialNo:       "123adskf23ek2jrh",
		LastAddress:    "10.8.240.133:42123",
		LastConnTime:   utils.GetCurrentMS(),
		CreateTime:     utils.GetCurrentMS(),
		LastDisconTime: 0,
	}
	err := dao.SaveDevice(device)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCreateSession(t *testing.T) {
	session := model.Session{
		Name:       "Go1 space",
		Type:       2,
		Owner:      1,
		CreateTime: utils.GetCurrentMS(),
	}
	err := dao.CreateSession(dao.DB(), session)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printj(session)
}

func TestAddAck(t *testing.T) {
	ack := model.Ack{
		MsgId:       1,
		SendCount:   1,
		ArriveCount: 0,
	}
	err := dao.AddAck(ack)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	printj(ack)
}

func TestAddMessage(t *testing.T) {
	body := map[string]interface{}{
		"url":      "http://www.sina.com.cn",
		"duration": "10",
	}
	bs, err := json.Marshal(body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	msg := model.Message{
		SenderId:   1,
		SessionId:  9,
		Type:       2,
		Status:     1,
		Sequence:   3,
		Body:       bs,
		CreateTime: utils.GetCurrentMS(),
	}
	dao.AddMessage(msg)
	printj(msg)
}

func TestDelMember(t *testing.T) {
	err := dao.DeleteMember(1, 3)
	if err != nil {
		log.Println(err.Error())
	}
}

func TestRenameSession(t *testing.T) {
	err := dao.RenameSession(1, "Happy Coder")
	if err != nil {
		log.Println(err.Error())
	}
}

func TestWithdrawMessage(t *testing.T) {
	affect, err := dao.WithdrawMessage(1, 1, 1)
	if err != nil {
		log.Error(err.Error())
		return
	}
	printl("update ", affect, " rows")
}

func TestAccumulateSendCount(t *testing.T) {
	err := dao.AccumulateAckSendCount(1)
	if err != nil {
		printl(err.Error())
	}
}

func TestAccumulateArriveCount(t *testing.T) {
	err := dao.AccumulateAckArriveCount(1)
	if err != nil {
		printl(err.Error())
	}
}

func TestGetOfflineMessages(t *testing.T) {
	msgs, err := dao.GetOfflineMsgs(18)
	if err != nil {
		printl(err.Error())
		return
	}
	printj(msgs)
}

func TestGetMessageList(t *testing.T) {
	msgs, err := dao.GetMessagesSeqAfter(9, 30, 10)
	if err != nil {
		log.Error(err.Error())
		return
	}
	for _, msg := range msgs {
		body := &rpc.Words{}
		err := proto.Unmarshal(msg.Body, body)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		printj(msg)
		printj(body)
	}
}
