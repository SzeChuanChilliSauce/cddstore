package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/tendermint/tendermint/crypto/merkle"
)

func main() {
	keyPtr := flag.String("key", "", "key")
	valuePtr := flag.String("value", "", "value")
	rootPtr := flag.String("root", "", "root hash in hex encoded form")
	proofPtr := flag.String("proof", "", "proof in json form")

	flag.Parse()

	err := verify(*proofPtr, *rootPtr, *keyPtr, *valuePtr)
	if err != nil {
		fmt.Printf("kv onchain verify (%s,%s) failed\n", *keyPtr, *valuePtr)
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Printf("kv onchain verify (%s,%s) succeeded\n", *keyPtr, *valuePtr)
}

func verify(proofJson, rootHex, key, value string) error {
	merkleProof := &merkle.Proof{Ops: make([]merkle.ProofOp, 1)}
	err := merkleProof.UnmarshalJSON([]byte(proofJson))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return err
	}

	rootBytes, err := hex.DecodeString(rootHex)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return err
	}

	prf := rootmulti.DefaultProofRuntime()
	return prf.VerifyValue(merkleProof, rootBytes, fmt.Sprintf("/%s", key), []byte(value))
}