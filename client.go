package chiarpc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseUrl = "https://localhost"
const (
	DaemonPort    = 55400
	FullNodePort  = 8555
	WalletPort    = 9256
	FarmerPort    = 8559
	HarvesterPort = 8560
	TestPort      = 8080
)

type Client struct {
	BaseUrl string
	client  *http.Client
}

type ClientSettings struct {
	PathToCertFile   string
	PathToCertSecret string
	BaseUrl          string
}

func NewRPCClient(settings ClientSettings) (*Client, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if settings.PathToCertFile == "" {
		settings.PathToCertFile = fmt.Sprintf("%s/.chia/mainnet/config/ssl/ca/private_ca.crt", homePath)
	}
	if settings.PathToCertSecret == "" {
		settings.PathToCertSecret = fmt.Sprintf("%s/.chia/mainnet/config/ssl/ca/private_ca.key", homePath)
	}
	if settings.BaseUrl == "" {
		settings.BaseUrl = baseUrl
	}

	cert, err := tls.LoadX509KeyPair(settings.PathToCertFile, settings.PathToCertSecret)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &Client{
		BaseUrl: settings.BaseUrl,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates:       []tls.Certificate{cert},
					InsecureSkipVerify: true,
				},
			},
		},
	}
	return client, err
}

func (c Client) makeRPCCall(method string, rpcMethod string, port uint16, data map[string]interface{}, queryParams map[string]string) ([]byte, error) {
	if method == "" {
		method = http.MethodPost
	}

	url := fmt.Sprintf("%s:%d/%s", c.BaseUrl, port, rpcMethod)

	if data == nil {
		data = map[string]interface{}{}
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// query param stuff will be nil for the time being as there is no need for query params
	var queryString string

	if queryParams != nil {
		q := req.URL.Query()
		for key, val := range queryParams {
			q.Add(key, val)
		}
		queryString = q.Encode()
		req.URL.RawQuery = queryString
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
