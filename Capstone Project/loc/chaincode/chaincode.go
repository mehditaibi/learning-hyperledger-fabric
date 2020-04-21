package main 

import (
	"fmt"
	"github.com/hyperledger/fabric/core/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type LocAsset struct { 

}

func (t *LocAsset) Init(stub shim.ChaincodeStubInterface) peer.Response { 
	// Get the args from the transaction proposal
	args := stub.GetStringArgs()
	if len(args) !=2 { 
		return shim.Error("Incorrect Arguments. We are expecting a key and a value pair")
	}

	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil { 
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)
}

func (t *LocAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Reponse { 

	// Extract the function and args from the transaction proposal 

	fn, args := stub.GetFunctionAndParameters()

	var result string 

	var err error 

	if fn == "set" { 
		result, err = set(stub, args)
	} else { 
		result, err = get(stub, args)
	}

	if err != nil { 
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) { 
	if len(args) !=2 { 
		return "", fmt.Error("Incorrect Arguments. We are expecting a key and a value pair")
	}

	err := stub.PutState(args[0], [byte(args[1])])

	if err != nil {
		return "", fmt.Errorf("Failed to set the asset: %s", args[0])
	}

	return args[1], nil 
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) { 
	if len(args) != 1{ 
		return "", fmt.Error("Incorrect Arguments. We are expecting a key and a value pair")
	}

	value, err := stub.GetState(args[0])

	if err != nil {
		return "", fmt.Errorf("Failed to get the asset: %s", args[0])
	}

	if value == nil {
		return "", fmt.Errorf("Asset not found on BlockChain: %s", args[0])
	}

	return string(value), nil 
}

func main(){ 
	if err := shim.Start(new(LocAsset)); err != nil { 
		fmt.Printf("Error Starting LocAsset Chaincode: %s", err)
	}
}