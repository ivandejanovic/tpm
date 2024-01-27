/*
The MIT License (MIT)

Copyright (c) 2016-2023 Ivan Dejanovic

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOopLdERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF opConTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN opConNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package vm

import (
	"fmt"
	"github.com/ivandejanovic/tpm/common"
	"math"
	"math/bits"
)

// Bit Masks
const (
	instMask    uint64 = 0xFF
	regMask     uint64 = 0xF
	fullRegMask uint64 = 0x1F
)

//Rotation values
const (
	instRotation       int = 8
	operandOneRotation int = 12
	operandTwoRotation int = 15
	fullRegRotation    int = 13
)

//Register index
const (
	ipIndex  uint = 8
	spIndex  uint = 9
	accIndex uint = 10
)

// Opcodes
const (
	opHalt uint = 0x00
	opInN  uint = 0x01
	opOutN uint = 0x02
	opInS  uint = 0x03
	opOutS uint = 0x04
	opAdd  uint = 0x05
	opSub  uint = 0x06
	opMul  uint = 0x07
	opDiv  uint = 0x08
	opCon  uint = 0x09
	opPush uint = 0x0A
	opPop  uint = 0x0B
	opLd   uint = 0x0C
	opSt   uint = 0x0D
	opJmp  uint = 0x0E
	opJgr  uint = 0x0F
	opJge  uint = 0x10
	opJeq  uint = 0x11
	opJne  uint = 0x12
	opJle  uint = 0x13
	opJls  uint = 0x14
)

const (
	regNum         uint   = 8
	memCap         uint64 = 50000
	varMemTop      uint64 = 50000
	varMemBottom   uint64 = 40001
	stackMemTop    uint64 = 40000
	stackMemBottom uint64 = 30001
)

type vm struct {
	ip  uint64
	sp  uint64
	acc uint64
	reg [regNum]uint64
	mem [memCap]uint64
}

func Execute(code []string) {
	vm := &vm{ip: 0, sp: 0, acc: 0}

	vm.load(code)
	vm.execute()
}

func (vm *vm) load(code []string) {
	//Create and initialize op codes map
	opCodes := make(map[uint]string)

	opCodes[opHalt] = common.Halt
	opCodes[opInN] = common.InN
	opCodes[opOutN] = common.OutN
	opCodes[opInS] = common.InS
	opCodes[opOutS] = common.OutS
	opCodes[opAdd] = common.Add
	opCodes[opSub] = common.Sub
	opCodes[opMul] = common.Mul
	opCodes[opDiv] = common.Div
	opCodes[opCon] = common.Con
	opCodes[opPush] = common.Push
	opCodes[opPop] = common.Pop
	opCodes[opLd] = common.Ld
	opCodes[opSt] = common.St
	opCodes[opJmp] = common.Jmp
	opCodes[opJgr] = common.Jgr
	opCodes[opJge] = common.Jge
	opCodes[opJeq] = common.Jeq
	opCodes[opJne] = common.Jne
	opCodes[opJle] = common.Jle
	opCodes[opJls] = common.Jls

	for _, inst := range code {
		fmt.Println(inst)
	}
}

func (vm *vm) execute() {
	for {
		inst := vm.mem[vm.ip]
		opcode := uint(bits.RotateLeft64(inst, instRotation) & instMask)

		switch opcode {
		case opHalt:
			break
		case opInN:
			vm.processInN()
		case opOutN:
			vm.processOutN()
		case opInS:
			vm.processInS()
		case opOutS:
			vm.processOutS()
		case opAdd:
			vm.processAdd()
		case opSub:
			vm.processSub()
		case opMul:
			vm.processMul()
		case opDiv:
			vm.processDiv()
		case opCon:
			vm.processCon()
		case opPush:
			vm.processPush()
		case opPop:
			vm.processPop()
		case opLd:
			vm.processLd()
		case opSt:
			vm.processSt()
		case opJmp:
			vm.processJmp()
		case opJgr:
			vm.processJgr()
		case opJge:
			vm.processJge()
		case opJeq:
			vm.processJeq()
		case opJne:
			vm.processJne()
		case opJle:
			vm.processJle()
		case opJls:
			vm.processJls()
		default:
			fmt.Println("Unknow instruction. Aborting.")
			break
		}

		if vm.hasErrors() {
			vm.throwException()
			break
		}
	}
}

func (vm *vm) processInN() {
	var input float64 = 0.0
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println("")
		vm.setIOErrorFlag()
		return
	}
	vm.acc = math.Float64bits(input)
	vm.ip = vm.ip + 1
}

func (vm *vm) processOutN() {
	output := math.Float64frombits(vm.acc)
	fmt.Println(output)
	vm.ip = vm.ip + 1
}

func (vm *vm) processInS() {

}

func (vm *vm) processOutS() {

}

func (vm *vm) processAdd() {
	operand1 := vm.getOperand(operandOneRotation)
	operand2 := vm.getOperand(operandTwoRotation)
	result := operand1 + operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processSub() {
	operand1 := vm.getOperand(operandOneRotation)
	operand2 := vm.getOperand(operandTwoRotation)
	result := operand1 - operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processMul() {
	operand1 := vm.getOperand(operandOneRotation)
	operand2 := vm.getOperand(operandTwoRotation)
	result := operand1 * operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processDiv() {
	operand1 := vm.getOperand(operandOneRotation)
	operand2 := vm.getOperand(operandTwoRotation)
	result := operand1 / operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processCon() {

}

func (vm *vm) processPush() {
	//Check if stack is full
	if vm.sp <= stackMemBottom {
		vm.setMemErrorFlags("Push attempt on full stack.")
		return
	}

	regIndex := vm.getRegisterIndex()
	var value uint64 = 0

	if regIndex < regNum {
		value = vm.reg[regIndex]
	} else if regIndex == ipIndex {
		value = uint64(vm.ip)

	} else if regIndex == spIndex {
		value = uint64(vm.sp)
	} else if regIndex == accIndex {
		value = vm.acc
	} else {
		vm.setMemErrorFlags("Push attempt on unkown register.")
		return
	}

	vm.mem[vm.sp] = value
	vm.sp = vm.sp - 1
}

func (vm *vm) processPop() {
	//Check if stack is empty
	if vm.sp > stackMemTop {
		vm.setMemErrorFlags("Pop attempt on empty stack.")
		return
	}

	regIndex := vm.getRegisterIndex()
	value := vm.mem[vm.sp]
	vm.sp = vm.sp + 1

	if regIndex < regNum {
		vm.reg[regIndex] = value
	} else if regIndex == ipIndex {
		vm.ip = value
	} else if regIndex == spIndex {
		vm.sp = value
	} else if regIndex == accIndex {
		vm.acc = value
	} else {
		vm.setMemErrorFlags("Pop attempt on unkown register.")
		return
	}
}

func (vm *vm) processLd() {
	//regIndex := vm.getRegisterIndex()
}

func (vm *vm) processSt() {
	//regIndex := vm.getRegisterIndex()
}

func (vm *vm) processJmp() {

}

func (vm *vm) processJgr() {

}

func (vm *vm) processJge() {

}

func (vm *vm) processJeq() {

}

func (vm *vm) processJne() {

}

func (vm *vm) processJle() {

}

func (vm *vm) processJls() {

}

func (vm *vm) setIOErrorFlag() {
	fmt.Println("IO error flag set!")
}

func (vm *vm) setMemErrorFlags(err string) {
	fmt.Println(err)
}

func (vm *vm) hasErrors() bool {
	return false
}

func (vm *vm) throwException() {
	fmt.Println("Exception thrown!")
}

func (vm *vm) getOperand(rotation int) float64 {
	inst := vm.mem[vm.ip]
	operandReg := bits.RotateLeft64(inst, rotation) & regMask
	operand := math.Float64frombits(operandReg)

	return operand
}

func (vm *vm) getRegisterIndex() uint {
	inst := vm.mem[vm.ip]
	regIndex := uint(bits.RotateLeft64(inst, fullRegRotation) & fullRegMask)

	return regIndex
}
