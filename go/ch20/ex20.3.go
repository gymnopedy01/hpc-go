package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

//	func SendBook(name string, sender *fedex.FedexSender) {
//		sender.Send(name)
//	}
func SendBook(name string, sender *koreaPost.PostSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
