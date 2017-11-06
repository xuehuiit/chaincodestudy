/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 41132168111

*/

package main

import (

	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"


)

//定义一个机构体，作为chaincode的主对象，可以是任何符合go语言规范的命名方式
type simplechaincode struct {

}


/**
	系统初始化方法， 在部署chaincode的过程中当执行命令

    peer chaincode instantiate -o orderer.robertfabrictest.com:7050 -C
         roberttestchannel -n r_test_cc6 -v 1.0 -c '{"Args":["init","a","100","b","200"]}'
         -P "OR	('Org1MSP.member','Org2MSP.member')"

    的时候会调用该方法


*/
func (t *simplechaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println(" <<  ========  success init it is view in docker  ==========  >> ")
	return shim.Success([]byte("success init "))
}

/**

  主业务逻辑，在执行命令
  peer chaincode invoke -o 192.168.23.212:7050 -C roberttestchannel -n r_test_cc6 -c '{"Args":["invoke","a","b","1"]}'

  的时候系统会调用该方法并传入相关的参数，注意 "invoke" 之后的参数是需要传入的参数


*/
func (t *simplechaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("  ========  1、success it is view in docker  ========== ")
	fmt.Println("  ========  2、success it is view in docker  ========== ")
	fmt.Println("  ========  3、success it is view in docker  ========== ")
	fmt.Println("  ========  4、success it is view in docker  ========== ")
	return shim.Success([]byte("success invok "))

}


func main() {
	err := shim.Start(new(simplechaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
