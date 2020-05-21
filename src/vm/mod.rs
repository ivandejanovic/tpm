// The MIT License (MIT)
//
// Copyright (c) 2016-2017 Ivan Dejanovic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

use std::io::stdin;
use std::collections::HashMap;
use std::mem::transmute;

use common::INSTRUCTIONS;

// Bit Masks
const INST_MASK: u64 = 0xFF;
const REG_MASK: u64 = 0xF;
const FULL_REG_MASK: u64 = 0x1F;

//Rotation values
const INST_ROTATION: u32 = 8;
const OPERAND_ONE_ROTATION: u32 = 12;
const OPERANT_TWO_ROTATION: u32 = 15;
const FULL_REG_ROTATION: u32 = 13;

//Register index
const IP_INDEX: usize = 8;
const SP_INDEX: usize = 9;
const ACC_INDEX: usize = 10;

// Opcodes
const HALT: u64 = 0x00;
const IN_N: u64 = 0x01;
const OUT_N: u64 = 0x02;
const IN_S: u64 = 0x03;
const OUT_S: u64 = 0x04;
const ADD: u64 = 0x05;
const SUB: u64 = 0x06;
const MUL: u64 = 0x07;
const DIV: u64 = 0x08;
const CON: u64 = 0x09;
const PUSH: u64 = 0x0A;
const POP: u64 = 0x0B;
const LD: u64 = 0x0C;
const ST: u64 = 0x0D;
const JMP: u64 = 0x0E;
const JGR: u64 = 0x0F;
const JGE: u64 = 0x10;
const JEQ: u64 = 0x11;
const JNE: u64 = 0x12;
const JLE: u64 = 0x13;
const JLS: u64 = 0x14;

const REG_NUM: usize = 8;
const MEM_CAP: usize = 50000;
const VAR_MEM_TOP: usize = 50000;
const VAR_MEM_BOTTOM: usize = 40001;
const STACK_MEM_TOP: usize = 40000;
const STACK_MEM_BOTTOM: usize = 30001;

pub struct Vm {
    ip: usize,
    sp: usize,
    acc: u64,
    reg: [u64; REG_NUM],
    mem: [u64; MEM_CAP],
}

impl Vm {
    pub fn new() -> Box<Vm> {
        Box::new(Vm {
            ip: 0,
            sp: STACK_MEM_TOP,
            acc: 0,
            reg: [0; REG_NUM],
            mem: [0; MEM_CAP],
        })
    }

    pub fn load_code(&mut self, code: Vec<String>) {
        println!("{}", code.len());

        let mut opcode_map: HashMap<&'static str, u64> = HashMap::new();

        opcode_map.insert(INSTRUCTIONS.halt, HALT);
        opcode_map.insert(INSTRUCTIONS.in_n, IN_N);
        opcode_map.insert(INSTRUCTIONS.out_n, OUT_N);
        opcode_map.insert(INSTRUCTIONS.in_s, IN_S);
        opcode_map.insert(INSTRUCTIONS.out_s, OUT_S);
        opcode_map.insert(INSTRUCTIONS.add, ADD);
        opcode_map.insert(INSTRUCTIONS.sub, SUB);
        opcode_map.insert(INSTRUCTIONS.mul, MUL);
        opcode_map.insert(INSTRUCTIONS.div, DIV);
        opcode_map.insert(INSTRUCTIONS.con, CON);
        opcode_map.insert(INSTRUCTIONS.push, PUSH);
        opcode_map.insert(INSTRUCTIONS.pop, POP);
        opcode_map.insert(INSTRUCTIONS.ld, LD);
        opcode_map.insert(INSTRUCTIONS.st, ST);
        opcode_map.insert(INSTRUCTIONS.jmp, JMP);
        opcode_map.insert(INSTRUCTIONS.jgr, JGR);
        opcode_map.insert(INSTRUCTIONS.jge, JGE);
        opcode_map.insert(INSTRUCTIONS.jeq, JEQ);
        opcode_map.insert(INSTRUCTIONS.jne, JNE);
        opcode_map.insert(INSTRUCTIONS.jle, JLE);
        opcode_map.insert(INSTRUCTIONS.jls, JLS);

        for instruction in code {
            println!("{}", instruction);
        }
    }

    pub fn execute(&mut self) {
        loop {
            let inst = self.mem[self.ip];
            let opcode = inst.rotate_left(INST_ROTATION) & INST_MASK;

            match opcode {
                HALT => break,
                IN_N => self.process_in_n(),
                OUT_N => self.process_out_n(),
                IN_S => self.process_in_s(),
                OUT_S => self.process_out_s(),
                ADD => self.process_add(),
                SUB => self.process_sub(),
                MUL => self.process_mul(),
                DIV => self.process_div(),
                CON => self.process_con(),
                PUSH => self.process_push(),
                POP => self.process_pop(),
                LD => self.process_ld(),
                ST => self.process_st(),
                JMP => self.process_jmp(),
                JGR => self.process_jgr(),
                JGE => self.process_jge(),
                JEQ => self.process_jeq(),
                JNE => self.process_jne(),
                JLE => self.process_jle(),
                JLS => self.process_jls(),
                _ => {
                    println!("Unknow instruction. Aborting.");
                    break;
                }

            }

            if self.has_errors() {
                self.throw_exception();
                break;
            }
        }
    }

    fn process_in_n(&mut self) {
        let mut input_text = String::new();
        stdin().read_line(&mut input_text).expect(
            "failed to read from stdin",
        );

        let trimmed = input_text.trim();
        match trimmed.parse::<f64>() {
            Ok(i) => {
                self.acc = unsafe { transmute::<f64, u64>(i) };
                self.ip = self.ip + 1;
            }
            Err(..) => self.set_io_error_flag(),
        };
    }

    fn process_out_n(&mut self) {
        let output = unsafe { transmute::<u64, f64>(self.acc) };
        println!("{}", output);
        self.ip = self.ip + 1;
    }

    fn process_in_s(&mut self) {}

    fn process_out_s(&mut self) {}

    fn process_add(&mut self) {
        let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);

        let result = operand1 + operand2;

        self.acc = unsafe { transmute::<f64, u64>(result) };
        self.ip = self.ip + 1;
    }

    fn process_sub(&mut self) {
        let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);

        let result = operand1 - operand2;

        self.acc = unsafe { transmute::<f64, u64>(result) };
        self.ip = self.ip + 1;
    }

    fn process_mul(&mut self) {
        let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);

        let result = operand1 * operand2;

        self.acc = unsafe { transmute::<f64, u64>(result) };
        self.ip = self.ip + 1;
    }

    fn process_div(&mut self) {
        let operand1 = self.get_operand(OPERAND_ONE_ROTATION);
        let operand2 = self.get_operand(OPERANT_TWO_ROTATION);

        let result = operand1 / operand2;

        self.acc = unsafe { transmute::<f64, u64>(result) };
        self.ip = self.ip + 1;
    }

    fn process_con(&mut self) {}

    fn process_push(&mut self) {
        //check if stack is full
        if self.sp <= STACK_MEM_BOTTOM {
            self.set_mem_error_flag("Push attempt on full stack.");
            return;
        }

        let reg_index = self.get_register_index();
        let mut value: u64;

        // register push
        if reg_index < REG_NUM {
            value = self.reg[reg_index];
        // ip push
        } else if reg_index == IP_INDEX {
            value = self.ip as u64;
        // sp push
        } else if reg_index == SP_INDEX {
            value = self.sp as u64;
        // acc push
        } else if reg_index == ACC_INDEX {
            value = self.acc;
        // error
        } else {
            self.set_mem_error_flag("Push attempt on unknown register.");
            return;
        }

        self.mem[self.sp] = value;
        self.sp = self.sp - 1;
    }

    fn process_pop(&mut self) {
        //check if stack is empty
        if self.sp > STACK_MEM_TOP {
            self.set_mem_error_flag("Pop attempt on empty stack.");
            return;
        }

        let reg_index = self.get_register_index();
        let value = self.mem[self.sp];
        self.sp = self.sp + 1;

        // register pop
        if reg_index < REG_NUM {
            self.reg[reg_index] = value;
        // ip pop
        } else if reg_index == IP_INDEX {
            self.ip = value as usize;
        // sp pop
        } else if reg_index == SP_INDEX {
            self.sp = value as usize;
        // acc pop
        } else if reg_index == ACC_INDEX {
            self.acc = value;
        // error
        } else {
            self.set_mem_error_flag("Pop attempt on unknown register.");
            return;
        }
    }

    fn process_ld(&mut self) {
        let reg_index = self.get_register_index();
    }

    fn process_st(&mut self) {
        let reg_index = self.get_register_index();
    }

    fn process_jmp(&mut self) {}

    fn process_jgr(&mut self) {}

    fn process_jge(&mut self) {}

    fn process_jeq(&mut self) {}

    fn process_jne(&mut self) {}

    fn process_jle(&mut self) {}

    fn process_jls(&mut self) {}

    fn set_io_error_flag(&mut self) {
        println!("IO error flag set!")
    }

    fn set_mem_error_flag(&mut self, error: &'static str) {
        println!("{}", error)
    }

    fn has_errors(&mut self) -> bool {
        false
    }

    fn throw_exception(&mut self) {
        println!("Exception thrown!")
    }

    fn get_operand(&mut self, rotation: u32) -> f64 {
        let inst = self.mem[self.ip];
        let operand_reg = (inst.rotate_left(rotation) & REG_MASK) as usize;

        unsafe { transmute::<u64, f64>(self.reg[operand_reg]) }
    }

    fn get_register_index(&mut self) -> usize {
        let inst = self.mem[self.ip];
        let reg_index = (inst.rotate_left(FULL_REG_ROTATION) & FULL_REG_MASK) as usize;

        reg_index
    }
}
