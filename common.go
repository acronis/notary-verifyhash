package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
)

type objectJSON struct {
	ETag      string `json:"eTag" `
	Key       string `json:"key" `
	Sequencer string `json:"sequencer"`
	Size      int    `json:"size"`
}

func proofFromJSON(j []byte) (rawProof []rlp.RawValue, err error) {
	var arrProof []string

	if json.Unmarshal(j, &arrProof) != nil {
		return rawProof, err
	}

	for _, v := range arrProof {
		var b []byte
		if b, err = hex.DecodeString(v); err != nil {
			return rawProof, err
		}
		rawProof = append(rawProof, b)
	}
	return rawProof, nil
}

func hashObject(b []byte) (hash string, err error) {
	var object objectJSON
	json.Unmarshal(b, &object)
	if !objectIsValid(object) {
		return hash, fmt.Errorf("%v", "Invalid Object")
	}
	byteObject, _ := json.Marshal(object)
	hashByte := sha256.Sum256(byteObject)
	hash = hex.EncodeToString(hashByte[:])
	return
}

func objectIsValid(object objectJSON) bool {
	if len(object.ETag) == 0 || len(object.Key) == 0 || len(object.Sequencer) == 0 {
		return false
	}
	return true
}
