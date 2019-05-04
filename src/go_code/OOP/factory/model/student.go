package model

//小写不能被调用
type student struct {
	Name string
	age  int
}

func NewStudent(n string, a int) *student {
	return &student{
		Name: n,
		age:  a,
	}
}

//如果字段首字母小写，则在其他包不能直接访问，我们可以提供一个方法，

func (s *student) GetAge() int {
	return s.age
}
