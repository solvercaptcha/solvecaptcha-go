package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	cutCaptcha := apisolvecaptcha.CutCaptcha{
		MiseryKey:  "a1488b66da00bf332a1488993a5443c79047e752",
		Url:        "https://filecrypt.co/Container/237D4D0995.html",
		DataApiKey: "SAb83IIB",
	}

	req := cutCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
