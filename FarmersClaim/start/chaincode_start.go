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
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	//"strconv" 
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"regexp"
)


var logger = shim.NewLogger("ClaimChaincode")


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
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

}

type Insured struct {
	FirstName              string `json:"firstName,omitempty"`
	LastName           	   string `json:"lastName,omitempty"`
	PhoneNo         	   string `json:"phoneNo,omitempty"`
	Email           	   string `json:"email,omitempty"`
	Dob             	   string `json:"dob,omitempty"`
	Ssn             	   string `json:"ssn,omitempty"`
	DrivingLicense         string `json:"drivingLicense,omitempty"`
}


type Adjuster struct {
	 
	EvaluationDateTime	string		`json:"evaluationDateTime,omitempty"` 
	LossAmount			string		`json:"lossAmount,omitempty"` 
	Remarks	    		string		`json:"remarks,omitempty"`

}

type Repair struct {
	 
	RepairDateTime			string		`json:"repairDateTime,omitempty"` 
	ItemRepaired			string		`json:"itemRepaired,omitempty"` 
	Cost	    			string		`json:"cost,omitempty"`

}

type Payment struct {
	 
	AccountNo				string		`json:"accountNo,omitempty"` 
	PaymentAmount			string		`json:"paymentAmount,omitempty"` 
	PaymentDateTime	    	string		`json:"paymentDateTime,omitempty"`

}

type Claim struct {
	 
	ClaimId	    		string		`json:"claimId,omitempty"` 
	PolicyNo			string		`json:"policyNo,omitempty"` 
	ClaimNo	    		string		`json:"claimNo,omitempty"`
	EstmLossAmount		string		`json:"estmLossAmount,omitempty"` 
	Status              string      `json:"status,omitempty"`
//	ExternalReport      string      `json:"externalReport,omitempty"`
//	LossDetails 		Loss 		`json:"lossDetails,omitempty"`
//	InsuredDetails 		Insured 	`json:"insuredDetails,omitempty"`
//	VehicleDetails 		Vehicle 	`json:"vehicleDetails,omitempty"`
//	AdjusterReport 		Adjuster 	`json:"adjusterReport,omitempty"`
//	RepairedDetails 	Repair 		`json:"repairedDetails,omitempty"`
//	PaymentDetails 		Payment 	`json:"paymentDetails,omitempty"`
//	SensorData 		    Sensor 		`json:"sensorData,omitempty"`

}



type Sensor struct {
    Latitude    *string `json:"latitude,omitempty"`
    Longitude   *string `json:"longitude,omitempty"`
//    Image   	*string `json:"image,omitempty"`
//    Voice   	*string `json:"voice,omitempty"`
}



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
	logger.Debug("Initialization Complete ")
	
	return nil, nil
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

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for Claim application creation")
	}
		var claimNo = args[0]
		var payload = args[1]
		
		payload = strings.Replace(payload, "^", "\"" , -1)
		b := []byte(payload)
		
		var c Claim
		var err = json.Unmarshal(b, &c)
		
		//DMV
		
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
		_ , err = save_changes(stub , c)
		
		bytes, err := stub.GetState(claimNo)
	
		err = json.Unmarshal(bytes, &c); 
	
	//err := stub.PutState(claimNo, bytes)
	if err != nil {
		logger.Error("Could not save claim  to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'claimApplicationCreation', description:" + claimNo + "' Successfully created'}"
	err = stub.SetEvent("Claim_Verification", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved claim application")
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
	
		
		var evaluationDateTime		= args[2]
		var lossAmount 				= args[3]
		var remarks					= args[4]
		var status					= args[5]
		
	
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim
		err = json.Unmarshal(laBytes, &claimApplication)
		
		claimApplication.AdjusterReport.EvaluationDateTime 	= evaluationDateTime
		claimApplication.AdjusterReport.LossAmount 			= lossAmount
		claimApplication.AdjusterReport.Remarks 			= remarks
		claimApplication.Status								= status
		
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
		
		var customEvent = "{eventType: 'claimApplicationUpdate', description:" + claimNo + "' : Investigation Report Submitted'}"
		err = stub.SetEvent("Investigation_Report", []byte(customEvent))
		if err != nil {
			return nil, err
		}
		
	} else if asset == "RepairInvoice"  {
	
		
		var repairDateTime 		= args[2]
		var itemRepaired 		= args[3]
		var cost 				= args[4]
		var status 				= args[5]
		
		laBytes, err := stub.GetState(claimNo)
		
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim
		err = json.Unmarshal(laBytes, &claimApplication)
		
		claimApplication.RepairedDetails.RepairDateTime 	= repairDateTime
		claimApplication.RepairedDetails.ItemRepaired 		= itemRepaired
		claimApplication.RepairedDetails.Cost 				= cost
		claimApplication.Status								= status
		
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
		
		var customEvent = "{eventType: 'claimApplicationUpdate', description:" + claimNo + "' : Repair Invoice Submitted'}"
		err = stub.SetEvent("Repair_Invoice", []byte(customEvent))
		if err != nil {
			return nil, err
		}
		
	} else if asset == "Payment"  {
	
		
		var accountNo 		= args[2]
		var paymentAmount 	= args[3]
		var paymentDateTime = args[4]
		var status 			= args[5]
			
	
		laBytes, err := stub.GetState(claimNo)
		if err != nil {
			logger.Error("Could not fetch claim application from ledger", err)
			return nil, err
		}
		var claimApplication Claim
		err = json.Unmarshal(laBytes, &claimApplication)
		
		claimApplication.PaymentDetails.AccountNo 			= accountNo
		claimApplication.PaymentDetails.PaymentAmount 		= paymentAmount
		claimApplication.PaymentDetails.PaymentDateTime 	= paymentDateTime
		claimApplication.Status								= status
		
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
		
		var customEvent = "{eventType: 'claimApplicationUpdate', description:" + claimNo + "' : Payment Processed'}"
		err = stub.SetEvent("Bank_Payment", []byte(customEvent))
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



type customEvent struct {
	Type       string `json:"type"`
	Decription string `json:"description"`
}



func save_changes(stub shim.ChaincodeStubInterface, c Claim) (bool, error) {

	bytes, err := json.Marshal(c)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error converting vehicle record: %s", err); return false, errors.New("Error converting claim record") }

	err = stub.PutState(c.ClaimNo, bytes)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error storing vehicle record: %s", err); return false, errors.New("Error storing claim record") }

	return true, nil
}

func retrieve_Claim(stub shim.ChaincodeStubInterface, claimNo string) (Claim, error) {

	var c Claim

	bytes, err := stub.GetState(claimNo);

	if err != nil {	fmt.Printf("RETRIEVE_claimId: Failed to invoke vehicle_code: %s", err); return c, errors.New("RETRIEVE_V5C: Error retrieving vehicle with v5cID = " + claimNo) }

	err = json.Unmarshal(bytes, &c);

    if err != nil {	fmt.Printf("RETRIEVE_claimId: Corrupt vehicle record "+string(bytes)+": %s", err); return c, errors.New("RETRIEVE_V5C: Corrupt vehicle record"+string(bytes))	}

	return c, nil
}