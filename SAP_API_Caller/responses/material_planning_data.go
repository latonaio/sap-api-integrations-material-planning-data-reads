package responses

type MaterialPlanningData struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MRPElementDocumentType         string `json:"MRPElementDocumentType"`
			Material                       string `json:"Material"`
			MRPPlant                       string `json:"MRPPlant"`
			MRPArea                        string `json:"MRPArea"`
			MRPPlanningSegment             string `json:"MRPPlanningSegment"`
			MRPPlanningSegmentType         string `json:"MRPPlanningSegmentType"`
			MRPElement                     string `json:"MRPElement"`
			MRPElementItem                 string `json:"MRPElementItem"`
			HasAcceptedShortage            bool   `json:"HasAcceptedShortage"`
			TimeHorizonInDays              string `json:"TimeHorizonInDays"`
			MaterialShortageDuration       string `json:"MaterialShortageDuration"`
			MaterialShortageDurnInWorkdays string `json:"MaterialShortageDurnInWorkdays"`
			MRPController                  string `json:"MRPController"`
			MRPSafetyDuration              string `json:"MRPSafetyDuration"`
			DaysOfSupplyDuration           string `json:"DaysOfSupplyDuration"`
			MaterialBaseUnit               string `json:"MaterialBaseUnit"`
			MRPElementAvailyOrRqmtDate     string `json:"MRPElementAvailyOrRqmtDate"`
			MRPSafetyDurationEndDate       string `json:"MRPSafetyDurationEndDate"`
			ProductionVersion              string `json:"ProductionVersion"`
			StorageLocation                string `json:"StorageLocation"`
			ExceptionMessageNumber         string `json:"ExceptionMessageNumber"`
			ExceptionMessageNumber2        string `json:"ExceptionMessageNumber2"`
			MRPAvailableQuantity           string `json:"MRPAvailableQuantity"`
			MaterialShortageQuantity       string `json:"MaterialShortageQuantity"`
			MaterialShortageEndDate        string `json:"MaterialShortageEndDate"`
			MaterialShortageStartDate      string `json:"MaterialShortageStartDate"`
			MRPElementOpenQuantity         string `json:"MRPElementOpenQuantity"`
			MRPElementQuantityIsFirm       bool   `json:"MRPElementQuantityIsFirm"`
			MaterialSafetyStockQty         string `json:"MaterialSafetyStockQty"`
			MRPElementBusinessPartner      string `json:"MRPElementBusinessPartner"`
			MRPElementIsPartiallyDelivered bool   `json:"MRPElementIsPartiallyDelivered"`
			MRPElementIsReleased           bool   `json:"MRPElementIsReleased"`
		} `json:"results"`
	} `json:"d"`
}
