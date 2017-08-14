package main

import (
	"fmt"
	"net"
	"os"
)
func checkError(err error,info string) (res bool){
	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

func main() {
	service := "127.0.0.1:8888"
	tcpAddr,err := net.ResolveTCPAddr("tcp", service)
	checkError(err, "ResolveTCPAddr")

	conn,err := net.DialTCP("tcp",nil,tcpAddr)//DialTCP在网络协议net上连接本地地址laddr和远端地址raddr,nil表示自动选择一个本地地址

	//向服务器发送消息
	go chatSend(conn)

	//读服务器发过来的数据
	buf := make([]byte,1024)
	for{
		length,err := conn.Read(buf)
		if checkError(err,"Connection") == false{
			conn.Close()
			fmt.Println("服务器挂掉")
			os.Exit(0)
		}
		fmt.Println(string(buf[:length]))
	}

}
func chatSend(conn *net.TCPConn) {
	var input string
	username := conn.LocalAddr().String()//获取本地ip

	for{
		/*读键盘输入的数据*/
		fmt.Scanln(&input)
		if input == "quit" {
			conn.Close()
			os.Exit(0)
		}
		/*发给服务器*/
		_,err := conn.Write([]byte(username+":::"+ input))
		if err != nil{
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}

}
