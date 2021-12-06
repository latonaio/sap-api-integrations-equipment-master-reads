package main

import (
	sap_api_caller "sap-api-integrations-equipment-master-reads/SAP_API_Caller"
	"sap-api-integrations-planned-equipment-master-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Equipment_Master_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

    caller.AsyncGetEquipmentMaster(
        inoutSDC.Equipment.Equipment,
        inoutSDC.Equipment.ValidityEndDate,
    )
}
