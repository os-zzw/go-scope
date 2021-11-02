package goroutinelocal

import (
	"fmt"
	"testing"
	"time"
)

// @author:      zhangzhewei
// @create:      2021-10-15 15:18
// @description:

func TestGID(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(GetGoID())
			fmt.Println(GetGoIDFast())
		}()
	}
	time.Sleep(1000000)
}

func BenchmarkGID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(GetGoID())
	}
}

func BenchmarkGIDFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(GetGoIDFast())
	}
}

func TestGoroutineLocal(t *testing.T) {
	k := "key"
	Set(k, "main goroutine value")
	for i := 0; i < 5; i++ {
		go func() {
			gid := GetGoID()
			Set(k, fmt.Sprintf("sub:%d goroutine value", gid))
			fmt.Println(Get(k))
		}()
	}
	fmt.Println(Get(k))
	time.Sleep(1 * time.Second)
}
