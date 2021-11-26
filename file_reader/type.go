package file_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	MRPElement    struct {
		MRPElement                     string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		MRPAvailableQuantity           float64     `json:"quantity"`
		PickedQuantity                 float64     `json:"picked_quantity"`
		Price                          float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string      `json:"api_schema"`
	MaterialCode                   string      `json:"material_code"`
	MRPPlant                       string      `json:"plant/supplier"`
	Stock                          float64     `json:"stock"`
	MRPElementDocumentType         string      `json:"document_type"`
	MRPElement                     string      `json:"document_no"`
	MRPElementAvailyOrRqmtDate     string      `json:"planned_date"`
	ValidatedDate                  string      `json:"validated_date"`
	Deleted                        string      `json:"deleted"`
}

type SDC struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	MaterialPlanningData struct {
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
	} `json:"MaterialPlanningData"`
	APISchema    string `json:"api_schema"`
	MaterialCode string `json:"material_code"`
	MRPPlant     string `json:"plant"`
	Deleted      string `json:"deleted"`
}