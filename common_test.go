package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
)

type proofFromJSONTest struct {
	input          string
	expectedOutput string
}

var proofFromJSONTests = []proofFromJSONTest{
	{
		`["f843a120357be57512aa15f40b6db08f1d3360e122e6fa2226bd559683c7675899df9a60a03132363030643434353761633661666138396539303334323333343536363537"]`,
		`["+EOhIDV75XUSqhX0C22wjx0zYOEi5voiJr1VloPHZ1iZ35pgoDEyNjAwZDQ0NTdhYzZhZmE4OWU5MDM0MjMzNDU2NjU3"]`,
	},
}

func TestProofFromJSON(t *testing.T) {
	for _, v := range proofFromJSONTests {
		var (
			r   []rlp.RawValue
			err error
		)
		if r, err = proofFromJSON([]byte(v.input)); err != nil {
			t.Error(err)
		}
		output, _ := json.Marshal(r)
		if !bytes.EqualFold(output, []byte(v.expectedOutput)) {
			t.Errorf("Expected hashObject to return (%s), Found (%s)\n", v.expectedOutput, output)
		}
	}
}

type hashObjectTest struct {
	input          []byte
	expectedOutput string
}

var hashObjectTests = []hashObjectTest{
	{
		[]byte(`
    {
        "eTag": "12600d4457ac6afa89e9034233456657",
        "key": "现代中文 - Better Chinese.jpg",
        "sequencer": "botary_bot",
        "size": 92040
    }
    `),
		"357be57512aa15f40b6db08f1d3360e122e6fa2226bd559683c7675899df9a60",
	},
}

func TestHashObject(t *testing.T) {
	for _, v := range hashObjectTests {
		output, err := hashObject(v.input)
		if err != nil {
			t.Error(err)
		}
		if !strings.EqualFold(output, v.expectedOutput) {
			t.Errorf("Expected hashObject to return (%s), Found (%s)\n", v.expectedOutput, output)
		}
	}
}

type objectIsValidTest struct {
	input          objectJSON
	expectedOutput bool
}

var objectIsValidTests = []objectIsValidTest{
	{
		objectJSON{
			ETag:      "12600d4457ac6afa89e9034233456657",
			Key:       "现代中文 - Better Chinese.jpg",
			Sequencer: "botary_bot",
			Size:      92040,
		},
		true,
	},
	{
		objectJSON{
			ETag:      "",
			Key:       "现代中文 - Better Chinese.jpg",
			Sequencer: "botary_bot",
			Size:      92040,
		},
		false,
	},
	{
		objectJSON{
			ETag:      "12600d4457ac6afa89e9034233456657",
			Key:       "",
			Sequencer: "botary_bot",
			Size:      92040,
		},
		false,
	},
	{
		objectJSON{
			ETag:      "12600d4457ac6afa89e9034233456657",
			Key:       "现代中文 - Better Chinese.jpg",
			Sequencer: "",
			Size:      92040,
		},
		false,
	},
}

func TestObjectIsValid(t *testing.T) {
	for _, v := range objectIsValidTests {
		output := objectIsValid(v.input)

		if output != v.expectedOutput {
			t.Errorf("Expected hashObject to return (%v), Found (%v)\n", v.expectedOutput, output)
		}
	}
}
