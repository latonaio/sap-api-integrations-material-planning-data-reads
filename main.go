package main

import (
    sap_api_caller "sap-api-integrations-material-planning-data-reads/SAP_API_Caller"
    "sap-api-integrations-material-planning-data-reads/file_reader"

    "github.com/latonaio/golang-logging-library/logger"
)

accepter := inoutSDC.Accepter
if len(accepter) == 0 || accepter[0] == "All" {
    accepter = []string{
        "MaterialPlanningData",
    }
}

func main() {
    l := logger.NewLogger()
    fr := file_reader.NewFileReader()
    inoutSDC := fr.ReadSDC("./Inputs//SDC_Material_Planning_Data_sample.json")
    caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
    )

    caller.AsyncGetMaterialPlanningData(
        inoutSDC.MRPElementDocumentType.Material.MRPPlant,
        accepter,
    )
}