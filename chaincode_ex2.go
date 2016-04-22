package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("init is called")
	var Aval int
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	// Write the state to the ledger
	err = stub.PutState("abc", []byte(strconv.Itoa(Aval))) //making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}

	jsonAsBytes, _ := json.Marshal("shubham") //marshal an emtpy array of strings to clear the index
	err = stub.PutState("currentuser", jsonAsBytes)
	fmt.Println("initial current user: shubham")
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)

	// Handle different functions
	if function == "init" { //initialize the chaincode state, used as reset
		return t.init(stub, args)
	} else if function == "delete" { //deletes an entity from its state
		return t.Delete(stub, args)
	} else if function == "write" { //writes a value to the chaincode state
		return t.Write(stub, args)
	} else if function == "init_currentuser" { //create a new marble
		return t.init_currentuser(stub, args)
	} else if function == "get_currentuser" { //change owner of a marble
		return t.get_currentuser(stub)
	}
	fmt.Println("run did not find func: " + function) //error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Delete - remove a key/value pair from state
// ============================================================================================================================
func (t *SimpleChaincode) Delete(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("delete is called")
	return nil, nil
}

// ============================================================================================================================
// Query - read a variable from chaincode state - (aka read)
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	fmt.Println("query is called")
	return nil, nil //send it onward
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Write - write variable into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) Write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("write is called")
	return nil, nil
}

// ============================================================================================================================
// Init Marble - create a new marble, store into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) init_currentuser(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	fmt.Println("init current user is called")
	var err error

	user := strings.ToLower(args[0])
	fmt.Println("putting: "+user);
	err = stub.PutState("currentuser", []byte(user)) //store marble with id as key
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ============================================================================================================================
// Set User Permission on Marble
// ============================================================================================================================
func (t *SimpleChaincode) get_currentuser(stub *shim.ChaincodeStub) ([]byte, error) {
	var err error
	fmt.Println("get current user is called")
	currentuserbytes, err := stub.GetState("currentuser")
	if err != nil {
		return nil, errors.New("Failed to get thing")
	}
	var username string
	json.Unmarshal(currentuserbytes, &username) //un stringify it aka JSON.parse()
	fmt.Println("current user: "+username)
	return currentuserbytes,err;
}
