package handler

import (
	log "github.com/sirupsen/logrus"
	"jim/common/rpc"
	"jim/common/utils"
	"jim/logic/cache"
	"jim/logic/dao"
	"jim/logic/model"
)

func Register(userId int64, token, addr, server, serialNo string) (deviceId, lastSequence int64, err error) {
	// 检查用户设备是否入库
	existInDB, device, err := dao.GetDevice(userId, serialNo)
	if err != nil {
		log.Error("register - get device in db fail:", err.Error())
		return
	}
	if !existInDB {
		device = &model.Device{
			UserId:       userId,
			SerialNo:     serialNo,
			LastConnTime: utils.GetCurrentMS(),
			LastAddress:  addr,
			LastSequence: 0,
			CreateTime:   utils.GetCurrentMS(),
		}
	} else {
		device.LastAddress = addr
		device.LastConnTime = utils.GetCurrentMS()
		// 检查是否在线
		existInCache, errr := cache.HasUserConn(userId, device.Id)
		if errr != nil {
			log.Error("register - check conn exist fail:", errr.Error())
			err = errr
			return
		}
		if existInCache {
			// 在线 将在线的连接踢下线
			SendKickoff(server, &rpc.Text{Value: addr})
		}
	}
	err = dao.SaveDevice(device)
	if err != nil {
		log.Error("register - save device in db fail:", err.Error())
		return
	}
	//更新redis中的device
	conn := &model.UserConn{
		Server:   server,
		Addr:     addr,
		DeviceId: device.Id,
	}
	err = cache.SaveUserConn(device.UserId, conn)
	if err != nil {
		log.Error("register - save user conn in redis fail:", err.Error())
		return
	}
	deviceId = device.Id
	// 返回该设备最后消息序号给接入服务器保存
	lastSequence = device.LastSequence
	return
}

func Offline(userId, deviceId, lastSequence int64) (err error) {
	// 删除用户设备的在线状态
	cache.RemoveUserConn(userId, deviceId)
	// 更新数据库表中的用户设备数据
	device, err := dao.GetDeviceById(deviceId)
	if err != nil {
		log.Error("offline - get device in db fail: ", err.Error())
		return
	}
	// 更新最后消息序号
	device.LastSequence = lastSequence
	device.LastDisconTime = utils.GetCurrentMS()
	dao.SaveDevice(device)
	return
}

func GetMembers(sessionId int64) (members *[]model.User, err error) {
	members, err = dao.GetMemberInSession(sessionId)
	if err != nil {
		log.Error("getmembers - get members fail: ", err.Error())
	}
	return
}

func GetSessions(userId int64) (sessions *[]model.Session, err error) {
	sessions, err = dao.GetSessionByUser(userId)
	if err != nil {
		log.Error("getsessions - get sessions fail: ", err.Error())
	}
	return
}

func GetSession(sessionId int64) (session *model.Session, members *[]model.User, err error) {
	session, err = dao.GetSession(sessionId)
	if err != nil {
		log.Error("getsession - get sessions fail: ", err.Error())
	}
	members, err = dao.GetMemberInSession(sessionId)
	if err != nil {
		log.Error("getsession - get members fail: ", err.Error())
	}
	return
}