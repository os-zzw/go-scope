package scope

// @author:      zhangzhewei
// @create:      2021-10-14 14:45
// @description:

type StringKey struct {
	key *Key
}

func (k *StringKey) Set(value string) bool {
	if k.key == nil {
		k.key = &Key{}
	}
	return k.key.Set(value)
}

func (k *StringKey) Get() string {
	if k.key == nil {
		k.key = &Key{}
	}
	get := k.key.Get()
	if get == nil {
		return ""
	}
	return get.(string)
}
