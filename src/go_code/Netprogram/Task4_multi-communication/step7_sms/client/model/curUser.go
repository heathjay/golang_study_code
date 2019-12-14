package model

import (
	"go_code/Netprogram/Task4_multi-communication/step7_sms/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
