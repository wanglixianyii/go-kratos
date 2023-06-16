package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
)

var Store = base64Captcha.DefaultMemStore

type InfoCaptcha struct {
	CaptchaId string
	PicPath   string
}

// GetCaptcha 生成验证码
func GetCaptcha(ctx context.Context) (*InfoCaptcha, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return nil, err
	}

	return &InfoCaptcha{
		CaptchaId: id,
		PicPath:   b64s,
	}, nil
}

// @Pram id 验证码id
// @Pram VerifyValue 用户输入的答案
// @Result true：正确，false：失败

func VerifyCaptcha(id, VerifyValue string) bool {
	// result 为步骤1 创建的图片验证码存储对象
	return Store.Verify(id, VerifyValue, true)
}

//注意 Get(codeId, false) 中的 false 参数
//当为 true 时，根据ID获取完验证码就要删除这个验证码
//当为 false 时，根据ID获取完验证码不删除

// @Pram codeId 验证码id
// @Result 验证码答案

func GetCodeAnswer(codeId string) string {
	// result 为步骤1 创建的图片验证码存储对象
	return Store.Get(codeId, false)
}
