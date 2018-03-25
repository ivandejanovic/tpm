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
	HALT  uint64 = 0x00
	IN_N  uint64 = 0x01
	OUT_N uint64 = 0x02
	IN_S  uint64 = 0x03
	OUT_S uint64 = 0x04
	ADD   uint64 = 0x05
	SUB   uint64 = 0x06
	MUL   uint64 = 0x07
	DIV   uint64 = 0x08
	CON   uint64 = 0x09
	PUSH  uint64 = 0x0A
	POP   uint64 = 0x0B
	LD    uint64 = 0x0C
	ST    uint64 = 0x0D
	JMP   uint64 = 0x0E
	JGR   uint64 = 0x0F
	JGE   uint64 = 0x10
	JEQ   uint64 = 0x11
	JNE   uint64 = 0x12
	JLE   uint64 = 0x13
	JLS   uint64 = 0x14
)

const (
	REG_NUM          uint = 8
	MEM_CAP          uint = 50000
	VAR_MEM_TOP      uint = 50000
	VAR_MEM_BOTTOM   uint = 40001
	STACK_MEM_TOP    uint = 40000
	STACK_MEM_BOTTOM uint = 30001
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
	opCodes := make(map[uint64]string)

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
		opcode := bits.RotateLeft64(inst, INST_ROTATION) & INST_MASK

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
	//let mut input_text = String::new();
	//        stdin().read_line(&mut input_text).expect(
	//            "failed to read from stdin",
	//        );
	//
	//        let trimmed = input_text.trim();
	//        match trimmed.parse::<f64>() {
	//            Ok(i) => {
	//                self.acc = unsafe { transmute::<f64, u64>(i) };
	//                self.ip = self.ip + 1;
	//            }
	//            Err(..) => self.set_io_error_flag(),
	//        };
}

func (vm *vm) processOutN() {
	//let output = unsafe { transmute::<u64, f64>(self.acc) };
	//        println!("{}", output);
	//        self.ip = self.ip + 1;
}

func (vm *vm) processInS() {

}

func (vm *vm) processOutS() {

}

func (vm *vm) processAdd() {
	//let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
	//        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);
	//
	//        let result = operand1 + operand2;
	//
	//        self.acc = unsafe { transmute::<f64, u64>(result) };
	//        self.ip = self.ip + 1;
}

func (vm *vm) processSub() {
	//let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
	//        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);
	//
	//        let result = operand1 - operand2;
	//
	//        self.acc = unsafe { transmute::<f64, u64>(result) };
	//        self.ip = self.ip + 1;
}

func (vm *vm) processMul() {
	//let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
	//        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);
	//
	//        let result = operand1 * operand2;
	//
	//        self.acc = unsafe { transmute::<f64, u64>(result) };
	//        self.ip = self.ip + 1;
}

func (vm *vm) processDiv() {
	//let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
	//        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);
	//
	//        let result = operand1 / operand2;
	//
	//        self.acc = unsafe { transmute::<f64, u64>(result) };
	//        self.ip = self.ip + 1;
}

func (vm *vm) processCon() {

}

func (vm *vm) processPush() {
	//check if stack is full
	//        if self.sp <= STACK_MEM_BOTTOM {
	//            self.set_mem_error_flag("Push attempt on full stack.");
	//            return;
	//        }
	//
	//        let reg_index = self.get_register_index();
	//        let mut value: u64;
	//
	//        // register push
	//        if reg_index < REG_NUM {
	//            value = self.reg[reg_index];
	//        // ip push
	//        } else if reg_index == IP_INDEX {
	//            value = self.ip as u64;
	//        // sp push
	//        } else if reg_index == SP_INDEX {
	//            value = self.sp as u64;
	//        // acc push
	//        } else if reg_index == ACC_INDEX {
	//            value = self.acc;
	//        // error
	//        } else {
	//            self.set_mem_error_flag("Push attempt on unknown register.");
	//            return;
	//        }
	//
	//        self.mem[self.sp] = value;
	//        self.sp = self.sp - 1;
}

func (vm *vm) processPop() {
	//check if stack is empty
	//        if self.sp > STACK_MEM_TOP {
	//            self.set_mem_error_flag("Pop attempt on empty stack.");
	//            return;
	//        }
	//
	//        let reg_index = self.get_register_index();
	//        let value = self.mem[self.sp];
	//        self.sp = self.sp + 1;
	//
	//        // register pop
	//        if reg_index < REG_NUM {
	//            self.reg[reg_index] = value;
	//        // ip pop
	//        } else if reg_index == IP_INDEX {
	//            self.ip = value as usize;
	//        // sp pop
	//        } else if reg_index == SP_INDEX {
	//            self.sp = value as usize;
	//        // acc pop
	//        } else if reg_index == ACC_INDEX {
	//            self.acc = value;
	//        // error
	//        } else {
	//            self.set_mem_error_flag("Pop attempt on unknown register.");
	//            return;
	//        }
}

func (vm *vm) processLd() {
	//let reg_index = self.get_register_index();
}

func (vm *vm) processSt() {
	//let reg_index = self.get_register_index();
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

func (vm *vm) getOperand(rotation uint32) float64 {
	//        let inst = self.mem[self.ip];
	//        let operand_reg = (inst.rotate_left(rotation) & REG_MASK) as usize;
	//
	//        unsafe { transmute::<u64, f64>(self.reg[operand_reg]) }

	return 0.0
}

func (vm *vm) getRegisterIndex() uint {
	//        let inst = self.mem[self.ip];
	//        let reg_index = (inst.rotate_left(FULL_REG_ROTATION) & FULL_REG_MASK) as usize;
	//
	//        reg_index

	return 0
}
