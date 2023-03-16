package emailservice

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"onecho_sso_backend/pkg/enums"
	"onecho_sso_backend/pkg/gredis"
	"onecho_sso_backend/pkg/logging"
	"onecho_sso_backend/pkg/setting"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"go.uber.org/zap"
)

func SendValidationEmail(ctx context.Context, to string) int {
	e := email.NewEmail()
	e.From = fmt.Sprintf("onecho <%s>", setting.AppSetting.Email)
	e.Subject = "注册验证码"
	tos := make([]string, 0)
	tos = append(tos, to)
	e.To = tos

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	validateCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf(`
	<div>
		<div>
			尊敬的用户，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, t, validateCode)

	e.HTML = []byte(content)
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", setting.AppSetting.Email, setting.AppSetting.EmailAuthToken, "smtp.126.com"))
	if err != nil {
		logging.Ctx(ctx).Error("send email error error", zap.Error(err))
		return enums.SERVER_ERROR
	}

	// add to redis
	err = gredis.Set(to, validateCode, 300)
	if err != nil {
		logging.Ctx(ctx).Error("set redis  error", zap.Error(err))
		return enums.SERVER_ERROR
	}
	return enums.SUCCESS
}

func ValidateEmail(ctx context.Context, email string, code string) bool {

	storedCode, err := gredis.Get(email)
	if err != nil {
		logging.Ctx(ctx).Debug("can not find validation code for email", zap.Error(err))
		return false
	}
	return strings.Trim(string(storedCode), "\"") == code
}
