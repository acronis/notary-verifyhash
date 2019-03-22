package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/onrik/gomerkle"
)

var (
	errInvalidCert        = errors.New("Invalid cert")
	errInvalidRoot        = errors.New("Invalid root")
	errInvalidProof       = errors.New("Invalid proof")
	errInvalidProofFormat = errors.New("Invalid proof format")
	errInvalidETag        = errors.New("Invalid eTag")
)

func verifyProof(proof []byte, root, cert, eTag string) error {
	if p, err := proofFromJSON(proof); err == nil {
		return verifyMerkleProof(p, root, eTag)
	}

	patriciaProof, err := patriciaTrieProofFromJSON(proof)
	if err != nil {
		return errInvalidProofFormat
	}

	return verifyPatriciaTrieProof(patriciaProof, root, cert, eTag)
}

// Merkle tree
func verifyMerkleProof(proof gomerkle.Proof, root, eTag string) error {
	rootByte, err := hex.DecodeString(root)
	if err != nil {
		return errInvalidRoot
	}

	tree := gomerkle.NewTree(sha256.New())
	eTagByte, _ := hex.DecodeString(eTag)
	valid := tree.VerifyProof(proof, rootByte, eTagByte)
	if valid {
		return nil
	}

	value := sha256.Sum256([]byte(eTag))
	valid = tree.VerifyProof(proof, rootByte, value[:])
	if !valid {
		return errors.New("Invalid proof or eTag")
	}

	return nil
}

func proofFromJSON(data []byte) (gomerkle.Proof, error) {
	hexProof := []map[string]string{}
	err := json.Unmarshal(data, &hexProof)
	if err != nil {
		return nil, err
	}

	proof := make(gomerkle.Proof, len(hexProof))
	for i := range hexProof {
		proof[i] = map[string][]byte{}
		for key, value := range hexProof[i] {
			proof[i][key], err = hex.DecodeString(value)
			if err != nil {
				return proof, err
			}
		}
	}

	return proof, nil
}

// Merkle patricia trie
func verifyPatriciaTrieProof(proof []rlp.RawValue, root, cert, eTag string) error {
	certByte, err := hex.DecodeString(cert)
	if err != nil {
		return errInvalidCert
	}

	value, err := trie.VerifyProof(common.HexToHash(root), certByte, proof)
	if err != nil {
		return errInvalidProof
	}
	equal := bytes.Equal(value, []byte(eTag))
	if !equal {
		return errInvalidETag
	}
	return nil
}

func patriciaTrieProofFromJSON(data []byte) ([]rlp.RawValue, error) {
	hexProof := []string{}
	err := json.Unmarshal(data, &hexProof)
	if err != nil {
		return nil, err
	}

	proof := make([]rlp.RawValue, len(hexProof))
	for i, v := range hexProof {
		proof[i], err = hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
	}

	return proof, nil
}
