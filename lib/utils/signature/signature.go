package signature

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
)

type Signature interface {
	SetSignature([]byte)
}

func MakeSignature(data any, password string) string {
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	count := dataType.NumField()

	pairs := []signaturePair{
		{
			Key:   "Password",
			Value: password,
		},
	}

	for i := 0; i < count; i++ {
		value := ""
		t := dataValue.Field(i).Type().String()

		// Другие типы игнорируются так как в структурах для хначений используются только данные типы
		switch t {
		case "bool":
			value = fmt.Sprint(dataValue.Field(i).Bool())
		case "string":
			value = dataValue.Field(i).String()
		case "int64":
			value = fmt.Sprint(dataValue.Field(i).Int())
		default:
			continue
		}

		pairs = append(pairs, signaturePair{
			Key:   dataType.Field(i).Name,
			Value: value,
		})

		dataValue.Field(i)
	}

	signature := makeSignature(pairs)
	//data.SetSignature(signature)
	return signature
}

func makeSignature(pairs []signaturePair) string {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key < pairs[j].Key
	})

	stringForHash := ""

	for _, item := range pairs {
		if item.Value == "" {
			continue
		}

		stringForHash += item.Value
	}

	hash := sha256.New()
	hash.Write([]byte(stringForHash))
	value := hash.Sum(nil)
	return hex.EncodeToString(value[:])
}

type signaturePair struct {
	Key   string
	Value string
}
