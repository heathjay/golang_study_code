package monster

import (
	"testing"
)

//测试用例
func TestStore(t *testing.T) {
	//先创建一个monster
	monster := &Monster{
		Name:  "jack",
		Age:   29,
		Skill: "fire",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.store（）错误, 希望是%v，实际是%v", true, res)
	}
	t.Logf("good")
}
func TestReStore(t *testing.T) {
	//先创建一个monster
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.Restore（）错误, 希望是%v，实际是%v", true, res)
	}

	//进一步判断是否正确
	if monster.Name != "jack" {
		t.Fatalf("monster.Restore（）错误, 希望是%v，实际是%v", "jack", monster.Name)
	}
	t.Logf("good")
}
