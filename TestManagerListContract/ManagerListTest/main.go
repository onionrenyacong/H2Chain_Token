package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/AElfProject/aelf-sdk.go/client"
	pb "github.com/AElfProject/aelf-sdk.go/protobuf/generated"
	"github.com/AElfProject/aelf-sdk.go/utils"
	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var aelfClient client.AElfClient
var fromAddress string

func aelfInit(host string, pk string) {
	aelfClient = client.AElfClient{
		Host:       host,
		Version:    "1.0",
		PrivateKey: pk,
	}
	fromAddress = aelfClient.GetAddressFromPrivateKey(aelfClient.PrivateKey)
}

func Initialize() {
	managerListAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.ManagerList")
	empty_buf, _ := proto.Marshal(&empty.Empty{})
	tx, _ := aelfClient.CreateTransaction(fromAddress, managerListAddr, "Initialize", empty_buf)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)
	sendResult, _ := aelfClient.SendTransaction(hex.EncodeToString(transactionByets))
	fmt.Println("txID = " + sendResult.TransactionID)
	for {
		res, _ := aelfClient.GetTransactionResult(sendResult.TransactionID)
		if res.Status == "MINED" {
			fmt.Println("Mined.")
			break
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func AddManager(addr string) {
	addr_sv := wrappers.StringValue{
		Value: addr,
	}
	addr_buf, _ := proto.Marshal(&addr_sv)
	managerListAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.ManagerList")
	tx, _ := aelfClient.CreateTransaction(fromAddress, managerListAddr, "AddManager", addr_buf)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)
	sendResult, _ := aelfClient.SendTransaction(hex.EncodeToString(transactionByets))
	fmt.Println("txID = " + sendResult.TransactionID)
	for {
		res, _ := aelfClient.GetTransactionResult(sendResult.TransactionID)
		if res.Status == "MINED" {
			fmt.Println("Mined.")
			break
		}
		fmt.Println(res.Status)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func RemoveManager(addr string) {
	addr_sv := wrappers.StringValue{
		Value: addr,
	}
	addr_buf, _ := proto.Marshal(&addr_sv)
	managerListAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.ManagerList")
	tx, _ := aelfClient.CreateTransaction(fromAddress, managerListAddr, "RemoveManager", addr_buf)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)
	sendResult, _ := aelfClient.SendTransaction(hex.EncodeToString(transactionByets))
	fmt.Println("txID = " + sendResult.TransactionID)
	for {
		res, _ := aelfClient.GetTransactionResult(sendResult.TransactionID)
		if res.Status == "MINED" {
			fmt.Println("Mined.")
			break
		}
		fmt.Println(res.Status)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func CheckManager(addr string) {
	addr_sv := wrappers.StringValue{
		Value: addr,
	}
	addr_buf, _ := proto.Marshal(&addr_sv)
	managerListAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.ManagerList")
	tx, _ := aelfClient.CreateTransaction(fromAddress, managerListAddr, "CheckManager", addr_buf)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)

	//注意：View
	sendResult, _ := aelfClient.ExecuteTransaction(hex.EncodeToString(transactionByets))
	res, _ := hex.DecodeString(sendResult)
	realRes := &wrapperspb.BoolValue{}
	proto.Unmarshal(res, realRes)
	fmt.Println(realRes.Value)
}

func transfer(to string, amount int64) {
	toAddr, _ := utils.Base58StringToAddress(to)
	transferInput := &pb.TransferInput{
		To:     toAddr,
		Symbol: "ELF",
		Amount: amount,
		Memo:   "test",
	}
	transferInputProto, _ := proto.Marshal(transferInput)
	multiTokenAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.Token")
	tx, _ := aelfClient.CreateTransaction(fromAddress, multiTokenAddr, "Transfer", transferInputProto)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)
	sendResult, _ := aelfClient.SendTransaction(hex.EncodeToString(transactionByets))
	fmt.Println("txID = " + sendResult.TransactionID)
	for {
		res, _ := aelfClient.GetTransactionResult(sendResult.TransactionID)
		if res.Status == "MINED" {
			fmt.Println("Mined.")
			break
		}
		fmt.Println(res.Status)
		time.Sleep(time.Duration(1) * time.Second)
	}

}

func balance(who string) {
	ownerAddr, _ := utils.Base58StringToAddress(who)
	balanceInput := &pb.GetBalanceInput{
		Owner:  ownerAddr,
		Symbol: "ELF",
	}
	balanceInputProto, _ := proto.Marshal(balanceInput)
	multiTokenAddr, _ := aelfClient.GetContractAddressByName("AElf.ContractNames.Token")
	tx, _ := aelfClient.CreateTransaction(fromAddress, multiTokenAddr, "GetBalance", balanceInputProto)
	signature, _ := aelfClient.SignTransaction(aelfClient.PrivateKey, tx)
	tx.Signature = signature
	transactionByets, _ := proto.Marshal(tx)
	sendResult, _ := aelfClient.ExecuteTransaction(hex.EncodeToString(transactionByets))
	res, _ := hex.DecodeString(sendResult)
	realRes := &pb.GetBalanceOutput{}
	proto.Unmarshal(res, realRes)
	fmt.Printf("Balance of %s = %d\n", who, realRes.Balance)
}

func main() {
	//第一个参数ip+端口，第二个参数私钥
	aelfInit("http://101.201.46.135:8000", "8e9d0c5741c72690cb0031894cd91bb7278395907d6631a7aa5b86b8beb75585")

	Initialize()
	//下面的都是Address
	transfer("zuPmz56jyfVcduKk3qgZbUXxY7ED3kggQTYj48tFrS2xpmwnY", 10)

	balance("2En93BYxtBWSghUZWnKudaTX9cdG59SRs4XTyqR5ALGj2REhM8")

	RemoveManager("2En93BYxtBWSghUZWnKudaTX9cdG59SRs4XTyqR5ALGj2REhM8")

	CheckManager("2eFQAZk5bYmeBVg1RN7qSAydS1AYyXrSyFVU9BfzBxFZh4csuf")

}
