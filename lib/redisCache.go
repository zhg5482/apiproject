package lib

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis" //redis
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	//定义常量
	RedisClient *redis.Pool
	REDIS_HOST string
	REDIS_DB int
)

//初始化连接池
func init()  {
	//从配置文件获取redis的ip和db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

//get key
func RedisGet(key string) (string,error){
	// 从池里获取连接
	rc := RedisClient.Get()
	//====================
	val,err :=  rc.Do("GET", key)
	if err != nil{
		fmt.Println("redis get key->value failed!")
	}
	res := string(val.([]byte))
	//====================
	// 用完后将连接放回连接池
	defer rc.Close()
	return res,err
}

//set key val expire
func RedisSet(key string,val string,expire time.Duration) bool {
	// 从池里获取连接
	rc := RedisClient.Get()
	//====================
	flag := true
	n,err := rc.Do("set", key, val)
	if err != nil {
		fmt.Println("redis set key->value failed!",err.Error())
		flag = false
	}else{
		if n == int64(1) {	//设置过期时间
			n,_ := rc.Do("EXPIRE", key, expire)
			if n == int64(1) {
				flag = true
				fmt.Println("redis set key->value success!")
			}
		}else if n == int64(0) {
			fmt.Println("redis set key->value has already existed!",err.Error())
			flag = false
		}
	}
	//====================
	// 用完后将连接放回连接池
	defer rc.Close()
	return flag
}