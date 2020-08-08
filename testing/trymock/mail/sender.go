package mail

//go:generate mockgen -destination=./mock/mock_mail.go github.com/cao7113/hellogolang/mockup/mail Sender

/*
    go:generate 这条语句，可分为以下部分：

	声明 //go:generate （注意不要留空格）
	使用 mockgen 命令
	定义 -destination
	定义 -package
	定义 source，此处为 mail 的包路径
	定义 interfaces，此处为 Male
*/

type Sender interface {
	Send(mail string) string
}
