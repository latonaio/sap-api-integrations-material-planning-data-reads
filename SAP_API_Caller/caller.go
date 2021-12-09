package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-material-planning-data-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
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

func (c *SAPAPICaller) AsyncGetMaterialPlanningData(mRPElementDocumentType, material, mRPPlant string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	func() {
		c.MRPItem(mRPElementDocumentType, material, mRPPlant)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) MRPItem(mRPElementDocumentType, material, mRPPlant string) {
	data, err := c.callMaterialPlanningDataSrvAPIRequirementMRPItem("SupplyDemandItems", mRPElementDocumentType, material, mRPPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMaterialPlanningDataSrvAPIRequirementMRPItem(api, mRPElementDocumentType, material, mRPPlant string) ([]sap_api_output_formatter.MRPItem, error) {
	url := strings.Join([]string{c.baseURL, "API_MRP_MATERIALS_SRV_01", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMRPItem(req, mRPElementDocumentType, material, mRPPlant)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMRPItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithMRPItem(req *http.Request, mRPElementDocumentType, material, mRPPlant string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("MRPElementDocumentType eq '%s' and Material eq '%s' and MRPPlant eq '%s'", mRPElementDocumentType, material, mRPPlant))
	req.URL.RawQuery = params.Encode()
}
