package scope

// @author:      zhangzhewei
// @create:      2021-10-14 14:45
// @description:

type Int64Key struct {
	key *Key
}

func (k *Int64Key) Set(value int64) bool {
	if k.key == nil {
		k.key = &Key{}
	}
	return k.key.Set(value)
}

func (k *Int64Key) Get() int64 {
	if k.key == nil {
		k.key = &Key{}
	}
	get := k.key.Get()
	if get == nil {
		return 0
	}
	return get.(int64)
}
