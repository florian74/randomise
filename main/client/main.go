package main

import (
	"context"
	"fmt"
	entities "github.com/florian74/randomise/entities"
	grpc "google.golang.org/grpc"
	"io"
	"os"
	"strings"
)

func main() {

	defaultType := "pets"
	defaultFields := []string{}

	if len(os.Args) > 1 {
		defaultType = "json"
		defaultFields = strings.Split(os.Args[1], ",")
	}

	dial, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		panic("could not start client" + err.Error())
	}

	if client := entities.NewRandomiseClient(dial); client != nil {

		streamClient, err := client.RandomStream(context.Background(), &entities.CommonRequest{ResponseType: defaultType, ResponseFields: defaultFields})
		if err != nil {
			panic("could not create stream client")
		}

		count := 0
		for {
			count = count + 1
			recv, err := streamClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Errorf("server exit with error %v", recv)
			}
			if defaultType != "json" {
				os.WriteFile(fmt.Sprintf("file%d", count), recv.Response, 0666)
			} else {
				fmt.Printf("json response is %s\n", string(recv.Response))
			}
		}

	}

}
