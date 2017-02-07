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

use std::collections::HashMap;

use common::INSTRUCTIONS;

// Opcodes
const HALT: i64 = 0x00;
const IN_N: i64 = 0x01;
const OUT_N: i64 = 0x02;
const IN_S: i64 = 0x03;
const OUT_S: i64 = 0x04;
const ADD: i64 = 0x05;
const SUB: i64 = 0x06;
const MUL: i64 = 0x07;
const DIV: i64 = 0x08;
const CON: i64 = 0x09;
const PUSH: i64 = 0x0A;
const POP: i64 = 0x0B;
const LD: i64 = 0x0C;
const ST: i64 = 0x0D;
const JMP: i64 = 0x0E;
const JGR: i64 = 0x0F;
const JGE: i64 = 0x10;
const JEQ: i64 = 0x11;
const JNE: i64 = 0x12;
const JLE: i64 = 0x13;
const JLS: i64 = 0x14;

const REG_NUM: usize = 8;
const INS_MEM_CAP: usize = 50000;

pub struct Vm {
    pc: usize,
    sp: usize,
    acc: i64,
    reg: [i64; REG_NUM],
    ins_mem: [i64; INS_MEM_CAP],
}

impl Vm {
    pub fn new() -> Box<Vm> {
        Box::new(Vm {
            pc: 0,
            sp: 0,
            acc: 0,
            reg: [0; REG_NUM],
            ins_mem: [0; INS_MEM_CAP],
        })
    }

    pub fn load_code(&self, code: Vec<String>) {
        println!("{}", code.len());

        let mut opcode_map: HashMap<&'static str, i64> = HashMap::new();

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

    pub fn execute(&self) {
        loop {
            let inst = self.ins_mem[self.pc];

            match inst {
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
        }
    }

    fn process_in_n(&self) {}

    fn process_out_n(&self) {}

    fn process_in_s(&self) {}

    fn process_out_s(&self) {}

    fn process_add(&self) {}

    fn process_sub(&self) {}

    fn process_mul(&self) {}

    fn process_div(&self) {}

    fn process_con(&self) {}

    fn process_push(&self) {}

    fn process_pop(&self) {}

    fn process_ld(&self) {}

    fn process_st(&self) {}

    fn process_jmp(&self) {}

    fn process_jgr(&self) {}

    fn process_jge(&self) {}

    fn process_jeq(&self) {}

    fn process_jne(&self) {}

    fn process_jle(&self) {}

    fn process_jls(&self) {}
}
