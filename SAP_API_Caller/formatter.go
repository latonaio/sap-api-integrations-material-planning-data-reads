package sap_api_caller

type MaterialPlanningDataReads struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	APISchema            string `json:"api_schema"`
	MaterialCode         string `json:"material_code"`
	MRPPlant             string `json:"plant"`
	Deleted              string `json:"deleted"`    
}

type MaterialPlanningData struct {
    MRPElementDocumentType         string      `json:"MRPElementDocumentType"`
    Material                       string      `json:"Material"`
    MRPPlant                       string      `json:"MRPPlant"`
    MRPArea                        string      `json:"MRPArea"`
    MRPPlanningSegment             string      `json:"MRPPlanningSegment"`
    MRPPlanningSegmentType         string      `json:"MRPPlanningSegmentType"`
    MRPElement                     string      `json:"MRPElement"`
    MRPElementItem                 string      `json:"MRPElementItem"`
    HasAcceptedShortage            string      `json:"HasAcceptedShortage"`
    TimeHorizonInDays              int         `json:"TimeHorizonInDays"`
    MaterialShortageDuration       int         `json:"MaterialShortageDuration"`
    MaterialShortageDurnInWorkdays int         `json:"MaterialShortageDurnInWorkdays"`
    MRPController                  string      `json:"MRPController"`
    MRPSafetyDuration              float64     `json:"MRPSafetyDuration"`
    DaysOfSupplyDuration           int         `json:"DaysOfSupplyDuration"`
    MaterialBaseUnit               string      `json:"MaterialBaseUnit"`
    MRPElementAvailyOrRqmtDate     string      `json:"MRPElementAvailyOrRqmtDate"`
    MRPSafetyDurationEndDate       string      `json:"MRPSafetyDurationEndDate"`
    ProductionVersion              string      `json:"ProductionVersion"`
    StorageLocation                string      `json:"StorageLocation"`
    ExceptionMessageNumber         string      `json:"ExceptionMessageNumber"`
    ExceptionMessageNumber2        string      `json:"ExceptionMessageNumber2"`
    MRPAvailableQuantity           float64     `json:"MRPAvailableQuantity"`
    MaterialShortageQuantity       float64     `json:"MaterialShortageQuantity"`
    MaterialShortageEndDate        string      `json:"MaterialShortageEndDate"`
    MaterialShortageStartDate      string      `json:"MaterialShortageStartDate"`
    MRPElementOpenQuantity         float64     `json:"MRPElementOpenQuantity"`
    MRPElementQuantityIsFirm       string      `json:"MRPElementQuantityIsFirm"`
    MaterialSafetyStockQty         float64     `json:"MaterialSafetyStockQty"`
    MRPElementBusinessPartner      string      `json:"MRPElementBusinessPartner"`
    MRPElementIsPartiallyDelivered string      `json:"MRPElementIsPartiallyDelivered"`
    MRPElementIsReleased           string      `json:"MRPElementIsReleased"`
}