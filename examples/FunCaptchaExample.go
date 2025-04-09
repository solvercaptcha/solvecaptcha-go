package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	funCaptcha := apisolvecaptcha.FunCaptcha{
		SiteKey: "69A21A01-CC7B-B9C6-0F9A-E7FA06677FFC",
		Url:     "https://mysite.com/page/with/funcaptcha",
	}

	req := funCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
