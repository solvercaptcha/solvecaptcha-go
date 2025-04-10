![go](https://github.com/user-attachments/assets/64cd57b2-4501-4eb6-80cb-68dac7e45832)

<a href="https://github.com/solvercaptcha/solvecaptcha-python"><img src="https://github.com/user-attachments/assets/37e1d860-033b-4cf3-a158-468fc6b4debc" width="82" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-javascript"><img src="https://github.com/user-attachments/assets/4d3b4541-34b2-4ed2-a687-d694ce67e5a6" width="36" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-go"><img src="https://github.com/user-attachments/assets/5e37ab36-f32f-464b-9d33-a335e2e1b13e" width="63" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-ruby"><img src="https://github.com/user-attachments/assets/0270d56f-79b0-4c95-9b09-4de89579914b" width="75" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-cpp"><img src="https://github.com/user-attachments/assets/36de8512-acfd-44fb-bb1f-b7c793a3f926" width="45" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-php"><img src="https://github.com/user-attachments/assets/e8797843-3f61-4fa9-a155-ab0b21fb3858" width="52" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-java"><img src="https://github.com/user-attachments/assets/a3d923f6-4fec-4c07-ac50-e20da6370911" width="50" height="30"></a>
<a href="https://github.com/solvercaptcha/solvecaptcha-csharp"><img src="https://github.com/user-attachments/assets/f4d449de-780b-49ed-bb0a-b70c82ec4b32" width="38" height="30"></a>

# Go captcha solver: Bypass reCAPTCHA, Cloudflare, hCaptcha, Amazon and more

Use the [Go captcha solver](https://solvecaptcha.com/captcha-solver/go-captcha-solver-bypass) to automatically bypass any captcha — including reCAPTCHA v2, Invisible, v3, Enterprise, hCaptcha, Cloudflare Turnstile, GeeTest sliders, Amazon AWS WAF, FunCaptcha, and both image and text-based captchas.

## ✅ Supported captcha solvers

To get started quickly, check out the [captcha solver API](https://solvecaptcha.com/captcha-solver-api) documentation.

Helpful links:

- [reCAPTCHA v2 solver](https://solvecaptcha.com/captcha-solver/recaptcha-v2-solver-bypass)
- [reCAPTCHA v3 solver](https://solvecaptcha.com/captcha-solver/recaptcha-v3-solver-bypass)
- [hCaptcha solver](https://solvecaptcha.com/captcha-solver/hcaptcha-solver-bypass)
- [Text and image captcha solver](https://solvecaptcha.com/captcha-solver/image-captcha-solver-bypass)
- [Cloudflare captcha solver (Turnstile)](https://solvecaptcha.com/captcha-solver/cloudflare-captcha-solver-bypass)
- [Amazon captcha solver (AWS WAF)](https://solvecaptcha.com/captcha-solver/amazon-captcha-solver-bypass)
- [GeeTest solver](https://solvecaptcha.com/captcha-solver/slider-captcha-solver-bypass)
- [FunCaptcha (Arkose Labs) solver](https://solvecaptcha.com/captcha-solver/funcaptcha-solver-bypass)
- Other types

### 🛠️ Features 

- Fast and fully automated captcha solving
- Native support for Go 1.18+
- Uses standard net/http and JSON APIs — no external dependencies
- Easy integration with any Go backend, CLI tool, or scraper
- Supports async solving via goroutines and channels
- Pay only for successful solves
- 99.9% uptime
- 24/7 developer support

### 📦 Use cases

- Accessibility automation
- Web scraping
- Automating form submissions in Go-powered services
- Running solvers in goroutines or background jobs
- QA pipelines and headless test automation
- Security testing and bot protection research

Need help integrating with your Go project? [Open an issue](https://github.com/solvercaptcha/solvecaptcha-go/issues) or fork this repo to contribute.

- [Go captcha solver: Bypass reCAPTCHA, Cloudflare, hCaptcha, Amazon and more](#go-captcha-solver-bypass-recaptcha-cloudflare-hcaptcha-amazon-and-more)
  - [Installation](#installation)
  - [Configuration](#configuration)
    - [Client instance options](#client-instance-options)
  - [Solve captcha](#solve-captcha)
    - [Captcha options](#captcha-options)
    - [Basic example](#basic-example)
    - [Normal Captcha](#normal-captcha)
    - [Text Captcha](#text-captcha)
    - [reCAPTCHA v2](#recaptcha-v2)
    - [reCAPTCHA v3](#recaptcha-v3)
    - [reCAPTCHA Enterprise](#recaptcha-enterprise)
    - [FunCaptcha](#funcaptcha)
    - [GeeTest](#geetest)
    - [GeeTest V4](#geetest-v4)
    - [KeyCaptcha](#keycaptcha)
    - [Capy](#capy)
    - [Grid](#grid)
    - [Canvas](#canvas)
    - [ClickCaptcha](#clickcaptcha)
    - [Rotate](#rotate)
    - [Amazon WAF](#amazon-waf)
    - [Cloudflare Turnstile](#cloudflare-turnstile)
    - [Lemin Cropped Captcha](#lemin-cropped-captcha)
    - [CyberSiARA](#cybersiara)
    - [DataDome](#datadome)
    - [MTCaptcha](#mtcaptcha)
    - [Yandex](#yandex)
    - [Tencent](#tencent)
    - [atbCAPTCHA](#atbcaptcha)
    - [Cutcaptcha](#cutcaptcha)
    - [Friendly Captcha](#friendly-captcha)
    - [Audio Captcha](#audio-captcha)
  - [Other methods](#other-methods)
    - [send / getResult](#send--getresult)
    - [balance](#balance)
    - [report](#report)
  - [Proxies](#proxies)
  - [Examples](#examples)
- [Get in touch](#get-in-touch)
- [License](#license)

## Installation

To install the api client, use this:

```bash
go get -u github.com/solvercaptcha/solvecaptcha-go
```

## Configuration

Import the module like this:

```go
import (
        "github.com/solvercaptcha/solvecaptcha-go"
)
```

`Client` instance can be created like this:

```go
client := apisolvecaptcha.NewClient("YOUR_API_KEY")
```

There are few options that can be configured:

```go
client.DefaultTimeout = 120
client.RecaptchaTimeout = 600
client.PollingInterval = 100
```

### Client instance options

| Option            | Default value | Description   |
| ----------------- | ------------- | ------------- |
| default_timeout   | 120           | Timeout in seconds for all captcha types except reCAPTCHA. Defines how long the module tries to get the answer from `res.php` API endpoint    |
| recaptcha_timeout | 600           | Timeout for reCAPTCHA in seconds. Defines how long the module tries to get the answer from `res.php` API endpoint                             |
| polling_interval  | 10            | Interval in seconds between requests to `res.php` API endpoint, setting values less than 5 seconds is not recommended                         |


To get the answer manually use [GetResult method](#send--getresult)

## Solve captcha

When you submit any image-based captcha use can provide additional options to help solvecaptcha workers to solve it properly.

### Captcha options

| Option         | Default Value | Description                                                                                        |
| -------------- | ------------- | -------------------------------------------------------------------------------------------------- |
| numeric        | 0             | Defines if captcha contains numeric or other symbols [see more info in the API docs][post options] |
| min_len        | 0             | minimal answer length                                                                              |
| max_len        | 0             | maximum answer length                                                                              |
| phrase         | 0             | defines if the answer contains multiple words or not                                               |
| case_sensitive | 0             | defines if the answer is case sensitive                                                            |
| calc           | 0             | defines captcha requires calculation                                                               |
| lang           | -             | defines the captcha language, see the [list of supported languages]                                |
| hint_img       | -             | an image with hint shown to workers with the captcha                                               |
| hint_text      | -             | hint or task text shown to workers with the captcha                                                |

Below you can find basic examples for every captcha type, check out the code below.

### Basic example

Example below shows a basic solver call example with error handling.

```go
captcha := apisolvecaptcha.Normal{
   File: "/path/to/normal.jpg",
}

code, err := client.Solve(captcha.ToRequest())
if err != nil {
	if err == apisolvecaptcha.ErrTimeout {
		log.Fatal("Timeout");
	} else if err == apisolvecaptcha.ErrApi {
		log.Fatal("API error");
	} else if err == apisolvecaptcha.ErrNetwork {
		log.Fatal("Network error");
	} else {
		log.Fatal(err);
	}
}
fmt.Println("code "+code)
```

### Normal Captcha

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_normal_captcha)</sup>

To bypass a normal captcha (distorted text on image) use the following method. This method also can be used to recognize any text on the image.

```go
captcha:= apisolvecaptcha.Normal{
   File: "/path/to/normal.jpg",
   Numeric: 4,
   MinLen: 4,
   MaxLen: 20,
   Phrase: true,
   CaseSensitive: true,
   Lang: "en",
   HintImgFile: "/path/to/hint.jpg",
   HintText: "Type red symbols",
}
```

### Text Captcha

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_text_captcha)</sup>

This method can be used to bypass a captcha that requires to answer a question provided in clear text.

```go
captcha:= apisolvecaptcha.Text{
   Text: "If tomorrow is Saturday, what day is today?",
   Lang: "en",
}
```

### reCAPTCHA v2

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_recaptchav2_new)</sup>

Use this method to solve reCAPTCHA V2 and obtain a token to bypass the protection.

```go
captcha := apisolvecaptcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Invisible: true,
   Action: "verify",
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### reCAPTCHA v3

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_recaptchav3)</sup>

This method provides reCAPTCHA V3 solver and returns a token.

```go
captcha := apisolvecaptcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Version: "v3",
   Action: "verify",
   Score: 0.3,
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### reCAPTCHA Enterprise

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_recaptcha_enterprise)</sup>

reCAPTCHA Enterprise can be used as reCAPTCHA V2 and reCAPTCHA V3. Below is a usage example for both versions.

```go
// reCAPTCHA V2
captcha:=  apisolvecaptcha.ReCaptcha({
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Invisible: true,
   Action: "verify",
   Enterprise: true,
})

// reCAPTCHA V3
captcha := apisolvecaptcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Version: "v3",
   Action: "verify",
   Score: 0.3,
   Enterprise: true,
}

req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### FunCaptcha

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_funcaptcha_new)</sup>

FunCaptcha (Arkoselabs) solving method. Returns a token.

```go
captcha := apisolvecaptcha.FunCaptcha{
   SiteKey: "69A21A01-CC7B-B9C6-0F9A-E7FA06677FFC",
   Url: "https://mysite.com/page/with/funcaptcha",
   Surl: "https://client-api.arkoselabs.com",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
   Data: map[string]string{"anyKey":"anyValue"},
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### GeeTest

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_geetest)</sup>

Method to solve GeeTest puzzle captcha. Returns a set of tokens as JSON.

```go
captcha := apisolvecaptcha.GeeTest{
   GT: "f2ae6cadcf7886856696502e1d55e00c",
   ApiServer: "api-na.geetest.com",
   Challenge: "12345678abc90123d45678ef90123a456b",
   Url: "https://mysite.com/captcha.html",
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### GeeTest V4

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_geetest_v4)</sup>

Use this method to solve GeeTest v4. Returns the response in JSON.

```go
captcha:= apisolvecaptcha.GeeTestV4{
    CaptchaId: "e392e1d7fd421dc63325744d5a2b9c73",
    Url: "https://www.site.com/page/",
}
```

### KeyCaptcha

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_keycaptcha)</sup>

Token-based method to solve KeyCaptcha.

```go
captcha:= apisolvecaptcha.KeyCaptcha{
   UserId: 10,
   SessionId: "493e52c37c10c2bcdf4a00cbc9ccd1e8",
   WebServerSign: "9006dc725760858e4c0715b835472f22",
   WebServerSign2: "9006dc725760858e4c0715b835472f22",
   Url: "https://www.keycaptcha.ru/demo-magnetic/",
}
req := captchaToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### Capy

Token-based method to bypass Capy puzzle captcha.

```go
captcha:= apisolvecaptcha.Capy{
   SiteKey: "PUZZLE_Abc1dEFghIJKLM2no34P56q7rStu8v",
   Url: "https://www.mysite.com/captcha/",
}
req := captchaToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### Grid

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_grid)</sup>

Grid method is originally called Old reCAPTCHA V2 method. The method can be used to bypass any type of captcha where you can apply a grid on image and need to click specific grid boxes. Returns numbers of boxes.

```go
captcha:= apisolvecaptcha.Grid{
    File: "path/to/captcha.jpg",
    Rows: 3,
    Cols: 3,
    PreviousId: 0,
    CanSkip: false,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Select all images with an Orange",
}
```

### Canvas

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#canvas)</sup>

Canvas method can be used when you need to draw a line around an object on image. Returns a set of points' coordinates to draw a polygon.

```go
captcha:= apisolvecaptcha.Canvas{
    File: "path/to/captcha.jpg",
    PreviousId: 0,
    CanSkip: false,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Draw around apple",
}
```

### ClickCaptcha

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_clickcaptcha)</sup>

ClickCaptcha method returns coordinates of points on captcha image. Can be used if you need to click on particular points on the image.

```go
captcha:= apisolvecaptcha.Coordinates{
    File: "path/to/captcha.jpg",
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Connect the dots",
}
```

### Rotate

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_rotatecaptcha)</sup>

This method can be used to solve a captcha that asks to rotate an object. Mostly used to bypass FunCaptcha. Returns the rotation angle.

```go
captcha:= apisolvecaptcha.Rotate{
    File: "path/to/captcha.jpg",
    Angle: 15,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Put the images in the correct way",
}
```

### Amazon WAF

Use this method to solve Amazon WAF Captcha also known as AWS WAF Captcha is a part of Intelligent threat mitigation for Amazon AWS. Returns JSON with the token.

```go
captcha:= apisolvecaptcha.AmazonWAF {
   Iv: "CgAHbCe2GgAAAAAj",
   SiteKey: "0x1AAAAAAAAkg0s2VIOD34y5",
   Url: "https://non-existent-example.execute-api.us-east-1.amazonaws.com/latest",
   Context: "9BUgmlm48F92WUoqv97a49ZuEJJ50TCk9MVr3C7WMtQ0X6flVbufM4n8mjFLmbLVAPgaQ1Jydeaja94iAS49ljb",
   ChallengeScript: "https://41bcdd4fb3cb.610cd090.us-east-1.token.awswaf.com/41bcdd4fb3cb/0d21de737ccb/cd77baa6c832/challenge.js"
   CaptchaScript: "https://41bcdd4fb3cb.610cd090.us-east-1.captcha.awswaf.com/41bcdd4fb3cb/0d21de737ccb/cd77baa6c832/captcha.js"
}
```

### Cloudflare Turnstile

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#solving_cloudflare_turnstile)</sup>

Use this method to solve Cloudflare Turnstile. Returns JSON with the token.

```go
captcha:= apisolvecaptcha.CloudflareTurnstile{
   SiteKey: "0x1AAAAAAAAkg0s2VIOD34y5",
   Url: "http://mysite.com/",
}
```

### Lemin Cropped Captcha

Use this method to solve Lemin Captcha challenge. Returns JSON with answer containing the following values: answer, challenge_id.

```go
captcha:= Lemin{
   CaptchaId: "CROPPED_3dfdd5c_d1872b526b794d83ba3b365eb15a200b",
   Url:   "https://www.site.com/page/",
   DivId:     "lemin-cropped-captcha",
   ApiServer: "api.leminnow.com",
}
```

### CyberSiARA

Use this method to solve CyberSiARA and obtain a token to bypass the protection.

```go
captcha:= apisolvecaptcha.CyberSiARA{
   MasterUrlId: "12333-3123123",
   Url: "https://test.com",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}
```

### DataDome

Use this method to solve DataDome and obtain a token to bypass the protection.

> [!IMPORTANT]
> To solve the DataDome captcha, you must use a proxy. It is recommended to use [residential proxies].

```go
captcha:= apisolvecaptcha.DataDome{
   Url: "https://test.com",
   CaptchaUrl: "https://test.com/captcha/",
   Proxytype: "http",
   Proxy: "proxyuser:strongPassword@123.123.123.123:3128",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}
```

### MTCaptcha

Use this method to solve MTCaptcha and obtain a token to bypass the protection.

```go
captcha:= apisolvecaptcha.MTCaptcha{
   Url: "https://service.mtcaptcha.com/mtcv1/demo/index.html",
   SiteKey: "MTPublic-DemoKey9M",
}
```

### Yandex

Use this method to solve Yandex and obtain a token to bypass the protection.

```go
captcha:= apisolvecaptcha.Yandex{
   Url: "https://example.com",
   SiteKey: "Y5Lh0tiycconMJGsFd3EbbuNKSp1yaZESUOIHfeV",
}
```

### Tencent

Use this method to solve Tencent and obtain a token to bypass the protection.

```go
tencentCaptcha := apisolvecaptcha.Tencent{
   AppId: "2092215077",
   Url:   "http://lcec.lclog.cn/cargo/NewCargotracking?blno=BANR01XMHB0004&selectstate=BLNO",
}
```

### atbCAPTCHA

Use this method to solve atbCAPTCHA and obtain a token to bypass the protection.

```go
atbCaptcha := apisolvecaptcha.AtbCAPTCHA{
   AppId:     "af23e041b22d000a11e22a230fa8991c",
   Url:       "https://www.playzone.vip/",
   ApiServer: "https://cap.aisecurius.com",
}
```

### Cutcaptcha

Use this method to solve Cutcaptcha and obtain a token to bypass the protection.

```go
captcha:= apisolvecaptcha.Cutcaptcha{
   MiseryKey: "a1488b66da00bf332a1488993a5443c79047e752",
   DataApiKey: "SAb83IIB",
   Url: "https://example.cc/foo/bar.html",
}
```

### Friendly Captcha

Use this method to solve Friendly Captcha and obtain a token to bypass the protection.

```go
captcha:= apisolvecaptcha.Friendly{
   Url: "https://example.com",
   SiteKey: "2FZFEVS1FZCGQ9",
}
```

### Audio Captcha

Use this method to solve Audio captcha and obtain a token to bypass the protection.

```go
audio := apisolvecaptcha.Audio{
   Base64: fileBase64Str,
   Lang:   "en",
}
```

## Other methods

### Send / GetResult

These methods can be used for manual captcha submission and answer polling.

```go
id, err := client.Send(captchaToRequest())
if err != nil {
   log.Fatal(err);
}

time.Sleep(10 * time.Second)

code, err := client.GetResult(id)
if err != nil {
   log.Fatal(err);
}

if code == nil {
   log.Fatal("Not ready")
}

fmt.Println("code "+*code)
```

### balance

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#additional)</sup>

Use this method to get your account's balance

```go
balance, err := client.GetBalance()
if err != nil {
   log.Fatal(err);
}
```

### report

<sup>[API method description.](https://solvecaptcha.com/captcha-solver-api#complain)</sup>

Use this method to report good or bad captcha answer.

```go
err := client.Report(id, true) // solved correctly
err := client.Report(id, false) // solved incorrectly
```

## Error Handling

When using the SolveCaptcha API, it's important to handle errors properly to ensure smooth interaction with the service. Below is an example of error handling in the `solvecaptcha-go` library:

```go
func main() {
	result, err := client.Text("If tomorrow is Saturday, what day is today?")
	if err != nil {
		switch e := err.(type) {
		case *api.ValidationException:
			// Invalid parameters passed
			log.Fatalf("Validation error: %v", e)
		case *api.NetworkException:
			// Network error occurred
			log.Fatalf("Network error: %v", e)
		case *api.ApiException:
			// API responded with an error
			log.Fatalf("API error: %v", e)
		case *api.TimeoutException:
			// CAPTCHA is not solved within the expected time
			log.Fatalf("Timeout error: %v", e)
		default:
			// Unknown error
			log.Fatalf("Unexpected error: %v", err)
		}
	}

	fmt.Printf("CAPTCHA solved! Result: %s\n", result.Text)
}
```

## Proxies

You can pass your proxy as an additional argument for methods: recaptcha, funcaptcha, geetest, geetest v4, keycaptcha, capy puzzle, lemin, turnstile, amazon waf, CyberSiARA, DataDome, MTCaptcha and etc. The proxy will be forwarded to the API to solve the captcha.

## Examples

Examples of solving all supported captcha types are located in the [examples] directory.

## Get in touch

<a href="mailto:info@solvecaptcha.com"><img src="https://github.com/user-attachments/assets/539df209-7c85-4fa5-84b4-fc22ab93fac7" width="80" height="30"></a>
<a href="https://solvecaptcha.com/support/faq#create-ticket"><img src="https://github.com/user-attachments/assets/be044db5-2e67-46c6-8c81-04b78bd99650" width="81" height="30"></a>

## License

The code in this repository is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

<!-- Shared links -->
[SolveCaptcha]: https://solvecaptcha.com/
[post options]: https://solvecaptcha.com/captcha-solver-api#normal_post
[list of supported languages]: https://solvecaptcha.com/solvecaptcha-api#language
[examples]: ./examples/
