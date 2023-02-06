package cryptosign

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/erebrus/api/v1/authenticate/flowid"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrFlowIdNotFound = errors.New("flow id not found")
)

func CheckSign(signature string, flowId string, message string) (string, bool, error) {
	// get flowid from the local db file
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)
	newMsgHash := crypto.Keccak256Hash([]byte(newMsg))
	signatureInBytes, err := hexutil.Decode(signature)
	if err != nil {
		return "", false, err
	}
	if signatureInBytes[64] == 27 || signatureInBytes[64] == 28 {
		signatureInBytes[64] -= 27
	}
	pubKey, err := crypto.SigToPub(newMsgHash.Bytes(), signatureInBytes)

	if err != nil {
		return "", false, err
	}

	//Get address from public key
	walletAddress := crypto.PubkeyToAddress(*pubKey)

	localData, exists := flowid.Data[flowId]
	if !exists {
		return "", false, errors.New("flow id not found")
	}
	if time.Now().Sub(localData.Timestamp) > 1*time.Hour {
		return "", false, errors.New("flow id expired for the request")
	}
	if strings.EqualFold(localData.WalletAddress, walletAddress.String()) {
		return localData.WalletAddress, true, nil
	} else {
		return "", false, nil
	} //equate the wallet address from the flow id and the reeived wallet address

}
