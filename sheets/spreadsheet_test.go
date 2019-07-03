package sheets

import (
	"fmt"
	"testing"
)

func TestCreateContentOnGoogleSheet(t *testing.T){
	Initialize("client_secret.json","1M3lEfISeJrxv2C9yQRouihpIbx_n1n6s5l5RtRZRPIY" )
	entity := map[string]string{"Nome":"Tom","CPF":"36163719883","Email":"tom3@guru.com.vc","Telefone":"11947588847", "Valor":"240000"}
	err := AddContent("Dados Básicos", entity)
	if err != nil{
		t.Error("Erro! Não foi possível inserir informações na planilha! ", err)
	}
}

func TestGetContentOnGoogleSheet(t *testing.T){
	Initialize("client_secret.json","1M3lEfISeJrxv2C9yQRouihpIbx_n1n6s5l5RtRZRPIY" )
	entity, err := getRowByValueInColumn("Dados Básicos", "tom2@guru.com.vc", "Email")
	if err != nil{
		t.Error("Erro! Não foi possível inserir informações na planilha! ", err)
	}else{
		for key, value := range entity {
			fmt.Print(key)
			fmt.Println("-" + value)
		}
	}
}
