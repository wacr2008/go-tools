package email

import (
	"github.com/go-gomail/gomail"
	"fmt"
)

const (
	smtpHost     = "smtphm.qiye.163.com"
	smtpPort     = 25
	smtpUsername = "gitlab@ctkjgroup.com"
	smtpPassword = "qiye@163"
)

func SendEmailCode(to, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", to)
	// 设置抄送
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "【FengWo】安全验证")
	tempHtml1 := `
<table cellpadding="0" cellspacing="0" style="border: 1px solid #cdcdcd; width: 640px; margin:auto;font-size: 12px; color: #1E2731; line-height: 20px;">
    <tr>
        <td colspan="3" align="center" style="background-color:#454c6d; height: 35px; padding: 30px 0"><a
                href="https://www.fengwo.com" target="_blank"><img height="60" src="http://upload.obsgs.com/%E8%A7%86%E9%A2%91.png"/></a></td>
    </tr>
    <tr style="height: 30px;">&nbsp;</tr>
    <tr>
        <td width="20"></td>
        <td style="line-height: 40px">
            您好：<br/>`
	tempHtml2 := "【FengWo】安全项设置安全验证: " + code + " <br/>"
	tempHtml3 := `
            出于安全原因，该验证码将于15分钟后失效。请勿将验证码透露给他人。<br/>
        </td>
        <td width="20"></td>
    </tr>
    <tr style="height: 20px;">&nbsp;</tr>
    <tr>
        <td width="20"></td>
        <td>
            <br/>
            FengWo团队<br/>
            <a href="https://www.huobi.co">https://www.fengwo.com</a><br/>
        </td>
        <td width="20"></td>
    </tr>
    <tr style="height: 50px;">&nbsp;</tr>
</table>`
	m.SetBody("text/html", tempHtml1+tempHtml2+tempHtml3)
	// 附件
	// m.Attach("./视频.png")

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		//return err
		// TODO 这里应该记录并重发
		fmt.Println(err)
	}
	return nil
}
