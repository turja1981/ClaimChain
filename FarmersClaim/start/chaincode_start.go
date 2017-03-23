/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
//	"github.com/hyperledger/fabric/core/chaincode/shim/table.pb"
	"fmt"
//	"net/http"
//	"io/ioutil"
//	"time"
	"strconv" 
	//"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"regexp"
)


var logger = shim.NewLogger("ClaimChaincode")


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type customEvent struct {
	Type       					string `json:"type"`
	Decription 					string `json:"description"`
	Owner	   					string `json:"owner"`
//	TransactionTime  			Time `json:"transactionTime"`
}



type Vehicle struct {
	Make            string `json:"make,omitempty"`
	Model           string `json:"model,omitempty"`
	VIN             string `json:"vin,omitempty"`
	Year           	string `json:"year,omitempty"`
}


type Loss struct {
	LossType            	string `json:"lossType,omitempty"`
	LossDateTime            string `json:"lossDateTime,omitempty"`
	LossDescription     	string `json:"lossDescription,omitempty"`
	LossAddress         	string `json:"lossAddress,omitempty"`
	LossCity            	string `json:"lossCity,omitempty"`
	LossState	    		string `json:"lossState,omitempty"`
	LossZipCode	    		string `json:"lossZipCode,omitempty"` 

} 

type Insured struct {
	FirstName              string `json:"firstName,omitempty"`
	LastName           	   string `json:"lastName,omitempty"`
	PhoneNo         	   string `json:"phoneNo,omitempty"`
	Email           	   string `json:"email,omitempty"`
	Dob             	   string `json:"dob,omitempty"`
	SSN             	   string `json:"ssn,omitempty"`
	DrivingLicense         string `json:"drivingLicense,omitempty"`
}


type Adjuster struct {
	 
	AdjusterZipCode			string		`json:"adjusterZipCode,omitempty"`
	AdjusterSpeciality		string		`json:"adjusterSpeciality,omitempty"`
	AdjusterFirstName		string		`json:"adjusterFirstName,omitempty"`
	AdjusterLastName		string		`json:"adjusterLastName,omitempty"`
	EvaluationDateTime		string		`json:"evaluationDateTime,omitempty"` 
	ApproveLossAmount		string		`json:"approveLossAmount,omitempty"` 
	Remarks	    			string		`json:"remarks,omitempty"`


}

type RepairShop struct {
	
	RepairShopName			string			`json:"repairShopName,omitempty"` 
	RepairZipCode			string			`json:"repairZipCode,omitempty"`	 
	RepairDateTime			string			`json:"repairDateTime,omitempty"` 
	ItemRepaired			[5]RepairItem	`json:"itemRepaired,omitempty"` 
	TotalCost	    		string			`json:"totalCost,omitempty"`


}

type RepairItem struct {
	 
	ItemId					string		`json:"itemId,omitempty"`
	ItemName				string		`json:"itemName,omitempty"` 
	ItemCost				string		`json:"itemCost,omitempty"` 

}



type Payment struct {
	 
	BankName				string		`json:"bankName,omitempty"` 
	AccountNo				string		`json:"accountNo,omitempty"` 
	PaymentAmount			string		`json:"paymentAmount,omitempty"` 
	PaymentDateTime	    	string		`json:"paymentDateTime,omitempty"`

}

type Claim struct {
	 
	ClaimNo	    				string		`json:"claimNo,omitempty"`	 
	PolicyNo					string		`json:"policyNo,omitempty"` 
	Status              		string      `json:"status,omitempty"`
	ExternalReport      		string      `json:"externalReport,omitempty"`
	LossDetails 				Loss 		`json:"lossDetails,omitempty"`
	InsuredDetails 				Insured 	`json:"insuredDetails,omitempty"`
	VehicleDetails 				Vehicle 	`json:"vehicleDetails,omitempty"`
	ThirdPartyInsuredDetails	Insured 	`json:"thirdPartyInsuredDetails,omitempty"`
	ThirdPartyVehicleDetails	Vehicle 	`json:"thirdPartyVehicleDetails,omitempty"`
	AdjusterReport 				Adjuster 	`json:"adjusterReport,omitempty"`
	RepairedDetails 			RepairShop 	`json:"repairedDetails,omitempty"`
	PaymentDetails 				Payment 	`json:"paymentDetails,omitempty"`
	SensorData 		    		Sensor 		`json:"sensorData,omitempty"`

}



type Sensor struct {
    Latitude    *string `json:"latitude,omitempty"`
    Longitude   *string `json:"longitude,omitempty"`
    Image   	*string `json:"image,omitempty"`
    Voice   	*string `json:"voice,omitempty"`
}


var  CLAIM_NO = 7000000
// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	debugLevel, _ := shim.LogLevel("DEBUG")
	shim.SetLoggingLevel(debugLevel)
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("Initialization Complete ")

	//err := t.createFraudTable(stub);
	
	logger.Debug("Initialization Complete ")
	
	return nil , nil
}

func (t *SimpleChaincode) readAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetLoanApplication")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing Claim No")
	}

	var claimNo = args[0]
	var c Claim
	bytes, err := stub.GetState(claimNo)
	
	err = json.Unmarshal(bytes, &c); 


	if err != nil {
		logger.Error("Could not fetch Claim application with No "+claimNo+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func (t *SimpleChaincode) readAssetObjectModel(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var claimObj Claim = Claim{}

    // Marshal and return
    stateJSON, err := json.Marshal(claimObj)
    if err != nil {
        return nil, err
    }
    return stateJSON, nil
}
//*************readAssetSamples*******************/

func (t *SimpleChaincode) readAssetSamples(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(samples), nil
}
//*************readAssetSchemas*******************/

func (t *SimpleChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(schemas), nil
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

// Handle different functions
    if function == "readAsset" {
        // gets the state for an assetID as a JSON struct
        return t.readAsset(stub, args)
    } else if function =="readAssetObjectModel" {
        return t.readAssetObjectModel(stub, args)
    }  else if function == "readAssetSamples" {
		// returns selected sample objects 
		return t.readAssetSamples(stub, args)
	} else if function == "readAssetSchemas" {
		// returns selected sample objects 
		return t.readAssetSchemas(stub, args)
	}
    
    return nil, errors.New("Received unknown invocation: " + function)
}



func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("______________Inside Invoke");
	
	 if function == "createAsset" {
        // create assetID
        return t.createAsset(stub, args)
    } else if function == "updateAsset" {
        // update assetID
        return t.updateAsset(stub, args)
    } 
    
	return nil, errors.New("Received unknown invocation: " + function)
}


func (t *SimpleChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreateLoanApplication")
	fmt.Printf("______________Inside createClaimApplication");

	if len(args) < 1 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for Claim application creation")
	}
		
		var payload = args[0]
		b := []byte(payload)
		
		var c Claim
		var err = json.Unmarshal(b, &c)
		
		var bytes []byte
		
		//DMV
		/*
		body := bytes.NewBuffer(b)
		r_dmv, _ := http.Post("https://claimnode.mybluemix.net/verify/DMV", "application/json", body)
		response_dmv, _ := ioutil.ReadAll(r_dmv.Body)
		
		var strDMVResponse  = string(response_dmv)
		strDMVResponse = strings.Replace(strDMVResponse, "\\", "" , -1)
		
		// ISO
		
		r_iso, _ := http.Post("https://claimnode.mybluemix.net/verify/ISO", "application/json", body)
		response_iso, _ := ioutil.ReadAll(r_iso.Body)
		
		var strISOResponse  = string(response_iso)
		strISOResponse = strings.Replace(strISOResponse, "\\", "" , -1)
		
		// Choicepoint
		
		r_choicepoint, _ := http.Post("https://claimnode.mybluemix.net/verify/Choicepoint", "application/json", body)
		response_choicepoint, _ := ioutil.ReadAll(r_choicepoint.Body)
		
		var strChoiceResponse  = string(response_choicepoint)
		strChoiceResponse = strings.Replace(strChoiceResponse, "\\", "" , -1)
		
	
		c.ExternalReport   = 	strDMVResponse + " , " +strISOResponse+ "  , " + strChoiceResponse
		*/
		
		flag , _  :=  t.checkFraudRecord(stub,c)
		if (!flag) {
			CLAIM_NO =  CLAIM_NO + 1
			c.ClaimNo = strconv.Itoa(CLAIM_NO)
			
			// Adjuster Assigment 

			c.AdjusterReport.AdjusterZipCode 	= c.LossDetails.LossZipCode
			c.AdjusterReport.AdjusterSpeciality = "Collision"
			c.AdjusterReport.AdjusterFirstName 	= "John"
			c.AdjusterReport.AdjusterLastName 	= "Doe"
			c.AdjusterReport.EvaluationDateTime = "03/02/2017"
		//	c.AdjusterReport.ApproveLossAmount	= "3000.00"


			c.RepairedDetails.RepairShopName	= "Quick Repair Shop"
			c.RepairedDetails.RepairZipCode  	= c.LossDetails.LossZipCode
			
			c.RepairedDetails.ItemRepaired[0].ItemId = "0"
			c.RepairedDetails.ItemRepaired[1].ItemId = "1"
			c.RepairedDetails.ItemRepaired[2].ItemId = "2"
			
			_ , err = save_changes(stub , c)
			
			if err != nil {
				logger.Error("Could not save claim to ledger", err)
				return nil, err
			}
			
			bytes, err = stub.GetState(c.ClaimNo)
			
			err = json.Unmarshal(bytes, &c); 
			
			var customEvent = "{\"ClaimNo\":\"" + c.ClaimNo +"\" ,  \"InsuredName\" :\""+c.InsuredDetails.FirstName+" "+c.InsuredDetails.LastName+"\" , \"Desc\":\"Claim  Created Successfully\"}"
			err = stub.SetEvent("Claim_Notification", []byte(customEvent))
			if err != nil {
				return nil, err
			}
			
		} else {
			returnMsg := "{"+"\"ReturnMessage\""+":"+"\"Poteltial Duplicate Claim\""+"}" 
			bytes = (([]byte)(returnMsg))
		
			var customEvent = "{\"PolicyNo\":\"" + c.PolicyNo +"\" ,  \"InsuredName\" :\""+c.InsuredDetails.FirstName+" "+c.InsuredDetails.LastName+"\" , \"Desc\":\"Potential Fraud Claim with SSN :"+c.InsuredDetails.SSN+" VIN :"+c.VehicleDetails.VIN +" and Loss Date:"+c.LossDetails.LossDateTime+"\" }"
			err = stub.SetEvent("Claim_Notification", []byte(customEvent))
			if err != nil {
				return nil, err
			}
			
			
			return nil, errors.New(returnMsg)
		}
		
		




	
	logger.Info("______Returning from create claim application :"+c.ClaimNo)
	return bytes, nil

}


//******************** updateAsset ********************/

func (t *SimpleChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdateLoanApplication")

	
	
	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for claim application update")
		}

	
	var asset					= args[0]
	var claimNo 				= args[1]
	
	
	if asset == "InvestigationReport"  {
	
/*		
		var evaluationDateTime		= args[2]
		var lossAmount 				= args[3]
		var remarks					= args[4]
		var status					= args[5]
*/		
		var payload = args[2]

		b := []byte(payload)
		var a Adjuster
		var err = json.Unmarshal(b, &a)
	
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim
		err = json.Unmarshal(laBytes, &claimApplication)
		
		claimApplication.AdjusterReport.EvaluationDateTime 	= a.EvaluationDateTime
		claimApplication.AdjusterReport.ApproveLossAmount 	= a.ApproveLossAmount
		claimApplication.AdjusterReport.Remarks 			= a.Remarks
		claimApplication.Status								= "Claim_Approved"
		
		laBytes, err = json.Marshal(&claimApplication)
		
		if err != nil {
		logger.Error("Could not marshal claim application post update", err)
		return nil, err
		}
		
		err = stub.PutState(claimNo, laBytes)
		if err != nil {
			logger.Error("Could not save claim application post update", err)
			return nil, err
		}
		
		var customEvent = "{\"ClaimNo\":\"" + claimApplication.ClaimNo +"\" ,  \"InsuredName\" :\""+claimApplication.InsuredDetails.FirstName+" "+claimApplication.InsuredDetails.LastName+"\" , \"Desc\":\"Investigation Report Submitted Successfully\"}"
		err = stub.SetEvent("Claim_Investigation_Report", []byte(customEvent))
		if err != nil {
			return nil, err
		}
		
	} else if asset == "RequestApproval"  {
	
		
		var payload = args[2]

		b := []byte(payload)
		var r RepairShop
		var err = json.Unmarshal(b, &r)
	
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim 
		err = json.Unmarshal(laBytes, &claimApplication)
		
		for index, each := range r.ItemRepaired {
		
			claimApplication.RepairedDetails.ItemRepaired[index].ItemId 		= r.ItemRepaired[index].ItemId
			claimApplication.RepairedDetails.ItemRepaired[index].ItemName 		= r.ItemRepaired[index].ItemName
			claimApplication.RepairedDetails.ItemRepaired[index].ItemCost 		= r.ItemRepaired[index].ItemCost
	
			logger.Debug("Divine value each item >>", r.ItemRepaired[index].ItemId)
			
			
			logger.Debug("Divine value [%d] is [%s]\n", index, each)
		}
		
		
		claimApplication.RepairedDetails.RepairDateTime 			= r.RepairDateTime
		claimApplication.RepairedDetails.TotalCost 					= r.TotalCost
		claimApplication.Status										= "Request_Approval"
		
		laBytes, err = json.Marshal(&claimApplication)
		
		if err != nil {
		logger.Error("Could not marshal claim application post update", err)
		return nil, err
		}

		err = stub.PutState(claimNo, laBytes)
		if err != nil {
			logger.Error("Could not save claim application post update", err)
			return nil, err
		}
		
		var customEvent = "{\"ClaimNo\":\"" + claimApplication.ClaimNo +"\" ,  \"InsuredName\" :\""+claimApplication.InsuredDetails.FirstName+" "+claimApplication.InsuredDetails.LastName+"\" , \"Desc\":\"Repair Approval Request Submitted Successfully\"}"
		err = stub.SetEvent("Claim_Request_Approval", []byte(customEvent))
		if err != nil {
			return nil, err
		}
		
	} else if asset == "RepairInvoice"  {
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim 
		err = json.Unmarshal(laBytes, &claimApplication)
	

		claimApplication.Status			= "Repair_Completed"
		
		laBytes, err = json.Marshal(&claimApplication)
		
		if err != nil {
		logger.Error("Could not marshal claim application post update", err)
		return nil, err
		}

		err = stub.PutState(claimNo, laBytes)
		if err != nil {
			logger.Error("Could not save claim application post update", err)
			return nil, err
		}
		
		var customEvent = "{\"ClaimNo\":\"" + claimApplication.ClaimNo +"\" ,  \"InsuredName\" :\""+claimApplication.InsuredDetails.FirstName+" "+claimApplication.InsuredDetails.LastName+"\" , \"Desc\":\"Repair Invoice Submitted Successfully\"}"
		err = stub.SetEvent("Claim_Repair_Invoice", []byte(customEvent))
		if err != nil {
			return nil, err
		}
		
	} else if asset == "ApproveRepairClaim"  {
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim 
		err = json.Unmarshal(laBytes, &claimApplication)
	

		claimApplication.Status	 = "Approve_Claim"
		
		laBytes, err = json.Marshal(&claimApplication)
		
		if err != nil {
		logger.Error("Could not marshal claim application post update", err)
		return nil, err
		}

		err = stub.PutState(claimNo, laBytes)
		if err != nil {
			logger.Error("Could not save claim application post update", err)
			return nil, err
		}
		
		var customEvent = "{\"ClaimNo\":\"" + claimApplication.ClaimNo +"\" ,  \"InsuredName\" :\""+claimApplication.InsuredDetails.FirstName+" "+claimApplication.InsuredDetails.LastName+"\" , \"Desc\":\"Repair Invoice Approved Successfully\"}"
		err = stub.SetEvent("Claim_Repair_Approval", []byte(customEvent))
		if err != nil {
			return nil, err
		}
	} else if asset == "Payment"  {
	
		var payload = args[2]

		b := []byte(payload)
		var p Payment
		var err = json.Unmarshal(b, &p)	
	
		laBytes, err := stub.GetState(claimNo)
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim
		err = json.Unmarshal(laBytes, &claimApplication)
		
		claimApplication.PaymentDetails.BankName 			= p.BankName
		claimApplication.PaymentDetails.AccountNo 			= p.AccountNo
		claimApplication.PaymentDetails.PaymentAmount 		= p.PaymentAmount
		claimApplication.PaymentDetails.PaymentDateTime 	= p.PaymentDateTime
		claimApplication.Status								= "Payment_Submitted"
		
		laBytes, err = json.Marshal(&claimApplication)
		
		if err != nil {
		logger.Error("Could not marshal claim application post update", err)
		return nil, err
		}

		err = stub.PutState(claimNo, laBytes)
		if err != nil {
			logger.Error("Could not save claim application post update", err)
			return nil, err
		}
		
		var customEvent = "{\"ClaimNo\":\"" + claimApplication.ClaimNo +"\" ,  \"InsuredName\" :\""+claimApplication.InsuredDetails.FirstName+" "+claimApplication.InsuredDetails.LastName+"\" , \"Desc\":\"Payment for Repair Submitted Successfully\"}"
		err = stub.SetEvent("Claim_Bank_Payment", []byte(customEvent))
		if err != nil {
			return nil, err
		}
	}


	logger.Info("Successfully updated claim application")
	return nil, nil
}


func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	logger.Debug("Entering GetCertAttribute")
	attr, err := stub.ReadCertAttribute(attributeName)
	if err != nil {
		return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
	}
	attrString := string(attr)
	return attrString, nil
}






func save_changes(stub shim.ChaincodeStubInterface, c Claim) (bool, error) {

	bytes, err := json.Marshal(c)

	if err != nil { logger.Error("SAVE_CHANGES: Error converting vehicle record: ", err); return false, errors.New("Error converting claim record") }

	key := c.InsuredDetails.SSN + c.VehicleDetails.VIN + c.LossDetails.LossDateTime
	
	thKey := c.ThirdPartyInsuredDetails.SSN + c.ThirdPartyVehicleDetails.VIN + c.LossDetails.LossDateTime
	
	logger.Debug("____________Save_changes for the key :- "+key)
	
	err = stub.PutState(key, bytes)

	if err != nil { logger.Error("SAVE_CHANGES: Error storing Claim record:", err); return false, errors.New("Error storing claim record") }
	
	logger.Debug("____________Save_changes for the key :- "+key)
	
	err = stub.PutState(thKey, bytes)

	if err != nil { logger.Error("SAVE_CHANGES: Error storing Claim record:", err); return false, errors.New("Error storing claim record") }
	
	logger.Debug("____________Save_changes for the Thirtd party Key :- "+thKey)

	err = stub.PutState(c.ClaimNo, bytes)
	
	if err != nil { logger.Error("SAVE_CHANGES: Error storing Claim record: ", err); return false, errors.New("Error storing claim record") }
	logger.Debug("Save Complete for the key :- "+c.ClaimNo)
	
	logger.Debug("____________Save_changes for the Claim :- "+c.ClaimNo)
	
	return true, nil
}

func (t *SimpleChaincode) checkFraudRecord(stub shim.ChaincodeStubInterface , c Claim ) (bool , error) {
	logger.Debug("Entering checkFraudRecord for the key :- "+c.InsuredDetails.SSN + c.VehicleDetails.VIN + c.LossDetails.LossDateTime)

	key := c.InsuredDetails.SSN + c.VehicleDetails.VIN + c.LossDetails.LossDateTime
	
	logger.Debug("checkFraudRecord with Key "+key+" from ledger")
	
	thkey :=  c.ThirdPartyInsuredDetails.SSN + c.ThirdPartyVehicleDetails.VIN + c.LossDetails.LossDateTime
	
	logger.Debug("checkFraudRecord with Key "+thkey+" from ledger")
	
	var dupClaim Claim  
	bytes, err := stub.GetState(key)
	
	if (bytes != nil ) {
	    err = json.Unmarshal(bytes, &dupClaim); 
	    
	    if err != nil {
		logger.Error("Could not fetch Claim application with Key "+key+" from ledger", err)
		
		}
	}
	
	if (c.InsuredDetails.SSN == dupClaim.ThirdPartyInsuredDetails.SSN && c.VehicleDetails.VIN == dupClaim.ThirdPartyVehicleDetails.VIN &&  c.LossDetails.LossDateTime == dupClaim.LossDetails.LossDateTime ) {
		logger.Debug("Duplicate Claim Found with Key :-"+dupClaim.ThirdPartyInsuredDetails.SSN + dupClaim.ThirdPartyVehicleDetails.VIN + dupClaim.LossDetails.LossDateTime)
		return true , nil 
	}

	
	if (c.InsuredDetails.SSN == dupClaim.InsuredDetails.SSN && c.VehicleDetails.VIN == dupClaim.VehicleDetails.VIN &&  c.LossDetails.LossDateTime == dupClaim.LossDetails.LossDateTime ) {
		logger.Debug("Duplicate Claim Found with Key :-"+dupClaim.InsuredDetails.SSN + dupClaim.VehicleDetails.VIN + dupClaim.LossDetails.LossDateTime)
		return true , nil 
	}

	if ( len(thkey) > 0) {
	
		bytes, err = stub.GetState(thkey)
		
		if (bytes != nil ) {
		    err = json.Unmarshal(bytes, &dupClaim); 
		}
	
	
		if (c.ThirdPartyInsuredDetails.SSN == dupClaim.InsuredDetails.SSN && c.ThirdPartyVehicleDetails.VIN == dupClaim.VehicleDetails.VIN &&  c.LossDetails.LossDateTime == dupClaim.LossDetails.LossDateTime ) {
			logger.Debug("Duplicate Claim Found with Key :-"+dupClaim.InsuredDetails.SSN + dupClaim.VehicleDetails.VIN + dupClaim.LossDetails.LossDateTime)
			return true , nil 
		}
	
	}
		
	return false, nil
	
}



func retrieve_Claim(stub shim.ChaincodeStubInterface, claimNo string) (Claim, error) {

	var c Claim

	bytes, err := stub.GetState(claimNo);

	if err != nil {	fmt.Printf("RETRIEVE_claimId: Failed to invoke vehicle_code: %s", err); return c, errors.New("RETRIEVE_V5C: Error retrieving vehicle with v5cID = " + claimNo) }

	err = json.Unmarshal(bytes, &c);

    if err != nil {	fmt.Printf("RETRIEVE_claimId: Corrupt vehicle record "+string(bytes)+": %s", err); return c, errors.New("RETRIEVE_V5C: Corrupt vehicle record"+string(bytes))	}

	return c, nil
}

func (t *SimpleChaincode) createFraudTable(stub shim.ChaincodeStubInterface) (error) {
	
	var column[] *shim.ColumnDefinition
	
	column[0].Name= "SSN"
	column[0].Type= shim.ColumnDefinition_STRING
	
	column[1].Name= "VIN"
	column[1].Type= shim.ColumnDefinition_STRING


	column[2].Name= "DOL"
	column[2].Type= shim.ColumnDefinition_STRING
 
	err:= stub.CreateTable("CHECK_FRAUD_TABLE" , column);
	return err
}

func (t *SimpleChaincode) insertFraudTable(stub shim.ChaincodeStubInterface , args []string) ([]byte, error) {

	var row shim.Row
	//var column *shim.Column
	var temp , temp1 , temp2 *shim.Column_String_
	temp.String_= args[0] //"12345"
	temp1.String_=args[1] //"999999"
	temp2.String_=args[2] // "03/01/1981"
	//column = &shim.Column{Value :"12345"}
	row.Columns[0].Value=temp
	row.Columns[1].Value=temp1
	row.Columns[2].Value=temp2
//	row.VIN ="XYZ" 
//	row.DOL ="03/01/2017" 

	stub.InsertRow("CHECK_FRAUD_TABLE" ,row)
	
	return nil , nil
}


/*	var key []shim.Column 
	var row shim.Row
//	var err error
	var temp , temp1 , temp2 *shim.Column_String_
	
	temp.String_= args[0] //"12345"
	temp1.String_= args[1] //"999999"
	temp2.String_= args[2] //"03/01/1981"
	
	key[0].Value = temp
	key[1].Value = temp1
	key[2].Value = temp2
	
	row , _= stub.GetRow("CHECK_FRAUD_TABLE" ,key )
	
	if (len(row.Columns) > 0) {
		return "TRUE" , nil
	}
	
	return "FALSE" , nil
	*/
