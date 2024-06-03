package main

import (
	"bytes"
	"testing"
)

var keyBytes = []byte{
	0x00, 0x01, 0x02, 0x03,
	0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b,
	0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13,
	0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b,
	0x1c, 0x1d, 0x1e, 0x1f,
}
var noncebytes = []byte{
	0x01, 0x02, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x4a,
	0x00, 0x00, 0x00, 0x09,
}
var couter = []byte{
	0x00, 0x00, 0x00, 0x02,
}

var plaintextBytes = []byte{
	0x4c, 0x61, 0x64, 0x69, 0x65, 0x73, 0x20, 0x61,
	0x65, 0x6d, 0x65, 0x6e, 0x20, 0x6f, 0x66, 0x20,
}

var testAesParam = aesParam{
	key:                   keyBytes,
	nonce:                 noncebytes,
	preCounterBlockSuffix: couter,
}

var aesCipherRes = []byte{
	0x4e, 0xbc, 0x4c, 0xf5,
	0x24, 0xae, 0x65, 0x39,
	0x55, 0x4a, 0xc3, 0xdf,
	0x6c, 0x05, 0x08, 0x5d,
}

var aesTagRes = []byte{
	0x45, 0x04, 0x78, 0xfe,
	0xa8, 0xa9, 0xff, 0x25,
	0x4c, 0x2a, 0xff, 0x15,
	0x88, 0xd7, 0x1b, 0x21,
}

var testChachaPolyParam = chachaPolyParam{
	key:                   keyBytes,
	nonce:                 noncebytes,
	preCounterBlockSuffix: couter,
}

var chaCipherRes = []byte{
	0x7b, 0x4a, 0xf9, 0x97,
	0xc8, 0xd6, 0x01, 0x94,
	0xfd, 0xed, 0x91, 0xb0,
	0x75, 0xc3, 0x84, 0x4e,
}

var chaTagRes = []byte{
	0xea, 0x86, 0xe7, 0x59,
	0xe6, 0x50, 0x84, 0xb3,
	0xd5, 0x24, 0x18, 0x9b,
	0x9c, 0xc1, 0x6e, 0xf7,
}

func TestAesParamEncrypt(t *testing.T) {
	cipher, tag, err := testAesParam.Encrypt(plaintextBytes)
	if err != nil {
		t.Error("test failed:", err)
	}
	t.Log("Cipher:", cipher)
	t.Log("Tag:", tag)
	if !bytes.Equal(cipher, aesCipherRes) {
		t.Errorf("Cipher not matched: want '%s', got '%s'", aesCipherRes, cipher)
	}
	if !bytes.Equal(tag, aesTagRes) {
		t.Errorf("Tag not matched: want '%s' got '%s'", aesTagRes, tag)
	}
}

func TestChachaPolyParamEncrypt(t *testing.T) {
	cipher, tag, err := testChachaPolyParam.Encrypt(plaintextBytes)
	if err != nil {
		t.Error("test failed:", err)
	}
	t.Log("Cipher:", cipher)
	t.Log("Tag:", tag)
	if !bytes.Equal(cipher, chaCipherRes) {
		t.Errorf("Cipher not matched: want '%s', got '%s'", chaCipherRes, cipher)
	}
	if !bytes.Equal(tag, chaTagRes) {
		t.Errorf("Tag not matched: want '%s' got '%s'", chaTagRes, tag)
	}
}