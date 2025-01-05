package authorizer

import (
	"context"
	"testing"

	"github.com/dstgo/wilson/framework/kratosx"
)

func TestYb_GetAccessToken(t *testing.T) {
	kctx := kratosx.MustContext(context.Background())
	ins := &yb{}
	info, err := ins.GetAccessToken(kctx, GetAccessTokenRequest{
		Ak:   "e4750b34230b48e1",
		Sk:   "b0891a7f6018e5a76b085e3cb9548edd",
		Code: "5050fd57038a4078196d5a250edf802697b0deb7bf7a71f9a9a323d229332eafc7e8a67895508f7a7318414f52ef2d6e2a329f1716ae428601d89172d5043e95b87941c1b53ff0ac4eac6a0e5abd5261d96728ec37f6d045c08f455ca1c6c5922fb34a9021cca99028ad8c1bdd4ba41ebff13c709f26a2174da4a0644ce9791921aa9edcc61be5e6dfe8e7107838f844e8ce9d9540ecb220af7cab4af763ca05f2834ab6b441dedf7e1df0d3284bae9878b96206d4f990ab8921ffd33455d33163f16e489284216e5d0c7b25b7f2e270731e43d96f32f6023c2849965f37cd5120a13e7fcd198350cc61b7a58658ef37",
	})
	if err != nil {
		t.Error("获取token失败" + err.Error())
		return
	}

	reply, err := ins.GetAuthInfo(kctx, GetAuthInfoRequest{
		Token: info.Token,
	})
	if err != nil {
		t.Error("获取授权信息失败" + err.Error())
		return
	}
	t.Log(reply)
}
