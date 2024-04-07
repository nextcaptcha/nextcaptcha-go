# NextCaptcha Golang SDK

NextCaptcha is a powerful captcha solving service that supports various types of captchas including reCAPTCHA v2, reCAPTCHA v2 Enterprise, reCAPTCHA v3, reCAPTCHA Mobile, hCaptcha, and FunCaptcha. With NextCaptcha, you can easily solve a variety of captcha challenges in your automation scripts and programs.

This SDK provides a simple and easy-to-use Golang interface for interacting with the NextCaptcha API. It supports all available captcha types and offers intuitive methods for solving different types of captchas.

## Installation

You can install the NextCaptcha Golang SDK using the go get command:

```shell
go get github.com/nextcaptcha/nextcaptcha-go
```
## Usage
To start using the NextCaptcha Golang SDK, you first need to obtain your API key (clientKey) from the [NextCaptcha](https://dashboard.nextcaptcha.com) Dashboard. Then, you can create a NextCaptchaAPI instance:
```go
import "github.com/nextcaptcha/nextcaptcha-go"

api := nextcaptcha.NewNextCaptchaAPI(clientKey, "", "", true)

```
Now, you can use the api object to solve various types of captchas.

To solve reCAPTCHA v2 challenges, use the RecaptchaV2 method:

```go
result, err := api.RecaptchaV2("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV2Options{})
```
Solving reCAPTCHA v2 Enterprise
To solve reCAPTCHA v2 Enterprise challenges, use the RecaptchaV2Enterprise method:
```go
result, err := api.RecaptchaV2Enterprise("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV2EnterpriseOptions{})
```
Solving reCAPTCHA v3
To solve reCAPTCHA v3 challenges, use the RecaptchaV3 method:

```go
result, err := api.RecaptchaV3("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV3Options{})
```
Solving reCAPTCHA Mobile
To solve reCAPTCHA Mobile challenges, use the RecaptchaMobile method:

```go
result, err := api.RecaptchaMobile("APP_KEY", nextcaptcha.RecaptchaMobileOptions{})
```
Solving hCaptcha
To solve hCaptcha challenges, use the HCaptcha method:

```go
result, err := api.HCaptcha("https://example.com", "SITE_KEY", nextcaptcha.HCaptchaOptions{})
```
Solving hCaptcha Enterprise
To solve hCaptcha Enterprise challenges, use the HCaptchaEnterprise method:

```go
result, err := api.HCaptchaEnterprise("https://example.com", "SITE_KEY", nextcaptcha.HCaptchaEnterpriseOptions{})
```
Solving FunCaptcha
To solve FunCaptcha challenges, use the FunCaptcha method:

```go
result, err := api.FunCaptcha("WEBSITE_PUBLIC_KEY", nextcaptcha.FunCaptchaOptions{})
```
Checking Account Balance
To check your NextCaptcha account balance, use the GetBalance method:

```go
balance, err := api.GetBalance()
fmt.Printf("Account balance: %s\n", balance)
```
Here is a complete example of using the NextCaptcha Golang SDK to solve a reCAPTCHA v2 challenge:

```go
package main

import (
    "fmt"
    "github.com/nextcaptcha/nextcaptcha-go"
)

func main() {
    api := nextcaptcha.NewNextCaptchaAPI("YOUR_CLIENT_KEY")

    result, err := api.RecaptchaV2("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV2Options{})
    if err != nil {
        fmt.Printf("Error solving reCAPTCHA: %v\n", err)
        return
    }
    
    if result.Status == "ready" {
        fmt.Printf("reCAPTCHA solved: %s\n", result.Solution)
    } else {
        fmt.Printf("Failed to solve reCAPTCHA: %s\n", result.Error)
    }
}
```

## Error Handling
If an error occurs while solving a captcha, the SDK will return a struct containing the error information. You can check the Status field to determine if the request was successful. If the status is "ready", the captcha has been successfully solved and the solution will be provided in the Solution field. If the status is "failed", the Error field will contain a description of the error.

## Contributing
If you find any bugs or have suggestions for improvement, please feel free to submit an issue or send a pull request. We welcome all contributions!

License
This project is licensed under the MIT License. For more information, please see the LICENSE file.