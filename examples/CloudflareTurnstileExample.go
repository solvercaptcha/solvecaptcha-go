package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	cloudflareTurnstile := apisolvecaptcha.CloudflareTurnstile{
		SiteKey: "0x4AAAAAAAChNiVJM_WtShFf",
		Url:     "https://ace.fusionist.io",
	}

	req := cloudflareTurnstile.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
