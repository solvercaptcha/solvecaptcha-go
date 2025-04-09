package main

import (
	"encoding/base64"
	"fmt"
	apisolvecaptcha "github.com/solvercaptcha/solvecaptcha-go"
	helper "github.com/solvercaptcha/solvecaptcha-go/examples/internal"
	"os"
)

func main() {
	client := apisolvecaptcha.NewClient(os.Args[1])

	assetsDir := helper.GetAssetsDir(os.Args[0])

	fileName := assetsDir + "/" + "audio-en.mp3"

	bs := helper.ReadFile2BiteSlice(fileName)

	if bs == nil {
		return
	}

	fileBase64Str := base64.StdEncoding.EncodeToString(bs)

	audio := apisolvecaptcha.Audio{
		Base64: fileBase64Str,
		Lang:   "en",
	}

	req := audio.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
