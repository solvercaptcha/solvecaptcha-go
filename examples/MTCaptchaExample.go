package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	mtCaptcha := apisolvecaptcha.MTCaptcha{
		SiteKey: "MTPublic-KzqLY1cKH",
		Url:     "https://solvecaptcha.com/demo/mtcaptcha",
	}

	req := mtCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
