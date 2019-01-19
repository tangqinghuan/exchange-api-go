package okex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/wanluc/exchange-api-go/utils"

	"github.com/pkg/errors"
)

const (
	// RestURL ...
	RestURL = "https://www.okex.com"
)

// Rest ...
type Rest interface {
	SetClient(client *http.Client)
	SetContext(ctx context.Context)
	Authenticate(key, secret, passphrase string)
	SetSigner(s func(key []byte) utils.Signer)
	Request(method, path string, params map[string]string, data interface{}, auth bool) ([]byte, error)
	Accounts() ([]*Account, error)
	Account(currency string) (*Account, error)
	Bills(currency, fromID, toID string, limit int) ([]*Bill, error)
	NewOrder(req *OrderRequest) (*OrderResponse, error)
	BatchNewOrder(req []*OrderRequest) (map[string][]*OrderResponse, error)
	CancelOrder(instrumentID, clientOID, orderID string) (*OrderResponse, error)
	BatchCancelOrder(req []*BatchCancelOrderRequest) (map[string]*BatchCancelOrderResponse, error)
	Candles(symbol string, granularity int32, start, end *time.Time) ([]*Candle, error)
}

type rest struct {
	baseURL    string
	key        string
	secret     string
	passphrase string
	signer     utils.Signer
	client     *http.Client
	ctx        context.Context
}

// NewRest ...
func NewRest(baseURL string) Rest {
	if baseURL == "" {
		baseURL = RestURL
	}
	return &rest{
		baseURL: baseURL,
		client:  &http.Client{},
		ctx:     context.Background(),
	}
}

// SetClient ...
func (r *rest) SetClient(client *http.Client) {
	r.client = client
}

// SetContext ...
func (r *rest) SetContext(ctx context.Context) {
	r.ctx = ctx
}

// Authenticate ...
func (r *rest) Authenticate(key, secret, passphrase string) {
	r.key = key
	r.secret = secret
	r.passphrase = passphrase
	r.signer = utils.NewHmacSha256Base64Signer([]byte(secret))
}

// SetSigner ....
func (r *rest) SetSigner(s func(key []byte) utils.Signer) {
	r.signer = s([]byte(r.secret))
}

// Request ...
func (r *rest) Request(method, path string, params map[string]string, data interface{}, auth bool) ([]byte, error) {
	reqURL := r.baseURL + path
	var reqData []byte
	var err error
	if data != nil {
		reqData, err = json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "json marshal request data")
		}
	}
	request, err := http.NewRequest(method, reqURL, bytes.NewReader(reqData))
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}
	request.WithContext(r.ctx)

	q := request.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	request.URL.RawQuery = q.Encode()

	if auth {
		if r.signer == nil {
			return nil, fmt.Errorf("please set the authentication info")
		}
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
		payload := timestamp + method + request.URL.RequestURI() + string(reqData)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("OK-ACCESS-KEY", r.key)
		request.Header.Set("OK-ACCESS-SIGN", r.signer.Sign([]byte(payload)))
		request.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
		request.Header.Set("OK-ACCESS-PASSPHRASE", r.passphrase)
	}

	log.Printf("request url : %s\n", request.URL.String())

	client := r.client
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer response.Body.Close()

	// 解析响应内容
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if response.StatusCode != http.StatusOK {
		var retErr ErrResponse
		if err := json.Unmarshal(content, &retErr); err != nil {
			return nil, errors.Wrap(err, "response body")
		}
		return nil, retErr
	}

	return content, err
}
