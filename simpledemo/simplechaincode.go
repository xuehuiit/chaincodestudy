/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321687

 */

package simpledemo


import (


	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"


)

type simplechaincode struct {

}


/**



 */
func (t *simplechaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {


	//获取传入参数值，在端采用数组的方式传入相关的参数
	_, args := stub.GetFunctionAndParameters()

	arg1 := args[0];
	arg2 := args[1];
	arg3 := args[2];

    //系统提示信息，该信息在客户端无法展现，只有进入响应的docker镜像才能正确展现
	fmt.Printf(" parm1: %s , parm2 : %s , parm3 : %s ", arg1,arg2,arg3)

	return shim.Success(nil)

}

/**



 */
func (t *simplechaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {



	//获取传入参数值，在端采用数组的方式传入相关的参数
	_, args := stub.GetFunctionAndParameters()

	arg1 := args[0];
	arg2 := args[1];
	arg3 := args[2];

	fmt.Printf(" parm1: %s , parm2 : %s , parm3 : %s ", arg1,arg2,arg3)

	return shim.Success(nil)


}
