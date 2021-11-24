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


func (c *SAPAPICaller) AsyncGetMaterialPlanningData(MRPElementDocumentType, Material, MRPPlant string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c.MaterialPlanningData(MRPElementDocumentType, Material, MRPPlant)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) MaterialPlanningData(MRPElementDocumentType, Material, MRPPlant string) {
	res, err := c.callMaterialPlanningDataSrvAPIRequirement("SupplyDemandItems", MRPElementDocumentType, Material, MRPPlant)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callMaterialPlanningDataSrvAPIRequirement(api, MRPElementDocumentType, Material, MRPPlant string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_MRP_MATERIALS_SRV_01", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "MRPElementDocumentType, Material, MRPPlant")
	params.Add("$filter", fmt.Sprintf("MRPElementDocumentType eq '%s' and Material eq '%s' and MRPPlant eq '%s'", MRPElementDocumentType, Material, MRPPlant))
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