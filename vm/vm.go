/*
The MIT License (MIT)

Copyright (c) 2016-2018 Ivan Dejanovic

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
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
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
	INST_MASK     uint64 = 0xFF
	REG_MASK      uint64 = 0xF
	FULL_REG_MASK uint64 = 0x1F
)

//Rotation values
const (
	INST_ROTATION        int = 8
	OPERAND_ONE_ROTATION int = 12
	OPERAND_TWO_ROTATION int = 15
	FULL_REG_ROTATION    int = 13
)

//Register index
const (
	IP_INDEX  uint = 8
	SP_INDEX  uint = 9
	ACC_INDEX uint = 10
)

// Opcodes
const (
	HALT  uint = 0x00
	IN_N  uint = 0x01
	OUT_N uint = 0x02
	IN_S  uint = 0x03
	OUT_S uint = 0x04
	ADD   uint = 0x05
	SUB   uint = 0x06
	MUL   uint = 0x07
	DIV   uint = 0x08
	CON   uint = 0x09
	PUSH  uint = 0x0A
	POP   uint = 0x0B
	LD    uint = 0x0C
	ST    uint = 0x0D
	JMP   uint = 0x0E
	JGR   uint = 0x0F
	JGE   uint = 0x10
	JEQ   uint = 0x11
	JNE   uint = 0x12
	JLE   uint = 0x13
	JLS   uint = 0x14
)

const (
	REG_NUM          uint   = 8
	MEM_CAP          uint64 = 50000
	VAR_MEM_TOP      uint64 = 50000
	VAR_MEM_BOTTOM   uint64 = 40001
	STACK_MEM_TOP    uint64 = 40000
	STACK_MEM_BOTTOM uint64 = 30001
)

type vm struct {
	ip  uint64
	sp  uint64
	acc uint64
	reg [REG_NUM]uint64
	mem [MEM_CAP]uint64
}

func Execute(code []string) {
	vm := &vm{ip: 0, sp: 0, acc: 0}

	vm.load(code)
	vm.execute()
}

func (vm *vm) load(code []string) {
	//Create and initialize op codes map
	opCodes := make(map[uint]string)

	opCodes[HALT] = common.HALT
	opCodes[IN_N] = common.IN_N
	opCodes[OUT_N] = common.OUT_N
	opCodes[IN_S] = common.IN_S
	opCodes[OUT_S] = common.OUT_S
	opCodes[ADD] = common.ADD
	opCodes[SUB] = common.SUB
	opCodes[MUL] = common.MUL
	opCodes[DIV] = common.DIV
	opCodes[CON] = common.CON
	opCodes[PUSH] = common.PUSH
	opCodes[POP] = common.POP
	opCodes[LD] = common.LD
	opCodes[ST] = common.ST
	opCodes[JMP] = common.JMP
	opCodes[JGR] = common.JGR
	opCodes[JGE] = common.JGE
	opCodes[JEQ] = common.JEQ
	opCodes[JNE] = common.JNE
	opCodes[JLE] = common.JLE
	opCodes[JLS] = common.JLS

	for _, inst := range code {
		fmt.Println(inst)
	}
}

func (vm *vm) execute() {
	for {
		inst := vm.mem[vm.ip]
		opcode := uint(bits.RotateLeft64(inst, INST_ROTATION) & INST_MASK)

		switch opcode {
		case HALT:
			break
		case IN_N:
			vm.processInN()
		case OUT_N:
			vm.processOutN()
		case IN_S:
			vm.processInS()
		case OUT_S:
			vm.processOutS()
		case ADD:
			vm.processAdd()
		case SUB:
			vm.processSub()
		case MUL:
			vm.processMul()
		case DIV:
			vm.processDiv()
		case CON:
			vm.processCon()
		case PUSH:
			vm.processPush()
		case POP:
			vm.processPop()
		case LD:
			vm.processLd()
		case ST:
			vm.processSt()
		case JMP:
			vm.processJmp()
		case JGR:
			vm.processJgr()
		case JGE:
			vm.processJge()
		case JEQ:
			vm.processJeq()
		case JNE:
			vm.processJne()
		case JLE:
			vm.processJle()
		case JLS:
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
	operand1 := vm.getOperand(OPERAND_ONE_ROTATION)
	operand2 := vm.getOperand(OPERAND_TWO_ROTATION)
	result := operand1 + operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processSub() {
	operand1 := vm.getOperand(OPERAND_ONE_ROTATION)
	operand2 := vm.getOperand(OPERAND_TWO_ROTATION)
	result := operand1 - operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processMul() {
	operand1 := vm.getOperand(OPERAND_ONE_ROTATION)
	operand2 := vm.getOperand(OPERAND_TWO_ROTATION)
	result := operand1 * operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processDiv() {
	operand1 := vm.getOperand(OPERAND_ONE_ROTATION)
	operand2 := vm.getOperand(OPERAND_TWO_ROTATION)
	result := operand1 / operand2
	vm.acc = math.Float64bits(result)
	vm.ip = vm.ip + 1
}

func (vm *vm) processCon() {

}

func (vm *vm) processPush() {
	//Check if stack is full
	if vm.sp <= STACK_MEM_BOTTOM {
		vm.setMemErrorFlags("Push attempt on full stack.")
		return
	}

	regIndex := vm.getRegisterIndex()
	var value uint64 = 0

	if regIndex < REG_NUM {
		value = vm.reg[regIndex]
	} else if regIndex == IP_INDEX {
		value = uint64(vm.ip)

	} else if regIndex == SP_INDEX {
		value = uint64(vm.sp)
	} else if regIndex == ACC_INDEX {
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
	if vm.sp > STACK_MEM_TOP {
		vm.setMemErrorFlags("Pop attempt on empty stack.")
		return
	}

	regIndex := vm.getRegisterIndex()
	value := vm.mem[vm.sp]
	vm.sp = vm.sp + 1

	if regIndex < REG_NUM {
		vm.reg[regIndex] = value
	} else if regIndex == IP_INDEX {
		vm.ip = value
	} else if regIndex == SP_INDEX {
		vm.sp = value
	} else if regIndex == ACC_INDEX {
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
	operandReg := bits.RotateLeft64(inst, rotation) & REG_MASK
	operand := math.Float64frombits(operandReg)

	return operand
}

func (vm *vm) getRegisterIndex() uint {
	inst := vm.mem[vm.ip]
	regIndex := uint(bits.RotateLeft64(inst, FULL_REG_ROTATION) & FULL_REG_MASK)

	return regIndex
}
