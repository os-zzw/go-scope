package goroutinelocal

import (
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/petermattis/goid"
)

// @author:      zhangzhewei
// @create:      2021-10-13 18:01
// @description:

var defaultGoroutineLocal = &GoroutineLocal{Values: &sync.Map{}}

func Get(k interface{}) interface{} {
	return defaultGoroutineLocal.Get(k)
}

func Set(k interface{}, v interface{}) {
	defaultGoroutineLocal.Set(k, v)
}

func Remove(k interface{}) {
	defaultGoroutineLocal.Remove(k)
}

func RemoveAll() {
	defaultGoroutineLocal.RemoveAll()
}

// 获取gid https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-08-goroutine-id.html
func GetGoID() int64 {
	buf := make([]byte, 64)
	stackMessage := string(buf[:runtime.Stack(buf, false)])
	idField := strings.Fields(strings.TrimPrefix(stackMessage, "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic("cannot get goroutine id err : %v")
	}
	return id
}

// 目前使用的是这个 获取goid
func GetGoIDFast() int64 {
	return goid.Get()
}

type GoroutineLocal struct {
	Values *sync.Map
}

func (r *GoroutineLocal) GetAll() interface{} {
	routineNo := GetGoIDFast()
	vm, ok := r.Values.Load(routineNo)
	if !ok {
		return nil
	}
	return vm
}

func (r *GoroutineLocal) Set(k interface{}, v interface{}) {
	routineNo := GetGoIDFast()
	values := r.Values
	vm, ok := values.Load(routineNo)
	if !ok {
		vm = make(map[interface{}]interface{})
		values.Store(routineNo, vm)
	}
	vm.(map[interface{}]interface{})[k] = v
}

func (r *GoroutineLocal) Get(k interface{}) interface{} {
	routineNo := GetGoIDFast()
	values := r.Values
	vm, ok := values.Load(routineNo)
	if !ok {
		return nil
	}
	return vm.(map[interface{}]interface{})[k]
}

func (r *GoroutineLocal) Remove(k interface{}) {
	routineNo := GetGoIDFast()
	vm, ok := r.Values.Load(routineNo)
	if !ok {
		return
	}
	delete(vm.(map[interface{}]interface{}), k)
}

func (r *GoroutineLocal) RemoveAll() {
	r.Values.Delete(GetGoIDFast())
}

func (r *GoroutineLocal) SetFin(k interface{}, v interface{}, fin func()) {
	r.Set(k, v)
	defer r.RemoveAll()
	fin()
}
