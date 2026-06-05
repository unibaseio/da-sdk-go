package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/unibaseio/da-sdk-go/build"
)

var Validators = map[string]func(string, string) error{
	"heartbeat.nickname": validateLettersOnly,
}

type Config struct {
	Wallet WalletConfig `json:"wallet"`
	API    APIConfig    `json:"api"`
	Remote RemoteConfig `json:"remote"`
	Chain  ChainConfig  `json:"chain"`
	Store  StoreConfig  `json:"store"`
}

type WalletConfig struct {
	Address string `json:"address,omitempty"`
}

type APIConfig struct {
	Endpoint string `json:"endpoint"`
	Expose   string `json:"expose"`
}

func newDefaultAPIConfig() APIConfig {
	return APIConfig{
		Endpoint: "127.0.0.1:8081",
	}
}

type ChainConfig struct {
	Type string `json:"type"`
}

func newDefaultChainConfig() ChainConfig {
	return ChainConfig{
		Type: build.OPSepolia,
	}
}

type RemoteConfig struct {
	URL string `json:"url"`
}

func newDefaultRemoteConfig() RemoteConfig {
	return RemoteConfig{
		URL: "http://127.0.0.1:8080",
	}
}

type StoreConfig struct {
	Stat    bool  `json:"stat"` // store new data or not
	MaxSize int64 `json:"max"`
}

func newDefaultStoreConfig() StoreConfig {
	return StoreConfig{
		Stat:    true,
		MaxSize: 1024 * 1024 * 1024 * 1024,
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		API:    newDefaultAPIConfig(),
		Remote: newDefaultRemoteConfig(),
		Chain:  newDefaultChainConfig(),
		Store:  newDefaultStoreConfig(),
	}
}

func (cfg *Config) WriteFile(file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close() // nolint: errcheck

	configString, err := json.MarshalIndent(*cfg, "", "\t")
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(f, string(configString))
	return err
}

func ReadFile(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	cfg := NewDefaultConfig()
	rawConfig, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if len(rawConfig) == 0 {
		return cfg, nil
	}

	err = json.Unmarshal(rawConfig, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) Set(dottedKey string, jsonString string) error {
	if !json.Valid([]byte(jsonString)) {
		jsonBytes, _ := json.Marshal(jsonString)
		jsonString = string(jsonBytes)
	}

	err := validate(dottedKey, jsonString)
	if err != nil {
		return err
	}

	keys := strings.Split(dottedKey, ".")
	for i := len(keys) - 1; i >= 0; i-- {
		jsonString = fmt.Sprintf(`{ "%s": %s }`, keys[i], jsonString)
	}

	decoder := json.NewDecoder(strings.NewReader(jsonString))
	decoder.DisallowUnknownFields()

	return decoder.Decode(&cfg)
}

func (cfg *Config) Get(key string) (interface{}, error) {
	v := reflect.Indirect(reflect.ValueOf(cfg))
	keyTags := strings.Split(key, ".")
OUTER:
	for j, keyTag := range keyTags {
		if v.Type().Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				jsonTag := strings.Split(
					v.Type().Field(i).Tag.Get("json"),
					",")[0]
				if jsonTag == keyTag {
					v = v.Field(i)
					if j == len(keyTags)-1 {
						return v.Interface(), nil
					}
					v = reflect.Indirect(v) // only attempt one dereference
					continue OUTER
				}
			}
		}

		return nil, fmt.Errorf("key: %s invalid for config", key)
	}
	// Cannot get here as len(strings.Split(s, sep)) >= 1 with non-empty sep
	return nil, fmt.Errorf("empty key is invalid")
}

// validate runs validations on a given key and json string. validate uses the
// validators map defined at the top of this file to determine which validations
// to use for each key.
func validate(dottedKey string, jsonString string) error {
	var obj interface{}
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		return err
	}
	// recursively validate sub-keys by partially unmarshalling
	if reflect.ValueOf(obj).Kind() == reflect.Map {
		var obj map[string]json.RawMessage
		err := json.Unmarshal([]byte(jsonString), &obj)
		if err != nil {
			return err
		}
		for key := range obj {
			err := validate(dottedKey+"."+key, string(obj[key]))
			if err != nil {
				return err
			}
		}
		return nil
	}

	if validationFunc, present := Validators[dottedKey]; present {
		return validationFunc(dottedKey, jsonString)
	}

	return nil
}

// validateLettersOnly validates that a given value contains only letters. If it
// does not, an error is returned using the given key for the message.
func validateLettersOnly(key string, value string) error {
	if match, _ := regexp.MatchString("^\"[a-zA-Z]+\"$", value); !match {
		return fmt.Errorf(`"%s" must only contain letters`, key)
	}
	return nil
}
