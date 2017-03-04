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
	//"strconv"
	//"strings"
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
//	LicensePlate	string `json:"licensePlate"`
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
	DrivingLicense         string `json:"DrivingLicense"`
}

type Property struct {
	PropertyAddress            	string `json:"PropertyAddress"`
	PropertyCity         		string `json:"PropertyCity"`
	PropertyState           	string `json:"PropertyState"`
	PropertyZip             	string `json:"PropertyZip"`
	IfRoofDamaged       		string `json:"ifRoofDamaged"`
	IfLightingCausedFire       	string `json:"ifLightingCausedFire"`
}


type Claim struct {
	 
	ClaimId	    		string		`json:"claimId"` 
	PolicyNo			string		`json:"policyNo"` 
	ClaimNo	    		string		`json:"claimNo"`
	EstmLossAmount		string		`json:"estmLossAmount"` 
	Status              string      `json:"status"`
	LossDetails 		Loss 		`json:"lossDetails"`
	InsuredDetails 		Insured 	`json:"insuredDetails"`
	VehicleDetails 		Vehicle 	`json:"vehicleDetails"`
//	PropertyDetails 	Property 	`json:"propertyDetails"`
}





// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("Initialization Complete ")
	logger.Info("Initialization Complete ")
	
	return nil, nil
}

func getClaimApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetLoanApplication")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing Claim No")
	}

	var claimNo = args[0]
	bytes, err := stub.GetState(claimNo)
	if err != nil {
		logger.Error("Could not fetch Claim application with No "+claimNo+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func createClaimApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreateLoanApplication")
	fmt.Printf("______________Inside createClaimApplication");

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for Claim application creation")
	}

	var err error
	var claimObj Claim 
	
	var claimNo = args[0]
	var claimApplicationInput = args[1]
	
	logger.Debug("Entering CreateLoanApplication " +claimNo)
	fmt.Printf("______________Inside createClaimApplication" + claimApplicationInput);

	b := []byte(claimApplicationInput)
	err = json.Unmarshal(b, &claimObj)
	
	 _ , err = save_changes(stub,claimObj)
	
	//err := stub.PutState(claimNo, bytes)
	if err != nil {
		logger.Error("Could not save claim  to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'claimApplicationCreation', description:" + claimNo + "' Successfully created'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved claim application")
	return nil, nil

}


func updateClaimApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdateLoanApplication")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for claim application update")
	}

	var claimNo = args[0]
	var status = args[1]

	laBytes, err := stub.GetState(claimNo)
	if err != nil {
		logger.Error("Could not fetch claim application from ledger", err)
		return nil, err
	}
	var claimApplication Claim
	err = json.Unmarshal(laBytes, &claimApplication)
	claimApplication.Status = status

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

	var customEvent = "{eventType: 'claimApplicationUpdate', description:" + claimNo + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully updated claim application")
	return nil, nil

}



// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("______________Inside Invoke");
	
	if function == "createClaimApplication" {
		fmt.Printf("______________Calling createClaimApplication");
		return createClaimApplication(stub, args)
	}
	/*if function == "createClaimApplication" {
		username, _ := GetCertAttribute(stub, "username")
		role, _ := GetCertAttribute(stub, "role")
		if role == "Claim_CSR" {
			
		} else {
			return nil, errors.New(username + " with role " + role + " does not have access to create a claim application")
		}
	}else if function == "updateClaimApplication" {
			username, _ := GetCertAttribute(stub, "username")
			role, _ := GetCertAttribute(stub, "role")
			if role == "Claim_UPDATE" {
				return updateClaimApplication(stub, args)
			} else {
				return nil, errors.New(username + " with role " + role + " does not have access to create a claim application")
			}
	}
	*/
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "getClaimApplication" {
		return getClaimApplication(stub, args)
	}

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

func  add_fnol(stub shim.ChaincodeStubInterface, claimObj Claim) ([]byte, error) {
   
    var err error
    fmt.Println("running add_fnol()")

    _ ,err = save_changes(stub, claimObj)
     
    if err != nil {
        return nil, err
    }
    return nil, nil
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