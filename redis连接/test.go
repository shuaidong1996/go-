package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func CheckErr(err error,info string) bool{
	if err!=nil{
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}
/*存取字符串*/
func redis_string() {
	// tcp连接redis
	rs, err := redis.Dial("tcp", "127.0.0.1:6379")
	CheckErr(err,"Dial")
	// 操作完后自动关闭
	defer rs.Close()

	//执行命令
	res,err := rs.Do("set","name", "dong")
	CheckErr(err, "set")
	fmt.Println(res)

	//第一种方法获取结果
	res,err = rs.Do("get", "name")
	CheckErr(err,"get")
	fmt.Println(res)
	fmt.Println(string(res.([]byte)))

	//第二种方法
	v,err := redis.String(rs.Do("get", "name"))
	fmt.Println(v)
}
/*存取集合*/
func redis_list() {
	rs,err := redis.Dial("tcp","localhost:6379")
	CheckErr(err, "dial")

	rs.Do("lpush","redisList","aaa")
	rs.Do("lpush","redisList","bbb")
	rs.Do("rpush","redisList","ccc")

	values,err := redis.Values(rs.Do("lrange","redisList","0","-1"))
	CheckErr(err, "lrange")
	//for遍历
	for _,val := range values{
		fmt.Print(string(val.([]byte)),"  ")
	}
	fmt.Println("")

	//scan函数
	var v1,v2 string
	redis.Scan(values, &v1,&v2)
	fmt.Println(v1,v2)
	rs.Do("del","redisList")
}

/*管道*/
func redis_channel() {
	rs,err := redis.Dial("tcp", "localhost:6379")
	CheckErr(err, "dial")
	rs.Send("set","name", "shuai")	//Send向连接的输出缓冲中写入命令
	rs.Send("get", "name")
	rs.Flush()	//Flush将连接的输出缓冲清空并写入服务器端
	v,err := rs.Receive()	// reply from SET
	fmt.Println(v)
	v,err = rs.Receive()	// reply from GET
	fmt.Println(string(v.([]byte)))
}

/*发布订阅*/
func subscribe(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	CheckErr(err,"dial")
	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe("ChatRoom")	//订阅CharRoom
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

/*发布者*/
func publish(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	CheckErr(err,"dial")
	defer c.Close()

	for {
		var s string
		fmt.Scanln(&s)
		_, err := c.Do("PUBLISH", "ChatRoom", s)
		CheckErr(err,"do")
	}
}

func main() {
	go subscribe()
	go subscribe()
	go subscribe()
	go subscribe()
	go subscribe()
	publish()
}
