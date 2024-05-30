package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// 1. Input
	input := new(InputJson)
	packet, aes, err := input.ReadFile()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 2. Add EDNS0 padding opt
	padding := new(DnsRROPT)
	padding.FillZero(512 - len(input.Packet) / 2 - 4)
	packet.AppendAdditionalRR(padding)
  paddingOnly := packet.Unmarshal()

	// 3. Encode 0x20
	if err := packet.question[0].Encode0x20(); err != nil {
		fmt.Println("error:", err)
		return
	}
	packet.Print()

	// 4. Encrypt w/ AES_256_GCM
	cipher, err := aes.Encrypt(packet.Unmarshal())
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Cipher")
	fmt.Printf("  Hex:     %s\n", hex.EncodeToString(cipher))
	fmt.Printf("  Length:  %d\n", len(cipher))

	// 5. Output
  output := new(OutputJson)
	if err := output.WriteFile(paddingOnly, cipher, aes); err != nil {
		fmt.Println("error:", err)
		return
	}
}
