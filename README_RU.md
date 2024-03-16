# NextCaptcha Golang SDK

NextCaptcha - это мощный сервис для решения капч, который поддерживает различные типы капч, включая reCAPTCHA v2,
reCAPTCHA v2 Enterprise, reCAPTCHA v3, reCAPTCHA Mobile, hCaptcha и FunCaptcha. С помощью NextCaptcha вы сможете легко
решать различные проблемы с капчей в своих скриптах и программах автоматизации.

Этот SDK предоставляет простой и удобный интерфейс Golang для взаимодействия с API NextCaptcha. Он поддерживает все
доступные типы captcha и предлагает интуитивно понятные методы решения различных типов captchas.

## Установка

Вы можете установить NextCaptcha Golang SDK с помощью команды go get:

```shell
go get github.com/nextcaptcha/nextcaptcha-go
```

## Использование
Чтобы начать использовать NextCaptcha Golang SDK, вам сначала нужно получить свой API-ключ (clientKey) на панели управления [NextCaptcha](https://dashboard.nextcaptcha.com). Затем вы можете создать экземпляр NextCaptchaAPI:
```go
import "github.com/nextcaptcha/nextcaptcha-go"

api := nextcaptcha.NewNextCaptchaAPI("YOUR_CLIENT_KEY")

```

Теперь вы можете использовать объект api для решения различных типов капч.

Для решения задач reCAPTCHA v2 используйте метод RecaptchaV2:

```go
result, err := api.RecaptchaV2("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV2Options{})

```
Решение reCAPTCHA v2 Enterprise
Чтобы решить проблемы reCAPTCHA v2 Enterprise, используйте метод RecaptchaV2Enterprise:
```go
result, err := api.RecaptchaV2Enterprise("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV2EnterpriseOptions{})
```

Решение reCAPTCHA v3
Чтобы решить проблемы reCAPTCHA v3, используйте метод RecaptchaV3:

```go
result, err := api.RecaptchaV3("https://example.com", "SITE_KEY", nextcaptcha.RecaptchaV3Options{})

```
Решение reCAPTCHA Mobile
Чтобы решить проблемы reCAPTCHA Mobile, используйте метод RecaptchaMobile:

```go
result, err := api.RecaptchaMobile("APP_KEY", nextcaptcha.RecaptchaMobileOptions{})
```

Решение проблем с hCaptcha
Чтобы решить проблемы с hCaptcha, используйте метод HCaptcha:

```go
result, err := api.HCaptcha("https://example.com", "SITE_KEY", nextcaptcha.HCaptchaOptions{})
```
Решение проблем hCaptcha Enterprise
Чтобы решить проблемы hCaptcha Enterprise, используйте метод HCaptchaEnterprise:

```go
result, err := api.HCaptchaEnterprise("https://example.com", "SITE_KEY", nextcaptcha.HCaptchaEnterpriseOptions{})
```

Решение проблем FunCaptcha
Чтобы решить проблемы FunCaptcha, используйте метод FunCaptcha:

```go
result, err := api.FunCaptcha("WEBSITE_PUBLIC_KEY", nextcaptcha.FunCaptchaOptions{})

```
Проверка баланса аккаунта
Чтобы проверить баланс вашего аккаунта NextCaptcha, используйте метод GetBalance:

```go
balance, err := api.GetBalance()
fmt.Printf("Баланс счета: %s\n", balance)
```

Вот полный пример использования NextCaptcha Golang SDK для решения задачи reCAPTCHA v2:

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
        fmt.Printf("Ошибка при решении reCAPTCHA: %v\n", err)
        return
    }
    
    if result.Status == "ready" {
        fmt.Printf("reCAPTCHA решена: %s\n", result.Solution)
    } else {
        fmt.Printf("Не удалось решить reCAPTCHA: %s\n", result.Error)
    }

}

```

## Обработка ошибок
Если при решении капчи произошла ошибка, SDK вернет структуру, содержащую информацию об ошибке. Вы можете проверить поле Status, чтобы определить, был ли запрос успешным. Если статус "готов", то капча была успешно решена и решение будет представлено в поле Solution. Если статус "не удалось", поле Error будет содержать описание ошибки.

## Вклад
Если вы обнаружили какие-либо ошибки или у вас есть предложения по улучшению, пожалуйста, не стесняйтесь оставить проблему или отправить запрос на исправление. Мы приветствуем любой вклад!

Лицензия
Этот проект лицензируется по лицензии MIT. Для получения дополнительной информации смотрите файл LICENSE.