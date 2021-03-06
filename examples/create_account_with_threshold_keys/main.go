package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	operatorAccountID, err := hedera.AccountIDFromString(os.Getenv("OPERATOR_ID"))
	if err != nil {
		panic(err)
	}

	operatorPrivateKey, err := hedera.Ed25519PrivateKeyFromString(os.Getenv("OPERATOR_KEY"))
	if err != nil {
		panic(err)
	}

	keys := make([]hedera.Ed25519PrivateKey, 3)
	pubKeys := make([]hedera.PublicKey, 3)

	fmt.Println("threshold key example")
	fmt.Println("Keys: ")

	for i := range keys {
		newKey, err := hedera.GenerateEd25519PrivateKey()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Key %v:\n", i)
		fmt.Printf("private = %v\n", newKey)
		fmt.Printf("public = %v\n", newKey.PublicKey())

		keys[i] = newKey
		pubKeys[i] = newKey.PublicKey()
	}

	client := hedera.ClientForTestnet()

	// A threshold key with a threshold of 2 and length of 3 requires
	// at least 2 of the 3 keys to sign anything modifying the account
	thresholdKey := hedera.NewThresholdKey(2).
		AddAll(pubKeys)

	fmt.Printf("threshold key %v\n", thresholdKey)

	transaction, err := hedera.NewAccountCreateTransaction().
		SetKey(thresholdKey).
		SetInitialBalance(hedera.NewHbar(6)).
		SetTransactionID(hedera.NewTransactionID(operatorAccountID)).
		SetTransactionMemo("sdk example create_account_with_threshold_keys/main.go").
		Build(client)

	if err != nil {
		panic(err)
	}

	transactionID, err := transaction.Sign(operatorPrivateKey).
		Execute(client)

	if err != nil {
		panic(err)
	}

	transactionReceipt, err := transactionID.GetReceipt(client)

	if err != nil {
		panic(err)
	}

	newAccountID := transactionReceipt.GetAccountID()

	fmt.Printf("account = %v\n", newAccountID)

	transferTx, err := hedera.NewCryptoTransferTransaction().
		SetTransactionID(hedera.NewTransactionID(newAccountID)).
		AddSender(newAccountID, hedera.HbarFrom(5, hedera.HbarUnits.Hbar)).
		AddRecipient(operatorAccountID, hedera.HbarFrom(5, hedera.HbarUnits.Hbar)).
		Build(client)

	if err != nil {
		panic(err)
	}

	// Manually sign with 2 of the private keys provided in the threshold
	transferID, err := transferTx.
		Sign(keys[0]).
		Sign(keys[1]).
		Execute(client)

	if err != nil {
		panic(err)
	}

	// Must wait for the transaction to go to consensus
	receipt, err := transferID.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status of transfer transaction: %v\n", receipt.Status)

	// Operator must be set
	client.SetOperator(operatorAccountID, operatorPrivateKey)

	balance, err := hedera.NewAccountBalanceQuery().
		SetAccountID(newAccountID).
		Execute(client)

	if err != nil {
		panic(err)
	}

	fmt.Printf("account balance after transfer: %v\n", balance)
}
