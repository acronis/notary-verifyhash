package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	// imports as package "cli"
)

type dataVerification struct {
	certificate string
	object      string
	root        string
	proof       string
}

func (d *dataVerification) setData(certificate, root, proof string) (err error) {
	d.certificate = certificate
	d.root = root
	d.proof = proof

	var messages []string
	if len(d.certificate) == 0 {
		messages = append(messages, "certificate")
	}
	if len(d.root) == 0 {
		messages = append(messages, "merkle root")
	}
	if len(d.proof) == 0 {
		messages = append(messages, "merkle proof")
	}
	if len(messages) != 0 {
		return fmt.Errorf("%v is not entered \n", strings.Join(messages, ", "))
	}

	return nil
}

func (d *dataVerification) getValFromTree() (val []byte, err error) {
	var (
		rootByte  common.Hash
		proofByte []rlp.RawValue
		hashByte  []byte
	)

	if proofByte, err = proofFromJSON([]byte(d.proof)); err != nil {
		return val, fmt.Errorf("Invalid proof")
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("\nInvalid root", r)
			}
		}()

		rootByte = common.HexToHash(d.root)
	}()

	if hashByte, err = hex.DecodeString(d.certificate); err != nil {
		return val, fmt.Errorf("Invalid certificate ID")
	}

	if val, err = trie.VerifyProof(rootByte, hashByte, proofByte); err != nil {
		return val, fmt.Errorf("Verification failed")
	}

	return val, err
}
