package main

import (
	"fmt"
	"os"

	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run HCaptchaExample.go YOUR_API_KEY")
		return
	}

	client := apisolvecaptcha.NewClient(os.Args[1])

	hcaptcha := apisolvecaptcha.HCaptcha{
		SiteKey: "10000000-ffff-ffff-ffff-000000000001",
		Url:     "https://accounts.hcaptcha.com/demo",
	}

	req := hcaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
