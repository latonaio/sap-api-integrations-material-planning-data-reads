# sap-api-integrations-material-planning-data-reads  
sap-api-integrations-material-planning-data-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で MRP データを取得するマイクロサービスです。    
sap-api-integrations-material-planning-data-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-material-planning-data-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_MRP_MATERIALS_SRV_01_0001/overview

## 動作環境
sap-api-integrations-material-planning-data-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-material-planning-data-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Product", "Plant", "Accounting" が指定されています。    
  
```
 "api_schema": "/SupplyDemandItems",
 "accepter": ["MaterialPlanningData"],
 "material_code": "100141",
 "plant": "1000",
 "deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
  "api_schema": "sap.s4.beh.product.v1.Product.Created.v1",
  "accepter": ["All"],
  "material_code": "21",
  "deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMaterialPlanningData(MRPElementDocumentType, Material, MRPPlant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialPlanningData":
			func() {
				c.MaterialPlanningData(MRPElementDocumentType, Material, MRPPlant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。   
以下の sample.json の例は、SAP MRPデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Batch" ～ "Supplier" は、/SAP_API_Output_Formatter/type.go 内 の type Batch struct {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。    

```
{
	"MRPElementDocumentType": "MD",
	"Material": "100141",
	"MRPPlant": "1000",
	"MRPArea": "1000",
	"MRPPlanningSegment": "",
	"MRPPlanningSegmentType": "",
	"MRPElement": "4500000073",
	"MRPElementItem": "10",
	"HasAcceptedShortage": "",
	"TimeHorizonInDays": "",
	"MaterialShortageDuration": "",
	"MaterialShortageDurnInWorkdays": "",
	"MRPController": "001",
	"MRPSafetyDuration": "999.9",
	"DaysOfSupplyDuration": "",
	"MaterialBaseUnit": "PC",
	"MRPElementAvailyOrRqmtDate": "",
	"MRPSafetyDurationEndDate": "",
	"ProductionVersion": "0001",
	"StorageLocation": "1000",
	"ExceptionMessageNumber": "",
	"ExceptionMessageNumber2": "",
	"MRPAvailableQuantity": "30",
	"MaterialShortageQuantity": "",
	"MaterialShortageEndDate": "",
	"MaterialShortageStartDate": "",
	"MRPElementOpenQuantity": "",
	"MRPElementQuantityIsFirm": false,
	"MaterialSafetyStockQty": "",
	"MRPElementBusinessPartner": "",
	"MRPElementIsPartiallyDelivered": false,
	"MRPElementIsReleased": false,
	"cursor": "/Users/kyotatashiro/go/src/sap-api-integrations-material-planning-data-reads/SAP_API_Caller/caller.go#L220",
	"function": "sap-api-integrations-material-planning-data-reads/SAP_API_Caller.(*SAPAPICaller).MaterialPlanningData",
	"level": "INFO",
	"time": "2021-11-20T18:28:36.143833+09:00"
}
```