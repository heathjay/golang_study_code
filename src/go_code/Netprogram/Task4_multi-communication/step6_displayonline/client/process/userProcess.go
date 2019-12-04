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
	"go_code/Netprogram/Task4_multi-communication/step6_displayonline/client/utils"
	"go_code/Netprogram/Task4_multi-communication/step6_displayonline/common/message"
	"net"
	"os"
)

type UserProcess struct {
	//字段分析，
	//login需要什么东西
	//暂时不需要任何字段

	//6.2增加一个字段表示当前是哪个用户的连接

}

/*
	6.4在process增加一个Register方法完成注册任务
	1. 用户登录显示在线用户
  - 在服务器端，UserMgr.go，维护用户在线列表，需要一个数据结构，onlineUsers
              - map[int]: 目的，能够让服务器和各个在线列表进行通信，维护多客户端
              - A-B客户端，维护服务器连接，关键在连接，点对点和群发都可以实现，
              - map[int] ? 用什么更为合理?-> userprocess里面有连接, *UserProcess
              - 100->*UserProces,200->*UserProcess
              - 增删改查都要有
      步骤：
      - 在服务器端维护一个onlineUsers map[int]*UserProcess(conn)
      - 创建一个新的结构或者文件，userMgr.go，该文件的作用是对onlineUsers：map的增删改查，
	  - 什么时候增删改查，改进common/message/message.go LoginResMes里面增加一个Users []int 在在线用户id返回
	  - 当用户登录后，可以显示当前在线用户列表，只有在登陆瞬间拿到
	  -
*/
func (this *UserProcess) Register(userId int, userPw string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()

	//2.准备发送信息
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.一个registerMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPw = userPw
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	//mes.Data = registerMes	//
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
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送消息错误:", err)
	}

	mes, err = tf.ReadPkg() //mes 就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg fails:", err)
	}

	//3.2将mes 反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("loginResMes unmarshal fails:", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("register succ")
		os.Exit(0)
	} else {
		fmt.Println("login err:", registerResMes.Error)
		os.Exit(0)
	}

	return
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

	//5.1手动添加用户在redis.
	fmt.Println("send  succ:", string(data))

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
		//fmt.Println("succ")
		/*
			4.1显示登陆成功的菜单..
				for 可以写在里面
				需要写一个协程，保持与服务器端的通信，如果服务器有数据推送给我们，可以直接get
		*/
		/*
			7.1显示当前在线用户列表，遍历loginResMes.UserId

		*/
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
		}
		fmt.Println("\n\n")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println("login err:", loginResMes.Error)
	}

	return
}
