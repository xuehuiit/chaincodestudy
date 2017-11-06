/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */

package main

import (
	//"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte){

}


func TestExample02_Init(t *testing.T) {

	scc := new(simplechaincode)
	stub := shim.NewMockStub("ex02", scc)


	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})


}