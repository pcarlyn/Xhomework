package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

type Item struct {
	Caption string  `json:"caption"`
	Weight  float32 `json:"weight"`
	Number  int     `json:"number"`
}

func PostRequest(item *Item) string {

	url := "http://localhost:8080/item"

	payload := strings.NewReader(fmt.Sprintf(`{
		"caption": "%s",
		"weight": %f,
		"number": %d
	}`, item.Caption, item.Weight, item.Number))

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return fmt.Sprintf("Ошибка при создании запроса: %v", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Sprintf("Ошибка при выполнении запроса: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Sprintf("Ошибка при чтении тела ответа: %v", err)
	}

	return (string(body))

}

func GetRequest(caption string) string {

	url := "http://localhost:8080/item/" + caption

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Sprintf("Ошибка при создании запроса: %v", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Sprintf("Ошибка при выполнении запроса: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Sprintf("Ошибка при чтении тела ответа: %v", err)
	}

	return (string(body))
}

var names []string = []string{"Azat", "David", "Radmil", "Max", "Maxim", "Slava"}

func Random() float32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Float32()
}

func RandomWeight() float32 {
	return Random() * 100
}

func RandomNumber() int {
	return int(Random() * 150)
}

func RandomIndex() int {
	return int(Random() * float32(len(names)))
}

func GenerateItem() *Item {
	randomIndex := RandomIndex()
	return &Item{
		Caption: names[randomIndex],
		Weight:  RandomWeight(),
		Number:  RandomNumber(),
	}
}

func GeneratePostItem(count int) []string {
	captions := make([]string, 0)
	for i := 0; i < count; i++ {
		item := GenerateItem()
		captions = append(captions, item.Caption)
		PostRequest(item)
	}
	return captions
}

func GenerateAndGet() []string {
	response := make([]string, 0)
	captions := GeneratePostItem(RandomIndex())
	for i := 0; i < len(captions); i++ {
		response = append(response, GetRequest(captions[i]))
	}
	return response
}

func main() {
	GenerateAndGet()
	// fmt.Println(GenerateAndGet())
}
