package scope

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

)

// @author:      zhangzhewei
// @create:      2021-10-13 20:25
// @description:

var ctx = context.Background()

func GetStr() interface{} {
	return "func get str"
}

func TestDefaultScope(t *testing.T) {
	beginScope()
	//key := &Key{DefaultValue: "adasd", Initializer: GetStr}
	fmt.Println(key.Get())
	endScope()
}

var key = &Key{}

func TestScope(t *testing.T) {
	beginScope()
	key.Set("asdasdas")
	fmt.Println(key.Get())
	endScope()
}

func TestLocal(t *testing.T) {
	local := &sync.Map{}
	key := &Key{}
	local.Store(&key, "adsad")
	load, _ := local.Load(&key)
	fmt.Println(load)
}

func TestCtxKey(t *testing.T) {
	ctxKey := &CtxKey{}
	RunWithNewScope(func() {
		ctxKey.Set(ctx)
		fmt.Println(ctxKey.Get())
	})
}

func TestStringKey(t *testing.T) {
	beginScope()
	strKey := &StringKey{}
	strKey.Set("hahdasdHHH")
	fmt.Println(strKey.Get())
	endScope()
}

func TestPointAddress(t *testing.T) {
	key := &Key{}
	fmt.Println(&key)
	key.DefaultValue = "1"
	fmt.Println(key.DefaultValue)
	Print(key)
	fmt.Println(&key)
	fmt.Println(key.DefaultValue)
}

func Print(key *Key) {
	key.DefaultValue = "2"
	fmt.Println(&key)
	fmt.Println(key.DefaultValue)
}

func TestMoreGoroutine(t *testing.T) {
	beginScope()
	strKey := &StringKey{}
	strKey.Set("goroutine main")
	for i := 0; i < 5; i++ {
		go func() {
			defer func() {
				if a := recover(); a != nil {
					fmt.Println("RECOVER", a)
				}
			}()
			RunWithNewScope(func() {
				panic("panic")
			})

		}()
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
	time.Sleep(time.Duration(3) * time.Second)
	t.Logf("main:" + strKey.Get())
	endScope()
	scopeGoroutineLocal.goroutineLocal.Values.Range(func(key, value interface{}) bool {
		t.Logf("map key is : %v, value is: %v", key, value)
		return true
	})
}

func TestSCopeUse(t *testing.T) {
	RunWithNewScope(func() {
		var key1 *StringKey
		var key2 *StringKey
		t.Log(key1 == nil)
		t.Log(key2 == nil)
		key1.Set("adasdasd key 1")
		key2.Set("adasdasd key 2")
		t.Log(key1.Get())
		t.Log(key2.Get())
	})
}
