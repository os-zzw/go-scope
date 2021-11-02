package scope

// @author:      zhangzhewei
// @create:      2021-10-13 18:01
// @description:

// 由于Key内部包含有func 不能作为map的key, 所以添加holder作为占位符 作为key使用
type Holder struct {
	h bool
}

type Key struct {
	DefaultValue interface{}
	Initializer  func() interface{}
	holder       *Holder //不需要实现 目的是作为map的key来用
}

func (k *Key) Get() interface{} {
	currentScope := getCurrentScope()
	if currentScope != nil {
		return currentScope.Get(k)
	}
	return k.DefaultValue
}

func (k *Key) Set(value interface{}) bool {
	currentScope := getCurrentScope()
	if currentScope != nil {
		currentScope.Set(k, value)
		return true
	}
	return false
}

// 提供给scope使用
func (k *Key) getHolder() *Holder {
	holder := k.holder
	if holder == nil {
		tmpHolder := &Holder{h: true}
		k.holder = tmpHolder
	}
	return k.holder
}
