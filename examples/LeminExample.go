package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	lemin := apisolvecaptcha.Lemin{
		CaptchaId: "CROPPED_d3d4d56_73ca4008925b4f83a8bed59c2dd0df6d",
		Url:       "https://www.site.com/page/",
		ApiServer: "api.leminnow.com",
	}

	req := lemin.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
