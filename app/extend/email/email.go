package email

// import (
// 	"log"
// 	"sync"
// 	"time"

// 	"gopkg.in/gomail.v2"
// )

// type EMAIL struct {
// 	params *EmailParam
// 	closed bool
// 	lock   sync.Mutex
// 	sender gomail.SendCloser
// }
// type EmailParam struct {
// 	// 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
// 	ServerHost string
// 	// 邮箱服务器端口，如腾讯企业邮箱为465
// 	ServerPort int
// 	// 发件人邮箱地址
// 	FromEmail string
// 	// 发件人邮箱密码
// 	FromPasswd string
// 	// 发件人别名
// 	Alias string
// }

// var g_handle *EMAIL

// func New(ep *EmailParam) *EMAIL {
// 	if g_handle == nil {
// 		g_handle = new(EMAIL)
// 		g_handle.params = ep
// 		g_handle.closed = true
// 		go g_handle.workThread()
// 	}
// 	return g_handle
// }

// func (c *EMAIL) workThread() {
// 	ticker := time.NewTicker(300 * time.Second) // 每隔5m进行一次
// 	for {
// 		<-ticker.C
// 		c.lock.Lock()
// 		if !c.closed {
// 			c.sender.Close()
// 			c.closed = true
// 		}
// 		c.lock.Unlock()
// 	}
// }

// func (c *EMAIL) getSender() gomail.SendCloser {
// 	if c.closed == true {
// 		c.closed = false
// 		c.initSender()
// 	}
// 	return c.sender
// }

// func (c *EMAIL) initSender() error {
// 	params := c.params
// 	smtp := gomail.NewDialer(params.ServerHost, params.ServerPort, params.FromEmail, params.FromPasswd)
// 	sendCloser, err := smtp.Dial()
// 	if err == nil {
// 		// sendCloser.Close()
// 		c.sender = sendCloser
// 	} else {
// 		log.Println(err)
// 	}
// 	return err
// }

// /*
// 发送邮件

// - [to] 发送给多个用户

// - [cc] 抄送给多个用户

// - [subject] 设置邮件主题

// - [body] 设置邮件正文
// */
// func (c *EMAIL) SendMail(to []string, cc []string, subject, body string) error {
// 	c.lock.Lock()
// 	defer c.lock.Unlock()
// 	mailTo := to
// 	m := gomail.NewMessage()
// 	//m.SetHeader("From", asial+"<"+c.params.FromEmail+">") // 添加别名  中文乱码
// 	m.SetAddressHeader("From", c.params.FromEmail, c.params.Alias)
// 	m.SetHeader("To", mailTo...)
// 	m.SetHeader("Cc", cc...)
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body)
// 	return c.getSender().Send(c.params.FromEmail, mailTo, m)
// }
