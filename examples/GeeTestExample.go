package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	geeTest := apisolvecaptcha.GeeTest{
		GT:        "f2ae6cadcf7886856696502e1d55e00c",
		Url:       "https://www.mysite.com/captcha/",
		ApiServer: "api-na.geetest.com",
		Challenge: "12345678abc90123d45678ef90123a456b",
	}

	req := geeTest.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
