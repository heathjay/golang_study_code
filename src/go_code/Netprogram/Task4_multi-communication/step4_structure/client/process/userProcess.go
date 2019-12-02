package process

/*
	step3: LoginMesResponse
		1. 客户端发送消息
		2.服务器接受消息，然后反序列化成对应的消息结构体
		3.服务器端根据序列化成对应的消息，判断是否登陆用户合法，返回LoginResMes
		4. 客户端解析返回的LoginResMes,显示对应的页面
	3.2
		优化封装：
			处理收发消息
			根据不同消息调用不同函数


*/

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/client/utils"
	"go_code/Netprogram/Task4_multi-communication/step4_structure/common/message"
	"net"
)

type UserProcess struct {
	//字段分析，
	//login需要什么东西
	//暂时不需要任何字段

}

//尽量用error的形式返回，
//error可以返回各种错误，自定义，描述信息
//变量尽量使用一样的
func (this *UserProcess) Login(userId int, userPw string) (err error) {
	//开始定协议
	//fmt.Printf("userId= %d,userPw = %s", userId, userPw)
	//return nil
	//1.连接到服务器
	//理论上说需要读取配置文件
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()
	//2.准备发送信息
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPw = userPw

	//4.marsh
	data, err := json.Marshal(loginMes)
	//mes.Data = loginMes	//
	if err != nil {
		fmt.Println("json error")
		return
	}

	//data 切片形式
	mes.Data = string(data)
	//6.mes 序列化 到发送切片并序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json error")
		return
	}

	//7.发送消息，
	//7.1 先发送消息长度
	//conn.Write（）byte 切片
	//先得到长度，然后得到切片encoding/binary---简单数字与字节序列 PutUnit32([]byte,uint32)

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write(pkgLeg) fails")
		return
	}
	//fmt.Println("send len succ:", len(data))

	//step3:发送消息本身

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fails")
		return
	}

	//3.1处理服务器返回消息
	// time.Sleep(3 * time.Second)
	// fmt.Println("sleep ")
	//3.2

	//4.1创建一个tf 实例,可以考虑全局话，
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg fails:", err)
	}
	//3.2将mes 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("loginResMes unmarshal fails:", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("succ")
		/*
			4.1显示登陆成功的菜单..
				for 可以写在里面
				需要写一个协程，保持与服务器端的通信，如果服务器有数据推送给我们，可以直接get
		*/
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println("login err:", loginResMes.Error)
	}

	return
}
