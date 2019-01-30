package utils

import (
	"log"
	"testing"
)

// SubString work well with unicode
func TestSubString(t *testing.T) {
	str := "hello 世界, 你好, world!"
	if v := SubString(str, 6, 10); v != "世界, 你好, wo" {
		t.Error(v)
	}
}

func TestRandomString(t *testing.T) {
	n := 32
	format := 2
	str := RandomString(n, format)
	log.Println(str)
	if len(str) != n {
		t.Error(str)
	}
}

func BenchmarkRandomString(b *testing.B) {
	n := 32
	format := 2
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		RandomString(n, format)
	}
}
