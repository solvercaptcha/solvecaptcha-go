package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	tencentCaptcha := apisolvecaptcha.Tencent{
		AppId: "2092215077",
		Url:   "http://lcec.lclog.cn/cargo/NewCargotracking?blno=BANR01XMHB0004&selectstate=BLNO",
	}

	req := tencentCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
