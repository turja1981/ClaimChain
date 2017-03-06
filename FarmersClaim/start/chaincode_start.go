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

const   CLAIM_SOURCE      	=  "claimsource"
const   INSURANCE_COMPANY 	=  "insuranceCompany"
const   ADJUSTER  			=  "adjuster"
const   BANK 				=  "bank"



// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


type Vehicle struct {
	Make            string `json:"make"`
	Model           string `json:"model"`
	VIN             string `json:"vin"`
	Year           	string `json:"year"`
}


type Loss struct {
	LossType            	string `json:"lossType"`
	LossDateTime            string `json:"lossDate"`
	LossDescription     	string `json:"lossDescription"`
	LossAddress         	string `json:"lossAddress"`
	LossCity            	string `json:"lossCity"`
	LossState	    		string `json:"lossState"`

}

type Insured struct {
	FirstName              string `json:"firstName"`
	LastName           	   string `json:"lastName"`
	PhoneNo         	   string `json:"phoneNo"`
	Email           	   string `json:"email"`
	Dob             	   string `json:"dob"`
	Ssn             	   string `json:"ssn"`
	DrivingLicense         string `json:"drivingLicense"`
}


type Adjuster struct {
	 
	EvaluationDateTime	string		`json:"evaluationDateTime"` 
	LossAmount			string		`json:"lossAmount"` 
	Remarks	    		string		`json:"remarks"`

}

type Repair struct {
	 
	RepairDateTime			string		`json:"repairDateTime"` 
	ItemRepaired			string		`json:"itemRepaired"` 
	Cost	    			string		`json:"cost"`

}

type Payment struct {
	 
	AccountNo				string		`json:"accountNo"` 
	PaymentAmount			string		`json:"paymentAmount"` 
	PaymentDateTime	    	string		`json:"paymentDateTime"`

}

type Claim struct {
	 
	ClaimId	    		string		`json:"claimId"` 
	PolicyNo			string		`json:"policyNo"` 
	ClaimNo	    		string		`json:"claimNo"`
	EstmLossAmount		string		`json:"estmLossAmount"` 
	Status              string      `json:"status"`
	ExternalReport      string      `json:"externalReport"`
	LossDetails 		Loss 		`json:"lossDetails"`
	InsuredDetails 		Insured 	`json:"insuredDetails"`
	VehicleDetails 		Vehicle 	`json:"vehicleDetails"`
	AdjusterReport 		Adjuster 	`json:"adjusterReport"`
	RepairedDetails 	Repair 		`json:"repairedDetails"`
	PaymentDetails 		Payment 	`json:"paymentDetails"`

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

func getClaimApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "getClaimApplication" {
		return getClaimApplication(stub, args)
	}

	return nil, nil
}

func createClaimApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
		
		var strClueResponse  = string(response_choicepoint)
		strClueResponse = strings.Replace(strClueResponse, "\\", "" , -1)
		
	
		c.ExternalReport   = 	strDMVResponse + " , " +strISOResponse+ "  , " + strClueResponse
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

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("______________Inside Invoke");
	
	if function == "createClaimApplication" {
		return createClaimApplication(stub , args)
	} else if function == "updateClaimApplication" {
		return updateClaimApplication(stub , function , args)
	}
		
	
	
	return nil, nil
}
func updateClaimApplication(stub shim.ChaincodeStubInterface, functionName string , args []string) ([]byte, error) {
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



// Invoke is our entry point to invoke a chaincode function




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