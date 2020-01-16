package helpers

import (
	"testing"
)

func Test_GenerateFromPassword(t *testing.T) {
	hash, err := GenerateFromPassword([]byte("MyPassword"))

	if err != nil {
		t.Error(err)
	} else {
		t.Log("GenerateFromPassword is ", string(hash), ", len:", len(string(hash)))
	}
}

func Test_CompareHashAndPassword(t *testing.T) {
	str := "524288$1$1$7c203a63d72384f8fc7accee77163fbd$df44d46c548910301d3c9d28a7e88c63424e5f6ee28a127fe8b79a35f6cf6738"
	err := CompareHashAndPassword([]byte(str), []byte("MyPassword"))
	if err != nil {
		t.Error(err)
	} else {
		t.Log("CompareHashAndPassword is OK")
	}
}
