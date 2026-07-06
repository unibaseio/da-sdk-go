package sdk

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/gorilla/websocket"
)

func WSConnect(baseURL string, au types.Auth) (*websocket.Conn, error) {
	url := strings.TrimPrefix(baseURL, "http")
	url = "ws" + url + "/v1/ws"

	aub, err := json.Marshal(au)
	if err != nil {
		return nil, err
	}

	header := make(http.Header)

	header.Add("Authorization", hex.EncodeToString(aub))
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
