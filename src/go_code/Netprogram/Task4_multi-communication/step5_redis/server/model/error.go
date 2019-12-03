package model

import (
	"errors"
)

/*
5.1 根据业务逻辑的需要，需要自定义一些错误
*/

var (
	//5.1用户没有存在
	ERROR_USER_NOTEXISTS = errors.New("该用户不存在")
	//5.1注册时用户已经存在
	ERROR_USER_EXISTS = errors.New("用户存在")
	//5.1密码不正确
	ERROR_USER_PWD = errors.New("密码不正确")
)
