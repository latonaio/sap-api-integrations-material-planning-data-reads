package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-planning-data-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToMRPItem(raw []byte, l *logger.Logger) ([]MRPItem, error) {
	pm := &responses.MRPItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MRPItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	mrpItem := make([]MRPItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		mrpItem = append(mrpItem, MRPItem{
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
		})
	}

	return mrpItem, nil
}
