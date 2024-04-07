package nextcaptcha

import (
	"os"
	"testing"
)

var clientKey string

func init() {
	clientKey = os.Getenv("NEXT_CAPTCHA_CLIENT_KEY")
}

func TestReCaptchaV2Task(t *testing.T) {

	api := NewNextCaptchaAPI(clientKey, "", "", true)
	api.RecaptchaV2("https://google.com", "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-", RecaptchaV2Options{})
	api.HCaptcha("https://accounts.hcaptcha.com/demo", "a5f74b19-9e45-40e0-b45d-47ff91b7a6c2", HCaptchaOptions{})
	api.FunCaptcha("2F1CD804-FE45-F12B-9723-240962EBA6F8", FunCaptchaOptions{Data: `{"blob":"Aevp3tBIkq8erBvyuXaLZQ==.nQfRqZK25AY+SSk3q4cOPmx7p11SDLbASdijSKZ6b2QFm1grgzyW8bUJcEUIruK6NkfX5xwQdtlZOsJY0AhbaVYx0ZNK5i1QCtf7q3zG+v55pUqiFXEDkSddKqnYH4zqHtwk5YWIiFFShWkhvBHs6Y4gtmvAX92kH7oDwgWRqqZawaysrDn3BUfaHGFSWlubD/ifXpYI0SvGqdTrtzMbV1tNQ8NxM2UOUZ57O5uvCwbJtNRdjEMhL7pYrFUokKe0nMzcmPJsPKnH8IXmEbPkaVEUWDZcXiZGziupAonHz0yd"}`})
}
