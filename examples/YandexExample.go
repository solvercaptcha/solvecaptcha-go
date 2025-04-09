package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	yandex := apisolvecaptcha.Yandex{
		SiteKey: "Y5Lh0tiycconMJGsFd3EbbuNKSp1yaZESUOIHfeV",
		Url:     "https://rutube.ru",
	}

	req := yandex.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
