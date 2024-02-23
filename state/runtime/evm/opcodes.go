package evm

import (
	"fmt"
)

// OpCode is the EVM operation code
type OpCode int

const (
	// STOP halts execution of the contract
	STOP OpCode = 0x0

	// ADD performs (u)int256 addition modulo 2**256
	ADD = 0x01

	// MUL performs (u)int256 multiplication modulo 2**256
	MUL = 0x02

	// SUB performs (u)int256 subtraction modulo 2**256
	SUB = 0x03

	// DIV performs uint256 division
	DIV = 0x04

	// SDIV performs int256 division
	SDIV = 0x05

	// MOD performs uint256 modulus
	MOD = 0x06

	// SMOD performs int256 modulus
	SMOD = 0x07

	// ADDMOD performs (u)int256 addition modulo N
	ADDMOD = 0x08

	// MULMOD performs (u)int256 multiplication modulo N
	MULMOD = 0x09

	// EXP performs uint256 exponentiation modulo 2**256
	EXP = 0x0A

	// SIGNEXTEND performs sign extends x from (b + 1) * 8 bits to 256 bits.
	SIGNEXTEND = 0x0B

	// LT performs int256 comparison
	LT = 0x10

	// GT performs int256 comparison
	GT = 0x11

	// SLT performs int256 comparison
	SLT = 0x12

	// SGT performs int256 comparison
	SGT = 0x13

	// EQ performs (u)int256 equality
	EQ = 0x14

	// ISZERO checks if (u)int256 is zero
	ISZERO = 0x15

	// AND performs 256-bit bitwise and
	AND = 0x16

	// OR performs 256-bit bitwise or
	OR = 0x17

	// XOR performs 256-bit bitwise xor
	XOR = 0x18

	// NOT performs 256-bit bitwise not
	NOT = 0x19

	// BYTE returns the ith byte of (u)int256 x counting from most significant byte
	BYTE = 0x1A

	// SHL performs a shift left
	SHL = 0x1B

	// SHR performs a logical shift right
	SHR = 0x1C

	// SAR performs an arithmetic shift right
	SAR = 0x1D

	// SHA3 performs the keccak256 hash function
	SHA3 = 0x20

	// ADDRESS returns the address of the executing contract
	ADDRESS = 0x30

	// BALANCE returns the address balance in wei
	BALANCE = 0x31

	// ORIGIN returns the transaction origin address
	ORIGIN = 0x32

	// CALLER returns the message caller address
	CALLER = 0x33

	// CALLVALUE returns the message funds in wei
	CALLVALUE = 0x34

	// CALLDATALOAD reads a (u)int256 from message data
	CALLDATALOAD = 0x35

	// CALLDATASIZE returns the message data length in bytes
	CALLDATASIZE = 0x36

	// CALLDATACOPY copies the message data
	CALLDATACOPY = 0x37

	// CODESIZE returns the length of the executing contract's code in bytes
	CODESIZE = 0x38

	// CODECOPY copies the executing contract bytecode
	CODECOPY = 0x39

	// GASPRICE returns the gas price of the executing transaction, in wei per unit of gas
	GASPRICE = 0x3A

	// EXTCODESIZE returns the length of the contract bytecode at addr
	EXTCODESIZE = 0x3B

	// EXTCODECOPY copies the contract bytecode
	EXTCODECOPY = 0x3C

	// RETURNDATASIZE returns the size of the returned data from the last external call in bytes
	RETURNDATASIZE = 0x3D

	// RETURNDATACOPY copies the returned data
	RETURNDATACOPY = 0x3E

	// EXTCODEHASH returns the hash of the specified contract bytecode
	EXTCODEHASH = 0x3F

	// BLOCKHASH returns the hash of the specific block. Only valid for the last 256 most recent blocks
	BLOCKHASH = 0x40

	// COINBASE returns the address of the current block's miner
	COINBASE = 0x41

	// TIMESTAMP returns the current block's Unix timestamp in seconds
	TIMESTAMP = 0x42

	// NUMBER returns the current block's number
	NUMBER = 0x43

	// DIFFICULTY returns the current block's difficulty
	DIFFICULTY = 0x44

	// GASLIMIT returns the current block's gas limit
	GASLIMIT = 0x45

	// CHAINID returns the id of the chain
	CHAINID = 0x46

	// SELFBALANCE returns the balance of the current account
	SELFBALANCE = 0x47

	// BASEFEE returns the current base fee value
	BASEFEE = 0x48

	// POP pops a (u)int256 off the stack and discards it
	POP = 0x50

	// MLOAD reads a (u)int256 from memory
	MLOAD = 0x51

	// MSTORE writes a (u)int256 to memory
	MSTORE = 0x52

	// MSTORE8 writes a uint8 to memory
	MSTORE8 = 0x53

	// SLOAD reads a (u)int256 from storage
	SLOAD = 0x54

	// SSTORE writes a (u)int256 to storage
	SSTORE = 0x55

	// JUMP performs an unconditional jump
	JUMP = 0x56

	// JUMPI performs a conditional jump if condition is truthy
	JUMPI = 0x57

	// PC returns the program counter
	PC = 0x58

	// MSIZE returns the size of memory for this contract execution, in bytes
	MSIZE = 0x59

	// GAS returns the remaining gas
	GAS = 0x5A

	// JUMPDEST corresponds to a possible jump destination
	JUMPDEST = 0x5B

	// PUSH0 represents the number 0
	PUSH0 = 0x5F

	// PUSH1 to PUSH32 place i bytes item on stack
	PUSH1 = 0x60
	PUSH2 = 0x61
	PUSH3 = 0x62
	PUSH4 = 0x63
	PUSH5 = 0x64
	PUSH6 = 0x65
	PUSH7 = 0x66
	PUSH8 = 0x67
	PUSH9 = 0x68
	PUSH10 = 0x69
	PUSH11 = 0x6A
	PUSH12 = 0x6B
	PUSH13 = 0x6C
	PUSH14 = 0x6D
	PUSH15 = 0x6E
	PUSH16 = 0x6F
	PUSH17 = 0x70
	PUSH18 = 0x71
	PUSH19 = 0x72
	PUSH20 = 0x73
	PUSH21 = 0x74
	PUSH22 = 0x75
	PUSH23 = 0x76
	PUSH24 = 0x77
	PUSH25 = 0x78
	PUSH26 = 0x79
	PUSH27 = 0x7A
	PUSH28 = 0x7B
	PUSH29 = 0x7C
	PUSH30 = 0x7D
	PUSH31 = 0x7E
	PUSH32 = 0x7F

	// DUP1 to DUP16 duplicates the i-th stack item
	DUP1 = 0x80
	DUP2 = 0x81
	DUP3 = 0x82
	DUP4 = 0x83
	DUP5 = 0x84
	DUP6 = 0x85
	DUP7 = 0x86
	DUP8 = 0x87
	DUP9 = 0x88
	DUP10 = 0x89
	DUP11 = 0x8A
	DUP12 = 0x8B
	DUP13 = 0x8C
	DUP14 = 0x8D
	DUP15 = 0x8E
	DUP16 = 0x8F

	// SWAP1 to SWAP16 exchanges 1st and i-th stack items
	SWAP1 = 0x90
	SWAP2 = 0x91
	SWAP3 = 0x92
	SWAP4 = 0x93
	SWAP5 = 0x94
	SWAP6 = 0x95
	SWAP7 = 0x96
	SWAP8 = 0x97
	SWAP9 = 0x98
	SWAP10 = 0x99
	SWAP11 = 0x9A
	SWAP12 = 0x9B
	SWAP13 = 0x9C
	SWAP14 = 0x9D
	SWAP15 = 0x9E
	SWAP16 = 0x9F

	// LOG0 to LOG4 logs data
	LOG0 = 0xA0
	LOG1 = 0xA1
	LOG2 = 0xA2
	LOG3 = 0xA3
	LOG4 = 0xA4

	// CREATE creates a new contract
	CREATE = 0xF0

	// CALL calls a contract
	CALL = 0xF1

	// CALLCODE calls a contract with a copy of its own code
	CALLCODE = 0xF2

	// RETURN halts execution and returns data
	RETURN = 0xF3

	// DELEGATECALL calls a contract using the same state
	DELEGATECALL = 0xF4

	// CREATE2 creates a new contract with a salt
	CREATE2 = 0xF5

	// STATICCALL calls a contract with a static state
	STATICCALL = 0xFA

	// REVERT halts execution and reverts state changes
	REVERT = 0xFD

	// INVALID invalid instruction
	INVALID = 0xFE

	// SELFDESTRUCT halts execution and marks the account for deletion
	SELFDESTRUCT = 0xFF
)

var opCodeToString = map[OpCode]string{
	STOP:           "STOP",
	ADD:            "ADD",
	MUL:            "MUL",
	SUB:            "SUB",
	DIV:            "DIV",
	SDIV:           "SDIV",
	MOD:            "MOD",
	SMOD:           "SMOD",
	EXP:            "EXP",
	NOT:            "NOT",
	LT:             "LT",
	GT:             "GT",
	SLT:            "SLT",
	SGT:            "SGT",
	EQ:             "EQ",
	ISZERO:         "ISZERO",
	SIGNEXTEND:     "SIGNEXTEND",
	AND:            "AND",
	OR:             "OR",
	XOR:            "XOR",
	BYTE:           "BYTE",
	SHL:            "SHL",
	SHR:            "SHR",
	SAR:            "SAR",
	ADDMOD:         "ADDMOD",
	MULMOD:         "MULMOD",
	SHA3:           "SHA3",
	ADDRESS:        "ADDRESS",
	BALANCE:        "BALANCE",
	ORIGIN:         "ORIGIN",
	CALLER:         "CALLER",
	CALLVALUE:      "CALLVALUE",
	CALLDATALOAD:   "CALLDATALOAD",
	CALLDATASIZE:   "CALLDATASIZE",
	CALLDATACOPY:   "CALLDATACOPY",
	CODESIZE:       "CODESIZE",
	CODECOPY:       "CODECOPY",
	GASPRICE:       "GASPRICE",
	EXTCODESIZE:    "EXTCODESIZE",
	EXTCODECOPY:    "EXTCODECOPY",
	RETURNDATASIZE: "RETURNDATASIZE",
	RETURNDATACOPY: "RETURNDATACOPY",
	EXTCODEHASH:    "EXTCODEHASH",
	BLOCKHASH:      "BLOCKHASH",
	COINBASE:       "COINBASE",
	TIMESTAMP:      "TIMESTAMP",
	NUMBER:         "NUMBER",
	DIFFICULTY:     "DIFFICULTY",
	GASLIMIT:       "GASLIMIT",
	CHAINID:        "CHAINID",
	SELFBALANCE:    "SELFBALANCE",
	BASEFEE:        "BASEFEE",
	POP:            "POP",
	MLOAD:          "MLOAD",
	MSTORE:         "MSTORE",
	MSTORE8:        "MSTORE8",
	SLOAD:          "SLOAD",
	SSTORE:         "SSTORE",
	JUMP:           "JUMP",
	JUMPI:          "JUMPI",
	PC:             "PC",
	MSIZE:          "MSIZE",
	GAS:            "GAS",
	JUMPDEST:       "JUMPDEST",
	CREATE:         "CREATE",
	CALL:           "CALL",
	RETURN:         "RETURN",
	CALLCODE:       "CALLCODE",
	DELEGATECALL:   "DELEGATECALL",
	CREATE2:        "CREATE2",
	STATICCALL:     "STATICCALL",
	REVERT:         "REVERT",
	SELFDESTRUCT:   "SELFDESTRUCT",
	PUSH0:          "PUSH0",
}

func opCodesToString(from, to OpCode, str string) {
	c := 1
	if from == LOG0 {
		c = 0
	}

	for i := from; i <= to; i++ {
		opCodeToString[i] = fmt.Sprintf("%s%d", str, c)
		c++
	}
}

func init() {
	// write push
	opCodesToString(PUSH0, PUSH32, "PUSH")
	// write dup
	opCodesToString(DUP1, DUP16, "DUP")
	// write log
	opCodesToString(LOG0, LOG4, "LOG")
	// write swap
	opCodesToString(SWAP1, SWAP16, "SWAP")
}

func (op OpCode) String() string {
	return opCodeToString[op]
}
