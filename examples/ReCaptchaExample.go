package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	reCaptcha := apisolvecaptcha.ReCaptcha{
		SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
		Url:     "ttps://mysite.com/page/with/recaptcha",
	}

	req := reCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
