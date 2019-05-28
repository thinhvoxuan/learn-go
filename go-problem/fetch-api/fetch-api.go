package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Company struct {
	Level     int    `json:"level"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsSponsor bool   `json:"isSponsor"`
	ID        string `json:"_id"`
	Shortname string `json:"shortname"`
	Fullname  string `json:"fullname"`
	Image     string `json:"image"`
}

type CompanyDetail struct {
	Level              int      `json:"level"`
	Width              int      `json:"width"`
	Height             int      `json:"height"`
	ContactEmails      []string `json:"contactEmails"`
	MaxAcceptedStudent int      `json:"maxAcceptedStudent"`
	MaxRegister        int      `json:"maxRegister"`
	AdminMaxRegister   int      `json:"adminMaxRegister"`
	IsSponsor          bool     `json:"isSponsor"`
	Active             bool     `json:"active"`
	StudentRegister    int      `json:"studentRegister"`
	StudentAccepted    int      `json:"studentAccepted"`
	ID                 string   `json:"_id"`
	Shortname          string   `json:"shortname"`
	Fullname           string   `json:"fullname"`
	Address            string   `json:"address"`
	V                  int      `json:"__v"`
	User               string   `json:"user"`
	Work               string   `json:"work"`
	Image              string   `json:"image"`
	InternshipFile     string   `json:"internshipFile"`
}

// List Response
type ListResponse struct {
	Items []Company `json:"items"`
}

// DetailResponse
type DetailResponse struct {
	Item CompanyDetail `json:"item"`
}

func fetchStatic(com Company) CompanyDetail {
	resp, _ := myClient.Get("http://cse.hcmut.edu.vn/new/home/company/id/" + com.ID)
	defer resp.Body.Close()

	var detailResponse DetailResponse
	json.NewDecoder(resp.Body).Decode(&detailResponse)
	return detailResponse.Item
}

func fetchAll() []Company {
	resp, _ := myClient.Get("http://cse.hcmut.edu.vn/new/home/company/all")
	defer resp.Body.Close()

	var listResponse ListResponse
	json.NewDecoder(resp.Body).Decode(&listResponse)
	return listResponse.Items
}

func main() {
	listCompany := fetchAll()
	println("ID", "|", "Fullname", "|", "Register", "|", "MaxRegister", "|", "MaxAccepted", "|", "Accepted")
	for _, c := range listCompany {
		cD := fetchStatic(c)
		println(c.ID, "|", c.Fullname, "|", cD.StudentRegister, "|", cD.MaxRegister, "|", cD.MaxAcceptedStudent, "|", cD.StudentAccepted)
	}
}
