package random

import (
	"fmt"
	entities "github.com/florian74/randomise/entities"
	"strings"
	"testing"
)

func TestJsonGenerate(t *testing.T) {

	action, err := generateJsonAction(&entities.CommonRequest{
		ResponseFields: strings.Split("A,B", ","),
		ResponseType:   "json",
	})
	if err != nil {
		t.Fatalf("No result")
	}

	if !strings.Contains(string(action), "A") {
		t.Fatalf("Wrong result")
	}
	fmt.Printf("OK")

}
