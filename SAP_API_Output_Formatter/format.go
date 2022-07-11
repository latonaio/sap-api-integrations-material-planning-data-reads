package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-planning-data-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func ConvertToMaterialPlanningData(raw []byte, l *logger.Logger) *MaterialPlanningData {
	pm := &responses.MaterialPlanningData{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &MaterialPlanningData{
		MRPElementDocumentType:         data.MRPElementDocumentType,
		Material:                       data.Material,
		MRPPlant:                       data.MRPPlant,
		MRPArea:                        data.MRPArea,
		MRPPlanningSegment:             data.MRPPlanningSegment,
		MRPPlanningSegmentType:         data.MRPPlanningSegmentType,
		MRPElement:                     data.MRPElement,
		MRPElementItem:                 data.MRPElementItem,
		HasAcceptedShortage:            data.HasAcceptedShortage,
		TimeHorizonInDays:              data.TimeHorizonInDays,
		MaterialShortageDuration:       data.MaterialShortageDuration,
		MaterialShortageDurnInWorkdays: data.MaterialShortageDurnInWorkdays,
		MRPController:                  data.MRPController,
		MRPSafetyDuration:              data.MRPSafetyDuration,
		DaysOfSupplyDuration:           data.DaysOfSupplyDuration,
		MaterialBaseUnit:               data.MaterialBaseUnit,
		MRPElementAvailyOrRqmtDate:     data.MRPElementAvailyOrRqmtDate,
		MRPSafetyDurationEndDate:       data.MRPSafetyDurationEndDate,
		ProductionVersion:              data.ProductionVersion,
		StorageLocation:                data.StorageLocation,
		ExceptionMessageNumber:         data.ExceptionMessageNumber,
		ExceptionMessageNumber2:        data.ExceptionMessageNumber2,
		MRPAvailableQuantity:           data.MRPAvailableQuantity,
		MaterialShortageQuantity:       data.MaterialShortageQuantity,
		MaterialShortageEndDate:        data.MaterialShortageEndDate,
		MaterialShortageStartDate:      data.MaterialShortageStartDate,
		MRPElementOpenQuantity:         data.MRPElementOpenQuantity,
		MRPElementQuantityIsFirm:       data.MRPElementQuantityIsFirm,
		MaterialSafetyStockQty:         data.MaterialSafetyStockQty,
		MRPElementBusinessPartner:      data.MRPElementBusinessPartner,
		MRPElementIsPartiallyDelivered: data.MRPElementIsPartiallyDelivered,
		MRPElementIsReleased:           data.MRPElementIsReleased,
	}
}
