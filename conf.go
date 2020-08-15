package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func FromFile(path string) (*Data, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromBytes(raw)
}

func FromString(s string) (*Data, error) {
	return FromBytes([]byte(s))
}

func FromBytes(raw []byte) (*Data, error) {
	var value map[string]interface{}
	err := yaml.Unmarshal(raw, &value)
	if err != nil {
		return nil, err
	}
	return &Data{value: value}, nil
}

func FromMap(value map[string]interface{}) *Data {
	return &Data{value: value}
}

type Data struct {
	value map[string]interface{}
}

func (data *Data) String() string {
	y, err := yaml.Marshal(data.value)
	if err != nil {
		return fmt.Sprintf("<conf: cannot marshal: %v>", err)
	}
	return string(y)
}

func (data *Data) GetString(key string, defaultValue string) string {
	if v, ok := data.value[key]; ok {
		switch s := v.(type) {
		case string:
			return s
		default:
			return fmt.Sprint(v)
		}
	}
	return defaultValue
}

func (data *Data) GetInt(key string, defaultValue int) int {
	if v, ok := data.value[key]; ok {
		switch n := v.(type) {
		case float64:
			return int(n)
		default:
			panic("whoops!")
		}
	}
	return defaultValue
}

func (data *Data) GetBool(key string, defaultValue bool) bool {
	if v, ok := data.value[key]; ok {
		switch b := v.(type) {
		case bool:
			return b
		case string:
			return "true" == b
		default:
			panic("whoops!")
		}
	}
	return defaultValue
}

func AsStrings(slice []interface{}) []string {
	var result []string
	for _, v := range slice {
		if s, ok := v.(string); ok {
			result = append(result, s)
		} else {
			return nil
		}
	}
	return result
}

func (data *Data) GetStrings(key string, defaultValue []string) []string {
	if v, ok := data.value[key]; ok {
		switch tv := v.(type) {
		case []interface{}:
			return AsStrings(tv)
		case []string:
			return tv
		default:
			panic("whoops!")
		}
	}
	return defaultValue
}
