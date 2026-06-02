package test

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/unibaseio/da-sdk-go/sdk"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

const url = "http://127.0.0.1:8086"

var jsonaddr = uuid.New().String()
var dataaddr = uuid.New().String()

type JsonStruct struct {
	Name  string
	Age   int
	Value string
}

func TestJson(t *testing.T) {
	js := JsonStruct{
		Name:  "test",
		Age:   10,
		Value: "aa",
	}
	jsb, err := json.Marshal(js)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsb))
	jsbi, err := json.MarshalIndent(js, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsbi))
	t.Fatal()
}

func TestList(t *testing.T) {
	err := sdk.ListAccountHub(url)
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.ListNeedleHub(url, jsonaddr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadJson(t *testing.T) {
	js := JsonStruct{
		Name: "test",
		Age:  10,
	}
	jsb, err := json.Marshal(js)
	if err != nil {
		t.Fatal(err)
	}
	length := rand.Int31n(16) + 16
	nkey := utils.RandomBytes(int(length))
	err = sdk.UploadHub(url, jsonaddr, hex.EncodeToString(nkey), string(jsb))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		nkey := uuid.New().String()
		length = rand.Int31n(1024 * 1024)
		nval := utils.RandomBytes(int(length))

		js := JsonStruct{
			Name:  "test",
			Age:   10,
			Value: hex.EncodeToString(nval),
		}
		jsb, err := json.Marshal(js)
		if err != nil {
			t.Fatal(err)
		}

		err = sdk.UploadHub(url, jsonaddr, nkey, string(jsb))
		if err != nil {
			t.Fatal(err)
		}
	}

	err = sdk.ListAccountHub(url)
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.ListNeedleHub(url, jsonaddr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadData(t *testing.T) {

	for i := 0; i < 10; i++ {
		length := rand.Int31n(16) + 16
		nkey := utils.RandomBytes(int(length))
		length = rand.Int31n(5 * 1024 * 1024)
		nval := utils.RandomBytes(int(length))

		err := sdk.UploadHubData(url, dataaddr, hex.EncodeToString(nkey), nval)
		if err != nil {
			t.Fatal(err)
		}
	}
}
