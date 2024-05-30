package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
)

// ---

type Input struct {
	Packet                []byte
	AesKey                []byte
	Nonce                 []byte
	PreCounterBlockSuffix []byte
}

type inputJson struct {
	Packet string `json:"packet"`
	AesKey string `json:"aes-key"`
	Nonce  string `json:"nonce"`

	// Note that the word "couter" indicates suffix for PreCounterBlock (J0)
	// which is 0x00000001 for 12-byte nonce (IV).
	// See NIST SP 800-38D, section 7.1.
	PreCounterBlockSuffix string `json:"counter"`
}

func (input *Input) ReadJsonFile() error {
	if len(os.Args) != 2 {
		return errors.New("Input JSON file not provided")
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}

	rawJson := new(inputJson)
	if err := json.Unmarshal(raw, rawJson); err != nil {
		return err
	}

	if input.Packet, err = hex.DecodeString(rawJson.Packet); err != nil {
		return err
	}
	if input.AesKey, err = hex.DecodeString(rawJson.AesKey); err != nil {
		return err
	}
	if input.Nonce, err = hex.DecodeString(rawJson.Nonce); err != nil {
		return err
	}
	if input.PreCounterBlockSuffix, err = hex.DecodeString(rawJson.PreCounterBlockSuffix); err != nil {
		return err
	}

	return nil
}

// ---

type Output struct {
	Key                   []byte
	Nonce                 []byte
	Packet                []byte
	CipherText            []byte
	PreCounterBlockSuffix []byte
}

type outputJson struct {
	Key        []string `json:"key"`
	Nonce      []string `json:"nonce"`
	Packet     []string `json:"packet"`
	CipherText []string `json:"ciphertext"`

	// Same as `inputJson` struct.
	PreCounterBlockSuffix []string `json:"counter"`
}

func (output *Output) WriteJsonFile() error {
	o := new(outputJson)

	o.Key = toStringSlice(output.Key)
	o.Nonce = toStringSlice(output.Nonce)
	o.Packet = toStringSlice(output.Packet)
	o.CipherText = toStringSlice(output.CipherText)
	o.PreCounterBlockSuffix = toStringSlice(output.PreCounterBlockSuffix)

	dat, err := json.Marshal(o)
	if err != nil {
		return err
	}

	if err := os.WriteFile("result.json", dat, 0644); err != nil {
		return err
	}

	return nil
}
