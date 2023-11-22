package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	grpc_app "rus-profile-test/internal/app/grpc"
	rest_app "rus-profile-test/internal/app/rest"
	"rus-profile-test/internal/config"
	"rus-profile-test/internal/domain/profiler"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	cfg.Swagger.FilePath = "../../api/profile_v1/service.swagger.json"
	mainApp := rest_app.NewApp(cfg)
	grpcApp := grpc_app.NewApp(cfg)

	go func() {
		err := mainApp.Run()
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := grpcApp.Run()
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 2)
	exitVal := m.Run()
	os.Exit(exitVal)

}

func TestGetCompanyA(t *testing.T) {
	cfg := config.GetConfig()
	profileResponse := profiler.Response{}
	get, err := http.Get(fmt.Sprintf("http://localhost:%d/find/inn/%s", cfg.REST.Port, "6514007911"))
	if err != nil {
		t.Errorf("%s", err)
	}
	err = json.NewDecoder(get.Body).Decode(&profileResponse)
	if err != nil {
		t.Errorf("%x", err)
	}
	if profileResponse.DirectorName != "Мкоян Татос Тельмани" {
		t.Errorf("Incorrect DirectorName - Expected Мкоян Татос Тельмани, get: %s", profileResponse.DirectorName)
	}
	if profileResponse.CompanyName != "ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ \"АРМЕН\"" {
		t.Errorf("Incorrect CompanyName - Expected ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ \"АРМЕН\", get: %s", profileResponse.CompanyName)
	}
	if profileResponse.Kpp != "651401001" {
		t.Errorf("Incorrect Kpp - Expected 651401001, get: %s", profileResponse.Kpp)
	}
	if profileResponse.Inn != "6514007911" {
		t.Errorf("Incorrect Inn - Expected 6514007911, get: %s", profileResponse.Inn)
	}

}

func TestGetIncorrectCompany(t *testing.T) {
	cfg := config.GetConfig()
	get, err := http.Get(fmt.Sprintf("http://localhost:%d/find/inn/%s", cfg.REST.Port, "1234560977"))
	if err != nil {
		t.Errorf("%s", err)
	}
	if get.StatusCode != 404 {
		t.Errorf("Incorrect Response Code- Expected 404, get: %d", get.StatusCode)
	}
}

func TestMalformedInn(t *testing.T) {
	cfg := config.GetConfig()
	get, err := http.Get(fmt.Sprintf("http://localhost:%d/find/inn/%s", cfg.REST.Port, "asdsadadsa"))

	if err != nil {
		t.Errorf("%s", err)
	}
	if get.StatusCode != 400 {
		t.Errorf("Incorrect Response Code- Expected 400, get: %d", get.StatusCode)
	}
}
