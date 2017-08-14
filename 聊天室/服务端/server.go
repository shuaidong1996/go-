package main

import (
	"fmt"
	"net"
)

func checkError(err error,info string) (res bool){
	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

func main() {
	//service := "127.0.0.1:8888"

	tcpAddr,err :=net.ResolveTCPAddr("tcp",service)
	checkError(err,"ResolveTCPAddr")

	lis, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err,"ListenTCP")

	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)

	//启动服务器广播线程（发送从客户端读到的数据）
	go echoHandler(&conns, messages)

	for{
		conn, err := lis.Accept()
		checkError(err, "Accept")
		fmt.Println("来了一个新客户端，地址为：",conn.RemoteAddr().String())
		conns[conn.RemoteAddr().String()] = conn
		/*启动 读客户端数据 的协程*/
		go Handler(conn, messages)
	}
}
func echoHandler(conns *map[string]net.Conn, messages chan string) {
	for{
		msg := <- messages //如果channel中有数据就继续执行
		/*遍历客户端，把消息发给每一个客户端*/
		for key, value := range *conns{
			_,err := value.Write([]byte(msg))
			fmt.Println("给",key,"发送了：",msg)
			if err != nil{
				fmt.Println(err.Error())
				delete(*conns,key)//根据key将元素从映射中删除
			}
		}
	}
}
func Handler(conn net.Conn, messages chan string) {
	buf := make([]byte,1024)
	/*开始读客户端的数据*/
	for{
		length,err:=conn.Read(buf)
		if checkError(err,"connection") == false{
			conn.Close()//如果读不到数据，就关闭这个连接
			break
		}
		/*如果读到了数据，放到channel中*/
		messages <- string(buf[:length])
	}
}

