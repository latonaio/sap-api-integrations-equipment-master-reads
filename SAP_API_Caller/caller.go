package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetEquipment(Equipment, ValidityEndDate string) {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		c.Equipment(Equipment, ValidityEndDate)
		wg.Done()
	}()
	go func() {
		c.BusinessPartner(Equipment, ValidityEndDate)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Equipment(Equipment, ValidityEndDate string) {
	res, err := c.callEquipmentSrvAPIRequirement("Equipment", Equipment, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) BusinessPartner(Equipment, ValidityEndDate string) {
	res, err := c.callEquipmentSrvAPIRequirementBusinessPartner("Equipment(Equipment='{Equipment}',ValidityEndDate=datetime'{ValidityEndDate}')/to_Partner", Equipment, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}


func (c *SAPAPICaller) callEquipmentSrvAPIRequirement(api, Equipment, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Equipment, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Equipment eq '%s' and ValidityEndDate eq '%s'", Equipment, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callEquipmentSrvAPIRequirementBusinessPartner(api, Equipment, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_EQUIPMENT", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "Equipment, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("Equipment eq '%s' and ValidityEndDate eq '%s'", Equipment, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}