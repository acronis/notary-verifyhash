package main

import (
	"bytes"
	"testing"
)

type setDataTest struct {
	input          dataVerification
	expectedOutput interface{}
}

var setDataTests = []setDataTest{
	{
		dataVerification{
			certificate: "e6b0b462e41a0acf4246a30686b069ccf178e1d1d295e7cafac41a4e9ffd7908",
			root:        "69208d1779de9eb8e2b01fc45a1b9f21e4ec75bcb1fabe62a4b6ba4d197df478",
			proof:       `["f843a120e6b0b462e41a0acf4246a30686b069ccf178e1d1d295e7cafac41a4e9ffd7908a03734366563333637303637346636386663326137623434646531353039643530"]`,
			object:      "",
		},
		nil,
	},
}

func TestSetData(t *testing.T) {
	for _, v := range setDataTests {
		var data dataVerification
		err := data.setData(v.input.certificate, v.input.root, v.input.proof)
		if err != v.expectedOutput {
			t.Errorf("Expected hashObject to return (%v), Found (%v)\n", v.expectedOutput, err)
		}
	}
}

type getValFromTreeTest struct {
	input          dataVerification
	expectedOutput []byte
}

var getValFromTreeTests = []getValFromTreeTest{
	{
		dataVerification{
			certificate: "e6b0b462e41a0acf4246a30686b069ccf178e1d1d295e7cafac41a4e9ffd7908",
			root:        "69208d1779de9eb8e2b01fc45a1b9f21e4ec75bcb1fabe62a4b6ba4d197df478",
			proof:       `["f843a120e6b0b462e41a0acf4246a30686b069ccf178e1d1d295e7cafac41a4e9ffd7908a03734366563333637303637346636386663326137623434646531353039643530"]`,
			object:      "",
		},
		[]byte("746ec3670674f68fc2a7b44de1509d50"),
	},
}

func TestGetValFromTree(t *testing.T) {
	for _, v := range getValFromTreeTests {
		output, err := v.input.getValFromTree()
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, v.expectedOutput) {
			t.Errorf("Expected hashObject to return (%v), Found (%v)\n", v.expectedOutput, err)
		}
	}
}

func TestMain(m *testing.M) {
}
