package utils_test

import (
	"testing"

	"github.com/wanluc/exchange-api-go/utils"
)

var key = []byte("NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j")
var payload = []byte("symbol=LTCBTC&side=BUY&type=LIMIT&quantity=1&price=0.1")

func TestHmacSha256HexSigner(t *testing.T) {
	signer := utils.NewHmacSha256HexSigner(key)
	s := signer.Sign(payload)
	t.Log(s)
}

func TestHmacSha512HexSigner(t *testing.T) {
	signer := utils.NewHmacSha512HexSigner(key)
	s := signer.Sign(payload)
	t.Log(s)
}

func TestHmacSha256Base64Signer(t *testing.T) {
	signer := utils.NewHmacSha256Base64Signer(key)
	s := signer.Sign(payload)
	t.Log(s)
}
