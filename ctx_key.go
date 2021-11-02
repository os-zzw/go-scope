package scope

import "context"

// @author:      zhangzhewei
// @create:      2021-10-14 14:46
// @description:

type CtxKey struct {
	key *Key
}

func (k *CtxKey) Set(value context.Context) bool {
	if k.key == nil {
		k.key = &Key{}
	}
	return k.key.Set(value)
}

func (k *CtxKey) Get() context.Context {
	if k.key == nil {
		k.key = &Key{}
	}
	get := k.key.Get()
	if get == nil {
		return nil
	}
	return get.(context.Context)
}
