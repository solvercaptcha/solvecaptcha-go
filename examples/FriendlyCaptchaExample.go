package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	friendlyCaptcha := apisolvecaptcha.Friendly{
		SiteKey: "FCMST5VUMCBOCGQ9",
		Url:     "https://geizhals.de/455973138?fsean=5901747021356",
	}

	request := friendlyCaptcha.ToRequest()

	token, captchaId, err := client.Solve(request)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
