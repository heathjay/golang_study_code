package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//编写一个monster结构体，字段name , string, age
//绑定方法store,可以将monster序列化并保存在文件中
type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println(err)
		return false
	}

	//保存到文件内
	filePath := "./monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//restore

func (this *Monster) ReStore() bool {
	//1.读取序列文件的内容
	filePath := "./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	//2.读取data 【】byte反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
