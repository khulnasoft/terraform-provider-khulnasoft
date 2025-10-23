package client

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	neturl "net/url"

	"github.com/khulnasoft/terraform-provider-khulnasoft/consts"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/http/httpproxy"
	"golang.org/x/time/rate"
)

// Client - API client
type Client struct {
	url        string
	saasUrl    string
	tokenUrl   string
	user       string
	password   string
	apiKey     string
	apiSecret  string
	token      string
	name       string
	gorequest  *gorequest.SuperAgent
	clientType string
	limiter    *rate.Limiter
}

const Csp string = "csp"
const Saas = "saas"
const SaasDev = "saasDev"

const UserAgentBase = "terraform-provider-khulnasoft"

var version string

// NewClient - initialize and return the Client
func NewClient(url, user, password string, verifyTLS bool, caCertByte []byte) *Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: !verifyTLS,
	}

	roots := x509.NewCertPool()
	if len(caCertByte) > 0 {
		roots.AppendCertsFromPEM(caCertByte)

		if verifyTLS {
			tlsConfig = &tls.Config{
				RootCAs: roots,
			}
		}
	}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	c := &Client{
		url:       url,
		user:      user,
		password:  password,
		gorequest: request,
		// we are setting rate limit for 10 connection per second
		limiter: rate.NewLimiter(10, 3),
	}

	// Determine if we need to use a proxy
	uURL, _ := neturl.Parse(c.url)
	proxy, _ := httpproxy.FromEnvironment().ProxyFunc()(uURL)
	if proxy != nil {
		c.gorequest.Proxy(proxy.String())
	}

	switch url {
	case consts.SaasUrl:
		c.clientType = Saas
		c.tokenUrl = consts.SaasTokenUrl
		c.saasUrl = consts.SaasUrl
		break
	case consts.SaasEu1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasEu1TokenUrl
		c.saasUrl = consts.SaasEu1Url
		break
	case consts.SaasAsia1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia1TokenUrl
		c.saasUrl = consts.SaasAsia1Url
		break
	case consts.SaasAsia2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia2TokenUrl
		c.saasUrl = consts.SaasAsia2Url
		break
	case consts.SaaSAu2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAu2TokenUrl
		c.saasUrl = consts.SaaSAu2Url
		break
	case consts.SaasDevUrl:
		c.clientType = SaasDev
		c.tokenUrl = consts.SaasDevTokenUrl
		c.saasUrl = consts.SaasDevUrl
		break
	default:
		c.clientType = Csp
		break
	}

	return c
}

// NewClientWithAPIKey - initialize and return the Client with API key authentication
func NewClientWithAPIKey(url, apiKey, apiSecret string, verifyTLS bool, caCertByte []byte) *Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: !verifyTLS,
	}

	roots := x509.NewCertPool()
	if len(caCertByte) > 0 {
		roots.AppendCertsFromPEM(caCertByte)

		if verifyTLS {
			tlsConfig = &tls.Config{
				RootCAs: roots,
			}
		}
	}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	c := &Client{
		url:       url,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		gorequest: request,
		// we are setting rate limit for 10 connection per second
		limiter: rate.NewLimiter(10, 3),
	}

	// Determine if we need to use a proxy
	uURL, _ := neturl.Parse(c.url)
	proxy, _ := httpproxy.FromEnvironment().ProxyFunc()(uURL)
	if proxy != nil {
		c.gorequest.Proxy(proxy.String())
	}

	switch url {
	case consts.SaasUrl:
		c.clientType = Saas
		c.tokenUrl = consts.SaasTokenUrl
		c.saasUrl = consts.SaasUrl
		break
	case consts.SaasEu1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasEu1TokenUrl
		c.saasUrl = consts.SaasEu1Url
		break
	case consts.SaasAsia1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia1TokenUrl
		c.saasUrl = consts.SaasAsia1Url
		break
	case consts.SaasAsia2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia2TokenUrl
		c.saasUrl = consts.SaasAsia2Url
		break
	case consts.SaaSAu2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAu2TokenUrl
		c.saasUrl = consts.SaaSAu2Url
		break
	case consts.SaasDevUrl:
		c.clientType = SaasDev
		c.tokenUrl = consts.SaasDevTokenUrl
		c.saasUrl = consts.SaasDevUrl
		break
	default:
		c.clientType = Csp
		break
	}

	return c
}

// NewClientWithTokenAuth - initialize and return the Client with username/password authentication
func NewClientWithTokenAuth(url, user, password string, verifyTLS bool, caCertByte []byte) *Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: !verifyTLS,
	}

	roots := x509.NewCertPool()
	if len(caCertByte) > 0 {
		roots.AppendCertsFromPEM(caCertByte)

		if verifyTLS {
			tlsConfig = &tls.Config{
				RootCAs: roots,
			}
		}
	}

	request := gorequest.New().TLSClientConfig(tlsConfig)

	c := &Client{
		url:       url,
		user:      user,
		password:  password,
		gorequest: request,
		// we are setting rate limit for 10 connection per second
		limiter: rate.NewLimiter(10, 3),
	}

	// Determine if we need to use a proxy
	uURL, _ := neturl.Parse(c.url)
	proxy, _ := httpproxy.FromEnvironment().ProxyFunc()(uURL)
	if proxy != nil {
		c.gorequest.Proxy(proxy.String())
	}

	switch url {
	case consts.SaasUrl:
		c.clientType = Saas
		c.tokenUrl = consts.SaasTokenUrl
		c.saasUrl = consts.SaasUrl
		break
	case consts.SaasEu1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasEu1TokenUrl
		c.saasUrl = consts.SaasEu1Url
		break
	case consts.SaasAsia1Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia1TokenUrl
		c.saasUrl = consts.SaasAsia1Url
		break
	case consts.SaasAsia2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAsia2TokenUrl
		c.saasUrl = consts.SaasAsia2Url
		break
	case consts.SaaSAu2Url:
		c.clientType = Saas
		c.tokenUrl = consts.SaasAu2TokenUrl
		c.saasUrl = consts.SaaSAu2Url
		break
	case consts.SaasDevUrl:
		c.clientType = SaasDev
		c.tokenUrl = consts.SaasDevTokenUrl
		c.saasUrl = consts.SaasDevUrl
		break
	default:
		c.clientType = Csp
		break
	}

	return c
}

func (cli *Client) SetAuthToken(token string) {
	cli.token = token
}

func (cli *Client) SetUrl(url string) {
	cli.url = url
}

func (cli *Client) GetAuthToken() (string, string, error) {
	var err error

	// Use API key authentication if API key is provided
	if cli.apiKey != "" && cli.apiSecret != "" {
		if cli.clientType == "csp" {
			_, err = cli.GetCspAuthTokenWithAPIKey()
		} else {
			_, _, err = cli.GetUSEAuthTokenWithAPIKey()
		}
	} else {
		// Use username/password authentication
		if cli.clientType == "csp" {
			_, err = cli.GetCspAuthToken()
		} else {
			_, _, err = cli.GetUSEAuthToken()
		}
	}

	if err != nil {
		return "", "", err
	}
	return cli.token, cli.url, nil
}

// GetAuthToken - Connect to Khulnasoft and return a JWT bearerToken (string)
func (cli *Client) GetCspAuthToken() (string, error) {
	resp, body, errs := cli.makeRequest().Post(cli.url + "/api/v1/login").
		Send(`{"id":"` + cli.user + `", "password":"` + cli.password + `"}`).End()
	if errs != nil {
		return "", getMergedError(errs)
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		cli.token = raw["token"].(string)
		return cli.token, nil
	}

	return "", fmt.Errorf("request failed. status: %s, response: %s", resp.Status, body)
}

// GetUSEAuthToken - Connect to Khulnasoft SaaS solution and return a JWT bearerToken (string)
func (cli *Client) GetUSEAuthToken() (string, string, error) {
	var provUrl string

	switch cli.url {
	case consts.SaasUrl:
		provUrl = consts.SaasProvUrl
		break
	case consts.SaasEu1Url:
		provUrl = consts.SaasEu1ProvUrl
		break
	case consts.SaasAsia1Url:
		provUrl = consts.SaasAsia1ProvUrl
		break
	case consts.SaasAsia2Url:
		provUrl = consts.SaasAsia2ProvUrl
		break
	case consts.SaaSAu2Url:
		provUrl = consts.SaasAu2ProvUrl
		break
	case consts.SaasDevUrl:
		provUrl = consts.SaasDevProvUrl
		break
	default:
		return "", "", fmt.Errorf(fmt.Sprintf("%v URL is not allowed USE url", cli.url))
	}

	resp, body, errs := cli.makeRequest().Post(cli.tokenUrl + "/v2/signin").
		Send(`{"email":"` + cli.user + `", "password":"` + cli.password + `"}`).End()
	if errs != nil {
		return "", "", getMergedError(errs)
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		data := raw["data"].(map[string]interface{})
		cli.token = data["token"].(string)
		//get the ese_url to make the API requests.
		request := cli.gorequest
		if request == nil {
			return "", "", fmt.Errorf("request is uninitialized")
		}

		request.Set("Authorization", "Bearer "+cli.token)
		events, body, errs := request.Clone().Get(provUrl + "/v1/envs").End()
		if errs != nil || events == nil {
			if events != nil {
				log.Println(events.StatusCode)
			}
			err := fmt.Errorf("error calling %s", provUrl)
			return "", "", err
		}

		if events.StatusCode == 200 {
			var raw map[string]interface{}
			_ = json.Unmarshal([]byte(body), &raw)
			data := raw["data"].(map[string]interface{})
			cli.url = "https://" + data["ese_url"].(string)
		}

		return cli.token, cli.url, nil
	}

	return "", "", fmt.Errorf("request failed. status: %s, response: %s", resp.Status, body)
}

// GetCspAuthTokenWithAPIKey - Connect to Khulnasoft using API key and return a JWT bearerToken (string)
func (cli *Client) GetCspAuthTokenWithAPIKey() (string, error) {
	resp, body, errs := cli.makeRequest().Post(cli.url + "/api/v1/login").
		Send(`{"id":"` + cli.apiKey + `", "password":"` + cli.apiSecret + `"}`).End()
	if errs != nil {
		return "", getMergedError(errs)
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		cli.token = raw["token"].(string)
		return cli.token, nil
	}

	return "", fmt.Errorf("request failed. status: %s, response: %s", resp.Status, body)
}

// GetUSEAuthTokenWithAPIKey - Connect to Khulnasoft SaaS solution using API key and return a JWT bearerToken (string)
func (cli *Client) GetUSEAuthTokenWithAPIKey() (string, string, error) {
	var provUrl string

	switch cli.url {
	case consts.SaasUrl:
		provUrl = consts.SaasProvUrl
		break
	case consts.SaasEu1Url:
		provUrl = consts.SaasEu1ProvUrl
		break
	case consts.SaasAsia1Url:
		provUrl = consts.SaasAsia1ProvUrl
		break
	case consts.SaasAsia2Url:
		provUrl = consts.SaasAsia2ProvUrl
		break
	case consts.SaaSAu2Url:
		provUrl = consts.SaasAu2ProvUrl
		break
	case consts.SaasDevUrl:
		provUrl = consts.SaasDevProvUrl
		break
	default:
		return "", "", fmt.Errorf(fmt.Sprintf("%v URL is not allowed USE url", cli.url))
	}

	resp, body, errs := cli.makeRequest().Post(cli.tokenUrl + "/v2/signin").
		Send(`{"email":"` + cli.apiKey + `", "password":"` + cli.apiSecret + `"}`).End()
	if errs != nil {
		return "", "", getMergedError(errs)
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		data := raw["data"].(map[string]interface{})
		cli.token = data["token"].(string)
		//get the ese_url to make the API requests.
		request := cli.gorequest
		if request == nil {
			return "", "", fmt.Errorf("request is uninitialized")
		}

		request.Set("Authorization", "Bearer "+cli.token)
		events, body, errs := request.Clone().Get(provUrl + "/v1/envs").End()
		if errs != nil || events == nil {
			if events != nil {
				log.Println(events.StatusCode)
			}
			err := fmt.Errorf("error calling %s", provUrl)
			return "", "", err
		}

		if events.StatusCode == 200 {
			var raw map[string]interface{}
			_ = json.Unmarshal([]byte(body), &raw)
			data := raw["data"].(map[string]interface{})
			cli.url = "https://" + data["ese_url"].(string)
		}

		return cli.token, cli.url, nil
	}

	return "", "", fmt.Errorf("request failed. status: %s, response: %s", resp.Status, body)
}

func (cli *Client) makeRequest() *gorequest.SuperAgent {
	userAgent := fmt.Sprintf("%s/%s", UserAgentBase, version)
	return cli.gorequest.Clone().Set("User-Agent", userAgent)
}