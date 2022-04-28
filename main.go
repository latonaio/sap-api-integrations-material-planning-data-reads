package main

import (
	sap_api_caller "sap-api-integrations-material-planning-data-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-material-planning-data-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Material_Planning_Data_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"MaterialPlanningData",
		}
	}

	caller.AsyncGetMaterialPlanningData(
		inoutSDC.MaterialPlanningData.MRPElementDocumentType,
		inoutSDC.MaterialPlanningData.Material,
		inoutSDC.MaterialPlanningData.MRPPlant,
		accepter,
	)
}
