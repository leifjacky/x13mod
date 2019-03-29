package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestX13BCD(t *testing.T) {
	targetHex := "fd21da3c2b0b6e5f438d365ef3f9a2e3a5018b9402fff79bec0a2e47657e802b"
	targetBytes, _ := hex.DecodeString(targetHex)
	headerHex := "600000006df1d486138ba00b9ceabc42c71c9d1925ea25896ee5555e78507642be52a4d761f93e1f507975d05a33de16bea7784f5cf492aa8e0eda6bf71029d78e3d939a5a2fa97e1d08379f000b3a0e"
	headerBytes, _ := hex.DecodeString(headerHex)
	rawData := bytes.NewBuffer(nil)
	for i := 0; i < len(headerBytes); i += 4 {
		for j := 0; j < 4; j++ {
			rawData.WriteByte(headerBytes[i+(4-j-1)])
		}
	}
	rawDataBytes := rawData.Bytes()
	hash := X13BCD(rawDataBytes)
	for i,j := 0, len(hash) - 1; i < j ;i,j = i+1,j-1{
		hash[i], hash[j] = hash[j], hash[i]
	}
	if !bytes.Equal(hash, targetBytes){
		t.Fatalf("wrong hash result, got %s, should be %s", hex.EncodeToString(hash), targetHex)
	}
}
