package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. AssetID is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                           "policyNo": {
                                "description": "Policy Number",
                                "type": "string"
                            },
                             "status": {
                                "description": "Status",
                                "type": "string"
                            },
                                                      
                            "lossDetails": {
                                "description": "lossDetails",
                                "properties": {
                                    "lossType": {
                                        "type": "string"
                                    },
                                    "lossDateTime":{
                                        "type": "string"
                                    },
                                    "lossDescription":{
                                        "type": "string"
                                    },
                                    "lossAddress": {
                                        "type": "string"
                                    },
                                    "lossCity": {
                                        "type": "string"
                                    },
                                    "lossState": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            
							 "insuredDetails": {
	                            "description": "insuredDetails",
	                            "properties": {
	                                "firstName": {
	                                    "type": "string"
	                                },
	                                "lastName":{
	                                    "type": "string"
	                                },
	                                "phoneNo":{
	                                    "type": "string"
	                                },
	                                "email": {
	                                    "type": "string"
	                                },
	                                "dob": {
	                                    "type": "string"
	                                },
	                                "ssn": {
	                                    "type": "string"
	                                },
	                                "drivingLicense": {
	                                    "type": "string"
	                                }
	                            },
	                            "type": "object"
	                        },
	                         "thirdPartyInsuredDetails": {
	                            "description": "insuredDetails",
	                            "properties": {
	                                "firstName": {
	                                    "type": "string"
	                                },
	                                "lastName":{
	                                    "type": "string"
	                                },

	                                "ssn": {
	                                    "type": "string"
	                                },

	                            },
	                            "type": "object"
	                        },
	                         "vehicleDetails": {
	                            "description": "vehicleDetails",
	                            "properties": {

	                                "vin":{
	                                    "type": "string"
	                                }

	                             },
	                            "type": "object"
	                        },
                         "thirdPartyVehicleDetails": {
	                            "description": "vehicleDetails",
	                            "properties": {
	                                "make": {
	                                    "type": "string"
	                                },
	                                "Model":{
	                                    "type": "string"
	                                },
	                                "vin":{
	                                    "type": "string"
	                                },
	                                "year": {
	                                    "type": "string"
	                                }
	                             },
	                            "type": "object"
	                        },                          	                          
                        },
                   },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "SIMPLE",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. AssetID is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an assetID for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "assetID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "assetstatus": {
                            "description": "transport entity currently in possession of asset",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "string"
                                },
                                "longitude": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
							 "role": {
                                "description": "person role",
                                "type": "string"
                            },"lastowner": {
                                "description": "lastowner name",
                                "type": "string"
                            },
							 "ownername": {
                                "description": "ownername",
                                "type": "string"
                            },
							 "ownerid": {
                                "description": "ownerid",
                                "type": "string"
                            },
							 "overallstatus": {
                                "description": "overallstatus",
                                "type": "string"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetSamples": {
            "description": "Returns a string generated from the schema containing sample Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSamples function",
                    "enum": [
                        "readAssetSamples"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected sample data",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "readAssetSchemas": {
            "description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSchemas function",
                    "enum": [
                        "readAssetSchemas"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected schemas",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. AssetID is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "claimId": {
                                "description": "The ID of a managed Claim. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "policyNo": {
                                "description": "Policy Number",
                                "type": "string"
                            },
                            "claimNo": {
                                "description": "Claim Number",
                                "type": "string"
                            },
                            "estmLossAmount": {
                                "description": "EstmLossAmount",
                                "type": "string"
                            },
                            "status": {
                                "description": "Status",
                                "type": "string"
                            },
                            "externalReport": {
                                "description": "ExternalReport",
                                "type": "string"
                            },
                            
                            "lossDetails": {
                                "description": "lossDetails",
                                "properties": {
                                    "lossType": {
                                        "type": "string"
                                    },
                                    "lossDateTime":{
                                        "type": "string"
                                    },
                                    "lossDescription":{
                                        "type": "string"
                                    },
                                    "lossAddress": {
                                        "type": "string"
                                    },
                                    "lossCity": {
                                        "type": "string"
                                    },
                                    "lossState": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            
							 "insuredDetails": {
	                            "description": "insuredDetails",
	                            "properties": {
	                                "firstName": {
	                                    "type": "string"
	                                },
	                                "lastName":{
	                                    "type": "string"
	                                },
	                                "phoneNo":{
	                                    "type": "string"
	                                },
	                                "email": {
	                                    "type": "string"
	                                },
	                                "dob": {
	                                    "type": "string"
	                                },
	                                "ssn": {
	                                    "type": "string"
	                                },
	                                "drivingLicense": {
	                                    "type": "string"
	                                }
	                            },
	                            "type": "object"
	                        },
	                         "vehicleDetails": {
	                            "description": "vehicleDetails",
	                            "properties": {
	                                "make": {
	                                    "type": "string"
	                                },
	                                "Model":{
	                                    "type": "string"
	                                },
	                                "vin":{
	                                    "type": "string"
	                                },
	                                "year": {
	                                    "type": "string"
	                                }
	                             },
	                            "type": "object"
	                        },
	                        
	                      "adjusterReport": {
	                            "description": "adjusterReport",
	                            "properties": {
	                                "evaluationDateTime": {
	                                    "type": "string"
	                                },
	                                "lossAmount":{
	                                    "type": "string"
	                                },
	                                "remarks":{
	                                    "type": "string"
	                                }
	                             },
	                            "type": "object"
	                        },  
	                        
	                     "repairedDetails": {
	                            "description": "repairedDetails",
	                            "properties": {
	                                "repairDateTime": {
	                                    "type": "string"
	                                },
	                                "itemRepaired":{
	                                    "type": "string"
	                                },
	                                "cost":{
	                                    "type": "string"
	                                }
	                             },
	                            "type": "object"
	                        },  
	                         
	                     "paymentDetails": {
	                            "description": "repairedDetails",
	                            "properties": {
	                                "accountNo": {
	                                    "type": "string"
	                                },
	                                "paymentAmount":{
	                                    "type": "string"
	                                },
	                                "paymentDateTime":{
	                                    "type": "string"
	                                }
	                             },
	                            "type": "object"
	                        }, 
	                        
	                     "sensorData": {
	                            "description": "sensorData",
	                            "properties": {
	                                "latitude": {
	                                    "type": "string"
	                                },
	                                "longitude":{
	                                    "type": "string"
	                                },
	                                "image":{
	                                    "type": "string"
	                                },
									"voice":{
	                                    "type": "string"
	                                }	                                
	                             },
	                            "type": "object"
	                        } 	                          	                          
                        },
                        "required": [
                            "claimNo"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    
    "objectModelSchemas": {
        "assetIDKey": {
            "description": "An object containing only an assetID for use as an argument to read or delete.",
            "properties": {
                "claimId": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "claimId": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "assetstatus": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "string"
                        },
                        "longitude": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
							 "role": {
                                "description": "person role",
                                "type": "string"
                            },
							"lastowner": {
                                "description": "lastowner name",
                                "type": "string"
                            },
							 "ownername": {
                                "description": "ownername",
                                "type": "string"
                            },
							 "ownerid": {
                                "description": "ownerid",
                                "type": "string"
                            },
							 "overallstatus": {
                                "description": "overallstatus",
                                "type": "string"
                            }
            },
            "required": [
                "claimId"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "SIMPLE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete asset state.",
            "properties": {
                "claimId": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "assetstatus": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "string"
                        },
                        "longitude": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
							 "role": {
                                "description": "person role",
                                "type": "string"
                            },
							"lastowner": {
                                "description": "lastowner name",
                                "type": "string"
                            },
							 "ownername": {
                                "description": "ownername",
                                "type": "string"
                            },
							 "ownerid": {
                                "description": "ownerid",
                                "type": "string"
                            },
							 "overallstatus": {
                                "description": "overallstatus",
                                "type": "string"
                            }
            },
            "type": "object"
        }
    }
}`
