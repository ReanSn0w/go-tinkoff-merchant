package signature

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
)

type Signature interface {
	SetSignature([]byte)
}

func MakeSignature(data any) string {
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	count := dataType.NumField()

	pairs := []signaturePair{}

	for i := 0; i < count; i++ {
		value := ""
		t := dataValue.Field(i).Type().String()
		switch t {
		case "bool":
			value = fmt.Sprint(dataValue.Field(i).Bool())
		case "string":
			value = dataValue.Field(i).String()
		case "int":
			value = fmt.Sprint(dataValue.Field(i).Int())
		case "float32", "float64":
			value = fmt.Sprint(dataValue.Field(i).Float())
		case "uint":
			value = fmt.Sprint(dataValue.Field(i).Uint())
		default:
			continue
		}

		pairs = append(pairs, signaturePair{
			Key:   dataType.Field(i).Tag.Get("json"),
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

		stringForHash += item.Key + item.Value
	}

	hash := md5.Sum([]byte(stringForHash))
	return hex.EncodeToString(hash[:])
}

type signaturePair struct {
	Key   string
	Value string
}
