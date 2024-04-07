package nextcaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	HOST                        = "https://api.nextcaptcha.com"
	TIMEOUT                     = 45 * time.Second
	PENDING_STATUS              = "pending"
	PROCESSING_STATUS           = "processing"
	READY_STATUS                = "ready"
	FAILED_STATUS               = "failed"
	RECAPTCHAV2_TYPE            = "RecaptchaV2TaskProxyless"
	RECAPTCHAV2_ENTERPRISE_TYPE = "RecaptchaV2EnterpriseTaskProxyless"
	RECAPTCHAV3_PROXYLESS_TYPE  = "RecaptchaV3TaskProxyless"
	RECAPTCHAV3_TYPE            = "RecaptchaV3Task"
	RECAPTCHA_MOBILE_TYPE       = "RecaptchaMobileProxyless"
	HCAPTCHA_TYPE               = "HCaptchaTask"
	HCAPTCHA_PROXYLESS_TYPE     = "HCaptchaTaskProxyless"
	HCAPTCHA_ENTERPRISE_TYPE    = "HCaptchaEnterpriseTask"
	FUNCAPTCHA_TYPE             = "FunCaptchaTask"
	FUNCAPTCHA_PROXYLESS_TYPE   = "FunCaptchaTaskProxyless"
)

type TaskBadParametersError struct {
	s string
}

func (e *TaskBadParametersError) Error() string {
	return e.s
}

type ApiClient struct {
	clientKey   string
	solftId     string
	callbackUrl string
	openLog     bool
	httpClient  *http.Client
}

func NewApiClient(clientKey, solftId, callbackUrl string, openLog bool) *ApiClient {
	return &ApiClient{
		clientKey:   clientKey,
		solftId:     solftId,
		callbackUrl: callbackUrl,
		openLog:     openLog,
		httpClient:  &http.Client{Timeout: TIMEOUT},
	}
}

func (c *ApiClient) getBalance() (string, error) {
	data := map[string]string{"clientKey": c.clientKey}
	resp, err := c.postJSON("/getBalance", data)
	if err != nil {
		if c.openLog {
			log.Printf("Error: %v", err)
		}
		return "", err
	}

	balance := resp["balance"].(string)
	if c.openLog {
		log.Printf("Balance: %s", balance)
	}

	return balance, nil
}

func (c *ApiClient) send(task map[string]interface{}) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"clientKey":   c.clientKey,
		"solftId":     c.solftId,
		"callbackUrl": c.callbackUrl,
		"task":        task,
	}
	resp, err := c.postJSON("/createTask", data)
	if err != nil {
		if c.openLog {
			log.Printf("Error: %v", err)
			log.Printf("Data: %v", data)
		}
		return nil, err
	}

	taskId := resp["taskId"].(float64)
	if c.openLog {
		log.Printf("Task %f created %v", taskId, resp)
	}

	startTime := time.Now()
	for {
		if time.Since(startTime) > TIMEOUT {
			return map[string]interface{}{
				"errorId":          12,
				"errorDescription": "Timeout",
				"status":           "failed",
			}, nil
		}

		data := map[string]any{
			"clientKey": c.clientKey,
			"taskId":    taskId,
		}
		resp, err := c.postJSON("/getTaskResult", data)
		if err != nil {
			if c.openLog {
				log.Printf("Error: %v", err)
			}
			return nil, err
		}

		status := resp["status"].(string)
		if c.openLog {
			log.Printf("Task status: %s", status)
		}

		if status == READY_STATUS {
			log.Printf("Task %f ready %v", taskId, resp)
			return resp, nil
		}
		if status == FAILED_STATUS {
			log.Printf("Task %f failed %v", taskId, resp)
			return resp, nil
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *ApiClient) postJSON(path string, data interface{}) (map[string]interface{}, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, HOST+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		if c.openLog {
			log.Printf("Error: %d %s", resp.StatusCode, string(body))
		}
		return nil, errors.New(string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

type NextCaptchaAPI struct {
	api *ApiClient
}

func NewNextCaptchaAPI(clientKey, solftId, callbackUrl string, openLog bool) *NextCaptchaAPI {
	log.Printf("NextCaptchaAPI created with clientKey=%s solftId=%s callbackUrl=%s", clientKey, solftId, callbackUrl)
	api := NewApiClient(clientKey, solftId, callbackUrl, openLog)
	return &NextCaptchaAPI{api}
}

type RecaptchaV2Options struct {
	RecaptchaDataSValue string
	IsInvisible         bool
	ApiDomain           string
}

func (api *NextCaptchaAPI) RecaptchaV2(websiteURL, websiteKey string, options RecaptchaV2Options) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":       RECAPTCHAV2_TYPE,
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
	if options.RecaptchaDataSValue != "" {
		task["recaptchaDataSValue"] = options.RecaptchaDataSValue
	}
	if options.IsInvisible {
		task["isInvisible"] = options.IsInvisible
	}
	if options.ApiDomain != "" {
		task["apiDomain"] = options.ApiDomain
	}
	return api.api.send(task)
}

type RecaptchaV2EnterpriseOptions struct {
	EnterprisePayload map[string]interface{}
	IsInvisible       bool
	ApiDomain         string
}

func (api *NextCaptchaAPI) RecaptchaV2Enterprise(websiteURL, websiteKey string, options RecaptchaV2EnterpriseOptions) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":       RECAPTCHAV2_ENTERPRISE_TYPE,
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
	if options.EnterprisePayload != nil {
		task["enterprisePayload"] = options.EnterprisePayload
	}
	if options.IsInvisible {
		task["isInvisible"] = options.IsInvisible
	}
	if options.ApiDomain != "" {
		task["apiDomain"] = options.ApiDomain
	}
	return api.api.send(task)
}

type RecaptchaV3Options struct {
	PageAction    string
	ApiDomain     string
	ProxyType     string
	ProxyAddress  string
	ProxyPort     int
	ProxyLogin    string
	ProxyPassword string
}

func (api *NextCaptchaAPI) RecaptchaV3(websiteURL, websiteKey string, options RecaptchaV3Options) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":       RECAPTCHAV3_PROXYLESS_TYPE,
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
	if options.PageAction != "" {
		task["pageAction"] = options.PageAction
	}
	if options.ApiDomain != "" {
		task["apiDomain"] = options.ApiDomain
	}
	if options.ProxyAddress != "" {
		task["type"] = RECAPTCHAV3_TYPE
		task["proxyType"] = options.ProxyType
		task["proxyAddress"] = options.ProxyAddress
		task["proxyPort"] = options.ProxyPort
		task["proxyLogin"] = options.ProxyLogin
		task["proxyPassword"] = options.ProxyPassword
	}
	return api.api.send(task)
}

type RecaptchaMobileOptions struct {
	AppPackageName string
	AppAction      string
}

func (api *NextCaptchaAPI) RecaptchaMobile(appKey string, options RecaptchaMobileOptions) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":   RECAPTCHA_MOBILE_TYPE,
		"appKey": appKey,
	}
	if options.AppPackageName != "" {
		task["appPackageName"] = options.AppPackageName
	}
	if options.AppAction != "" {
		task["appAction"] = options.AppAction
	}
	return api.api.send(task)
}

type HCaptchaOptions struct {
	IsInvisible       bool
	EnterprisePayload map[string]interface{}
	ProxyType         string
	ProxyAddress      string
	ProxyPort         int
	ProxyLogin        string
	ProxyPassword     string
}

func (api *NextCaptchaAPI) HCaptcha(websiteURL, websiteKey string, options HCaptchaOptions) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":       HCAPTCHA_PROXYLESS_TYPE,
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
	if options.IsInvisible {
		task["isInvisible"] = options.IsInvisible
	}
	if options.EnterprisePayload != nil {
		task["enterprisePayload"] = options.EnterprisePayload
	}
	if options.ProxyAddress != "" {
		task["type"] = HCAPTCHA_TYPE
		task["proxyType"] = options.ProxyType
		task["proxyAddress"] = options.ProxyAddress
		task["proxyPort"] = options.ProxyPort
		task["proxyLogin"] = options.ProxyLogin
		task["proxyPassword"] = options.ProxyPassword
	}
	return api.api.send(task)
}

type HCaptchaEnterpriseOptions struct {
	EnterprisePayload map[string]interface{}
	IsInvisible       bool
	ProxyType         string
	ProxyAddress      string
	ProxyPort         int
	ProxyLogin        string
	ProxyPassword     string
}

func (api *NextCaptchaAPI) HCaptchaEnterprise(websiteURL, websiteKey string, options HCaptchaEnterpriseOptions) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":       HCAPTCHA_ENTERPRISE_TYPE,
		"websiteURL": websiteURL,
		"websiteKey": websiteKey,
	}
	if options.EnterprisePayload != nil {
		task["enterprisePayload"] = options.EnterprisePayload
	}
	if options.IsInvisible {
		task["isInvisible"] = options.IsInvisible
	}
	task["proxyType"] = options.ProxyType
	task["proxyAddress"] = options.ProxyAddress
	task["proxyPort"] = options.ProxyPort
	task["proxyLogin"] = options.ProxyLogin
	task["proxyPassword"] = options.ProxyPassword
	return api.api.send(task)
}

type FunCaptchaOptions struct {
	WebsiteURL    string
	Data          string
	ProxyType     string
	ProxyAddress  string
	ProxyPort     int
	ProxyLogin    string
	ProxyPassword string
}

func (api *NextCaptchaAPI) FunCaptcha(websitePublicKey string, options FunCaptchaOptions) (map[string]interface{}, error) {
	task := map[string]interface{}{
		"type":             FUNCAPTCHA_PROXYLESS_TYPE,
		"websitePublicKey": websitePublicKey,
	}
	if options.WebsiteURL != "" {
		task["websiteURL"] = options.WebsiteURL
	}
	if options.Data != "" {
		task["data"] = options.Data
	}
	if options.ProxyAddress != "" {
		task["type"] = FUNCAPTCHA_TYPE
		task["proxyType"] = options.ProxyType
		task["proxyAddress"] = options.ProxyAddress
		task["proxyPort"] = options.ProxyPort
		task["proxyLogin"] = options.ProxyLogin
		task["proxyPassword"] = options.ProxyPassword
	}
	return api.api.send(task)
}
