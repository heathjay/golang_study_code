package model

/*
	5.1定一个UserDao 这个结构体，完成对User结构体的操作
		我们在服务器启动后就初始化一个userDao 的一个实例，在需要用的时候直接用，做成一个全局的变量，在需要和redis操作时就直接使用就可以了

*/
import (
	"encoding/json"
	"fmt"
	"go_code/Netprogram/Task4_multi-communication/step5_redis/common/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	//5.1要操作Redis,需要拿到连接池
	pool *redis.Pool
}

//5.1使用工厂模式创建一个userDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考一下再UserDao应该提供什么方法给我们
//5.1 根据用户id返回一个User实例 *err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//5.1 通过给定的id去redis查询这个用户
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//5.1需要把res反序列化成一个user对象
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.unmarshal err =", err)
		return
	}
	return
}

//5.1 完成登陆的校验,如果正确返回一个user实例，如果不正确则返回一个错误信息
func (this *UserDao) Login(userId int, userPw string) (user *User, err error) {
	//先从UserDao的连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//至少有这个用户存在，获取到用户实例，但是密码不一定正确，
	if user.UserPw != userPw {
		err = ERROR_USER_PWD
		return
	}
	return
}

/*6.1 增加register 方法*/
func (this *UserDao) Register(user *message.User) (err error) {
	//先从UserDao的连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//说明该用户还没有注册过
	//则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户出错:", err)
		return
	}
	return
}
