package main

import (
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	helper "github.com/solvercaptcha/solvecaptcha-go/examples/internal"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	assetsDir := helper.GetAssetsDir(os.Args[0])
	fileName := assetsDir + "/" + "normal.jpg"

	normal := apisolvecaptcha.Normal{
		File: fileName,
	}

	req := normal.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
