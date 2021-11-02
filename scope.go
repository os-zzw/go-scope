package scope

import (
	"go-scope/goroutinelocal"
	"sync"


)

// @author:      zhangzhewei
// @create:      2021-10-13 18:01
// @description:

var scopeGoroutineLocal = &GoroutineLocal{goroutineLocal: &goroutinelocal.GoroutineLocal{Values: &sync.Map{}}}

// scope 作为key 保证全局唯一
var scopeKey = Scope{}

// RunWithNewScope 在一个新的scope内 进行操作
func RunWithNewScope(runnable Runnable) {
	beginScope()
	defer endScope()
	runnable()
}

func SupplyWithNewScope(supplier ThrowableSupplier) (interface{}, error) {
	beginScope()
	defer endScope()
	return supplier()
}

// RunWithExistScope 在一个已经存在的scope内 进行操作
func RunWithExistScope(scope Scope, action Runnable) {
	oldScope := scopeGoroutineLocal.Get()
	scopeGoroutineLocal.Set(scope)
	defer func() {
		if oldScope != nil {
			scopeGoroutineLocal.Set(oldScope)
		} else {
			scopeGoroutineLocal.Remove()
		}
	}()
	action()
}

// SuppleWithExistScope 在一个已经存在的scope内 进行操作
func SuppleWithExistScope(scope Scope, supplier ThrowableSupplier) (interface{}, error) {
	oldScope := scopeGoroutineLocal.Get()
	scopeGoroutineLocal.Set(scope)
	defer func() {
		if oldScope != nil {
			scopeGoroutineLocal.Set(oldScope)
		} else {
			scopeGoroutineLocal.Remove()
		}
	}()
	return supplier()
}

//##### 下面是具体实现 #####
// beginScope 开启 一个goroutine 私有的scope
func beginScope() *Scope {
	scope := scopeGoroutineLocal.Get()
	if scope != nil {
		panic("beginScope scope != nil")
	}
	scope = &Scope{Value: &sync.Map{}}
	scopeGoroutineLocal.Set(scope)
	return scope.(*Scope)
}

// endScope 结束 goroutine 私有的scope
func endScope() {
	scopeGoroutineLocal.Remove()
}

// getCurrentScope 获取当前 goroutine 下的私有scope
func getCurrentScope() *Scope {
	scope := scopeGoroutineLocal.Get()
	if scope == nil {
		return nil
	}
	return scope.(*Scope)
}

type Scope struct {
	Value *sync.Map
}

func (s *Scope) Set(key *Key, value interface{}) {
	if value != nil {
		s.Value.Store(key.getHolder(), value)
	} else {
		s.Value.Delete(key.getHolder())
	}
}

func (s *Scope) Get(key *Key) interface{} {
	value, ok := s.Value.Load(key.getHolder())
	if ok {
		return value
	}
	initializer := key.Initializer
	if initializer != nil {
		return initializer()
	}
	if key.DefaultValue != nil {
		return key.DefaultValue
	}
	return nil
}

type GoroutineLocal struct {
	goroutineLocal *goroutinelocal.GoroutineLocal
}

func (local GoroutineLocal) Get() interface{} {
	return local.goroutineLocal.Get(scopeKey)
}

func (local GoroutineLocal) Set(value interface{}) {
	local.goroutineLocal.Set(scopeKey, value)
}

func (local GoroutineLocal) Remove() {
	local.goroutineLocal.RemoveAll()
}
