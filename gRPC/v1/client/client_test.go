package client

import (
	"context"
	"testing"

	"github.com/TheLazarusNetwork/erebrus/model"
)

func TestCreateClient(t *testing.T) {
	var client *ClientService
	data := new(model.Client)
	data.Name = "sambath kumar"
	data.Tags = []string{"home"}
	data.CreatedBy = "sambath@mail.com"
	data.Email = "sambath@mail.com"
	data.Enable = true
	data.AllowedIPs = []string{"0.0.0.0/0", "::/0"}
	data.Address = []string{"10.0.0.1/24"}
	response, err := client.CreateClient(context.Background(), data)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Sucess")
		t.Log(response)
	}
}
