package scope

// @author:      zhangzhewei
// @create:      2021-10-26 16:28
// @description:

type BoolKey struct {
	key *Key
}

func (k *BoolKey) Set(value bool) bool {
	if k.key == nil {
		k.key = &Key{}
	}
	return k.key.Set(value)
}

// 默认是false
func (k *BoolKey) Get() bool {
	if k.key == nil {
		k.key = &Key{}
	}
	get := k.key.Get()
	if get == nil {
		return false
	}
	return get.(bool)
}
