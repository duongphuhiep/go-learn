package main

import (
	"encoding/json"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestJsonUnmarshal(t *testing.T) {
	rawResp := []byte(`{
	"q": [{
		"alias": "a",
		"balance": 100.00
	}]
}`)
	var resp struct {
		Q []struct {
			Alias   string  `json:"alias"`
			Balance float64 `json:"balance"`
		} `json:"q"`
	}

	if err := json.Unmarshal(rawResp, &resp); err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestJsonMarshal(t *testing.T) {
	var req struct {
		Query string `json:"query,omitempty"`
		Set   struct {
			Uid     string  `json:"uid,omitempty"`
			Balance float64 `json:"balance,omitempty"`
		} `json:"set"`
	}
	req.Query = `
		q(func: eq(Wallet.alias, "a")) {
			v as uid
		}`
	req.Set.Uid = "uid(v)"
	req.Set.Balance = 201

	reqBytes, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(reqBytes))
}
