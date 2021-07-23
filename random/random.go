package random

import (
	"context"
	"encoding/json"
	"fmt"
	entities "github.com/florian74/randomise/entities"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var instance *RandomServer
var lock = &sync.Mutex{}

func GetInstance() *RandomServer {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &RandomServer{
				actions: map[string]func(request *entities.CommonRequest) ([]byte, error){
					"pets": fetchDogPictureAction,
					"json": generateJsonAction,
				},
				UnimplementedRandomiseServer: entities.UnimplementedRandomiseServer{}}
		} else {
			fmt.Println("instance already created.")
		}
	} else {
		fmt.Println("instance already created.")
	}

	return instance
}

type RandomServer struct {
	actions map[string]func(request *entities.CommonRequest) ([]byte, error)
	entities.UnimplementedRandomiseServer
}

func (server RandomServer) RandomStream(request *entities.CommonRequest, stream entities.Randomise_RandomStreamServer) error {
	for i := 0; i < 4; i++ {
		resp, err := server.Random(context.Background(), request)
		if err != nil {
			fmt.Printf("cannot generate, %v", err)
			return err
		}
		err = stream.Send(resp)
		if err != nil {
			fmt.Printf("an error on send, %v", err)
			return err
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}

func (server RandomServer) Random(ctx context.Context, entity *entities.CommonRequest) (*entities.CommonResponse, error) {

	fmt.Printf("Random called\n")
	var result []byte
	var err error

	for key, action := range server.actions {
		if key == entity.ResponseType {
			result, err = action(entity)
			if err != nil {
				return nil, fmt.Errorf("actions failed %s, %v", key, err)
			}
		}
	}

	return &entities.CommonResponse{
		Response: result,
	}, nil
}

func generateJsonAction(request *entities.CommonRequest) ([]byte, error) {

	if request.ResponseFields == nil {
		return nil, fmt.Errorf("no field are given in parameters")
	}

	data := make(map[string]string)
	for i := 0; i < len(request.ResponseFields); i++ {

		array := []rune(request.ResponseFields[i])
		if len(array) == 0 {
			return nil, fmt.Errorf("a field cannot be empty")
		}

		if array[0] < 'A' && array[0] > 'Z' {
			array[0] = array[0] + 'A' - 'a'
		}

		data[strings.TrimSpace(string(array))] = randSeq(10)
	}

	return json.Marshal(data)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type dogInfo struct {
	Message string
	Status  string
}

func fetchDogPictureAction(request *entities.CommonRequest) ([]byte, error) {
	info, err := get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return nil, fmt.Errorf("cannot access web site, %v", err)
	}
	var dogInfo dogInfo
	err = json.Unmarshal(info, &dogInfo)
	if err != nil {
		return nil, fmt.Errorf("cannot read data, %v", err)
	}
	return get(dogInfo.Message)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s failed, %+v", url, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s returned %s", url, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to extract data from GET %s, %+v", url, err)
	}

	return data, nil
}
