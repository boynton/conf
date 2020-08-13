package conf

import (
	"testing"
)

var test1 = `
interface: 127.0.0.1
http-port: 8080
https-port: 4443
domains:
- one
- two
- three
`

func TestConf(test *testing.T) {
	conf, err := FromString(test1)
	if err != nil {
		test.Errorf("%v", err)
	}
	if conf.GetString("interface", "0.0.0.0") != "127.0.0.1" {
		test.Errorf("Failure, getString(key) didn't return the expected value")
	}
	if conf.GetString("http-port", "80") != "8080" {
		test.Errorf("Failure, getString(key) didn't return the expected value")
	}
	if conf.GetInt("http-port", 80) != 8080 {
		test.Errorf("Failure, getInt(key) didn't return the expected value")
	}
	lst := conf.GetStrings("domains", nil)
	if lst == nil {
		test.Errorf("Cannot get slice of strings")
	} else {
		if len(lst) != 3 {
			test.Errorf("Cannot get correct length slice of strings")
		} else {
			ref := map[int]string{0: "one", 1: "two", 2: "three"}
			for i, s := range lst {
				r := ref[i]
				if r != s {
					test.Errorf("Mismatch on item %d of list of strings: got %q, expected %q", i, s, r)
				}
			}
		}
	}
}
