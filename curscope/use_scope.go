package curscope

import (
	"context"
	scope "go-scope"
)

// @author:      zhangzhewei
// @create:      2021-10-15 11:19
// @description:

var CtxKey = scope.CtxKey{}

// 使用之前首先要 设置SetCtx
func GetCtx() context.Context {
	return CtxKey.Get()
}

// 使用之前首先要 开启scope
func SetCtx(ctx context.Context) bool {
	return CtxKey.Set(ctx)
}
