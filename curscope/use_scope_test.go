package curscope

import (
	"context"
	"fmt"
	scope "go-scope"
	"testing"

)

// @author:      zhangzhewei
// @create:      2021-10-15 14:09
// @description:

var ctx = context.Background()

var int64Key = &scope.Int64Key{}

func TestCtxScope(t *testing.T) {
	scope.RunWithNewScope(func() {
		SetCtx(ctx)
		int64Key.Set(1312313)
		fmt.Println(GetCtx())
		fmt.Println(int64Key.Get())
	})
}
