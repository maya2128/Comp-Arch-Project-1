package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define a simple map to map opcode patterns to instruction mnemonics.
var opcodeToMnemonic = map[string]string{
	"000101":                     "B",
	"10001010000":                "AND",
	"10001011000":                "ADD",
	"1001000100":                 "ADDI",
	"10101010000":                "ORR",
	"10110100":                   "CBZ",
	"10110101":                   "CBNZ",
	"11001011000":                "SUB",
	"1101000100":                 "SUBI",
	"110100101":                  "MOVZ",
	"111100101":                  "MOVK",
	"11010011010":                "LSR",
	"11010011011":                "LSL",
	"11111000000":                "STUR",
	"11111000010":                "LDUR",
	"11010011100":                "ASR",
	"0x0":                        "NOP",
	"11101010000":                "EOR",
	"11111101101111101111100111": "BREAK",
	"000000":                     " ",
}

func main() {
	inputFileName := "test7_bin.txt"
	outputFileName := "team7_out.txt"

	// Open the input file for reading.
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file for writing.
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	//defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		binaryInstruction := strings.TrimSpace(line)
		disassembled, err := DisassembleInstruction(binaryInstruction)
		if err != nil {
			fmt.Printf("Error disassembling instruction: %v\n", err)
			continue
		}
		fmt.Fprintf(outputFile, "%s\n", disassembled)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	fmt.Println("Disassembly complete. Output written to", outputFileName)
}

func DisassembleInstruction(binaryInstruction string) (string, error) {
	// Assuming a fixed length binary instruction format, e.g., 16 characters.
	if len(binaryInstruction) != 32 {
		return "", fmt.Errorf("Invalid instruction length: %s", binaryInstruction)
	}

	// Extract the opcode part to determine the instruction mnemonic.
	opcode := binaryInstruction[:11]

	// Check if the opcode exists in the map.
	mnemonic, found := opcodeToMnemonic[opcode]
	if !found {
		return "", fmt.Errorf("Opcode not found: %s", opcode)
	}

	// Construct the disassembled instruction.
	disassembled := fmt.Sprintf("%s %s", mnemonic, binaryInstruction[11:])

	return disassembled, nil
}
