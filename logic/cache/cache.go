package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"jim/logic/core"
	_ "jim/logic/core"
	"jim/logic/model"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     core.AppConfig.Redis.Addr,
		Password: "",
		DB:       0,
	})
}

func GetUserMsgSequence(userId int64) (no int64, err error) {
	key := fmt.Sprintf("%s_user_seq_%d", core.AppConfig.Redis.Prefix, userId)
	cmd := client.Incr(context.Background(), key)
	if cmd.Err() != nil {
		err = cmd.Err()
		return
	}
	no = cmd.Val()
	return
}

func GetUserConn(userId int64, addr string, userConn *model.UserConn) error {
	key := fmt.Sprintf("%s_user_conn_%d", core.AppConfig.Redis.Prefix, userId)
	cmd := client.HGet(context.Background(), key, addr)
	return cmd.Scan(userConn)
}

func SaveUserConn(userId int64, userConn *model.UserConn) error {
	key := fmt.Sprintf("%s_user_conn_%d", core.AppConfig.Redis.Prefix, userId)
	cmd := client.HSet(context.Background(), key, userConn.Addr, userConn)
	return cmd.Err()
}

func RemoveUserConn(userId int64, addr string) error {
	key := fmt.Sprintf("%s_user_conn_%d", core.AppConfig.Redis.Prefix, userId)
	cmd := client.HDel(context.Background(), key, addr)
	return cmd.Err()
}

func ListUserConn(userId int64) (conns *map[string]model.UserConn, err error) {
	key := fmt.Sprintf("%s_user_conn_%d", core.AppConfig.Redis.Prefix, userId)
	cmd := client.HGetAll(context.Background(), key)
	m, err := cmd.Result()
	if err != nil {
		return
	}
	connd := map[string]model.UserConn{}
	for k, v := range m {
		conn := model.UserConn{}
		conn.UnmarshalBinary([]byte(v))
		connd[k] = conn
	}
	conns = &connd
	return
}
