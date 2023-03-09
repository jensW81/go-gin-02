package main

import (
	"bytes"
	"context"
	"encoding/json"
	go_gin "golang.source-fellows.com/seminar/go-gin/v2"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	callPut(ctx)
	callGet(ctx)
}

func callGet(ctx context.Context) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/api/getautos", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.SetBasicAuth("foo", "bar")
	response, err := http.DefaultClient.Do(request)
	if err == nil {
		all, _ := io.ReadAll(response.Body)
		log.Println(string(all))
	} else {
		log.Fatal(err)
	}
}

func callPut(ctx context.Context) {

	/*body := bytes.NewBufferString(`
		{
			"Kennzeichen": "Fzg from callPut"
		}
	`)*/
	audi := go_gin.Audi{Kennzeichen: "marshall"}
	marshal, err := json.Marshal(audi)
	if err != nil {
		log.Fatal(err)
		return
	}

	request, err := http.NewRequestWithContext(ctx,
		http.MethodPut,
		"http://localhost:8080/api/addauto",
		bytes.NewBuffer(marshal))
	if err != nil {
		log.Fatal(err)
	}

	request.SetBasicAuth("foo", "bar")
	request.Header.Add("x-trace-id", "123456")
	response, err := http.DefaultClient.Do(request)
	if err == nil {
		all, _ := io.ReadAll(response.Body)
		log.Println(string(all))
	} else {
		log.Fatal(err)
	}
}
