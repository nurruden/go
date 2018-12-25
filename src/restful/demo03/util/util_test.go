package util

import (
	"testing"
)

func TestGenShortId(t *testing.T) {
	shorId, err := GenShortId()

	if shorId == "" || err != nil {
		t.Error("GenShortId failed!")
	}
	t.Log("GenShortId test pass")
}

//func TestGetReqID(t *testing.T) {
//	reqId := GetReqID( )
//	if reqId == ""  {
//		t.Error("Get req ID failed!")
//	}
//	t.Log("pass")
//}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
