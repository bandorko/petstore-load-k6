package main

import (
	"context"
	"log"
	"net/http"

	petstore "github.com/Masato516/petstore-client-go"
	"github.com/bandorko/petstore-load-k6/pkg/provider"
)

func Setup() {
	config := petstore.NewConfiguration()
	config.Scheme = "https"
	config.Host = "petstore.swagger.io"
	client := petstore.NewAPIClient(config)
	resp, r, _ := client.PetApi.FindPetsByStatus(context.Background()).Status([]string{"available"}).Execute()
	provider.SetPetID(*resp[0].Id)
}

func Default(data interface{}) {
	config := petstore.NewConfiguration()
	config.Scheme = "https"
	config.Host = "petstore.swagger.io"
	config.HTTPClient = http.DefaultClient
	client := petstore.NewAPIClient(config)

	neededPetID := provider.GetPetID()

	resp, r, err := client.PetApi.GetPetById(context.Background(), neededPetID).Execute()
	if err != nil {
		log.Println(err)
		return
	}
	if r.StatusCode != 200 {
		log.Println("Status code is not 200 : ", r.StatusCode)
		return
	}
	if *resp.Id != neededPetID {
		log.Println("Didn't get the needed pet:", r.Body)
	}
}
