package signature_test

import (
	"testing"

	"github.com/ReanSn0w/go-tinkoff-merchant/lib/utils/signature"
)

func Test_MakeSignature(t *testing.T) {
	pairs := []struct {
		Value Data
		Hash  string
	}{
		{
			Value: Data{Name: "Dmitriy", Age: 28, Desc: "some user"},
			Hash:  "c1a1aec42244abe993b191ae2fc61305",
		},
	}

	for _, item := range pairs {
		s := signature.MakeSignature(item.Value)

		if s != item.Hash {
			t.Logf("signature not setted: \n value: %s \n hash: %s", s, item.Hash)
			t.Fail()
		}
	}
}

// MARK: - support

type Data struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Desc      string `json:"desc"`
	Signature []byte `json:"signature"`
}

func (d *Data) SetSignature(value []byte) {
	d.Signature = value
}
