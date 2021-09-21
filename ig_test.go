package igslim

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetUser(t *testing.T) {
	client := NewClient(os.Getenv("IGSESSIONID"))
	user, err := client.GetUser("TaylorSwift")
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(user)
	if user.Id != 11830955 {
		t.Error("wrong user id")
	}
	if user.FbId != 17841401648650184 {
		t.Error("wrong user fbid")
	}
	if user.UserName != "taylorswift" {
		t.Error("wrong user name")
	}
	if user.FullName != "Taylor Swift" {
		t.Error("wrong user full name")
	}
}
