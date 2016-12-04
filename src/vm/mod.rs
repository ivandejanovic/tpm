// The MIT License (MIT)
//
// Copyright (c) 2016 Ivan Dejanovic
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

const OPCODE: Opcode = Opcode {
    halt: 0x00,
    in_n: 0x01,
    out_n: 0x02,
    in_s: 0x03,
    out_s: 0x04,
    add: 0x05,
    sub: 0x06,
    mul: 0x07,
    div: 0x08,
    con: 0x09,
    push: 0x0A,
    pop: 0x0B,
    ld: 0x0C,
    st: 0x0D,
    jmp: 0x0E,
    jgr: 0x0F,
    jge: 0x10,
    jeq: 0x11,
    jne: 0x12,
    jle: 0x13,
    jls: 0x14,
};

const reg_num: usize = 8;
const ins_mem_cap: usize = 50000;

pub struct Vm {
    pc: i64,
    sp: i64,
    acc: i64,
    reg: [i64; reg_num],
    ins_mem: [i64; ins_mem_cap],
}

struct Opcode {
    pub halt: i64,
    pub in_n: i64,
    pub out_n: i64,
    pub in_s: i64,
    pub out_s: i64,
    pub add: i64,
    pub sub: i64,
    pub mul: i64,
    pub div: i64,
    pub con: i64,
    pub push: i64,
    pub pop: i64,
    pub ld: i64,
    pub st: i64,
    pub jmp: i64,
    pub jgr: i64,
    pub jge: i64,
    pub jeq: i64,
    pub jne: i64,
    pub jle: i64,
    pub jls: i64,
}

impl Vm {
    pub fn new() -> Box<Vm> {
        Box::new(Vm {
            pc: 0,
            sp: 0,
            acc: 0,
            reg: [0; reg_num],
            ins_mem: [0; ins_mem_cap],
        })
    }

    pub fn load_code(&self, code: Vec<String>) {
        println!("{}", code.len());

        let mut opcode_map: HashMap<&'static str, i64> = HashMap::new();

        opcode_map.insert(INSTRUCTIONS.halt, OPCODE.halt);
        opcode_map.insert(INSTRUCTIONS.in_n, OPCODE.in_n);
        opcode_map.insert(INSTRUCTIONS.out_n, OPCODE.out_n);
        opcode_map.insert(INSTRUCTIONS.in_s, OPCODE.in_s);
        opcode_map.insert(INSTRUCTIONS.out_s, OPCODE.out_s);
        opcode_map.insert(INSTRUCTIONS.add, OPCODE.add);
        opcode_map.insert(INSTRUCTIONS.sub, OPCODE.sub);
        opcode_map.insert(INSTRUCTIONS.mul, OPCODE.mul);
        opcode_map.insert(INSTRUCTIONS.div, OPCODE.div);
        opcode_map.insert(INSTRUCTIONS.con, OPCODE.con);
        opcode_map.insert(INSTRUCTIONS.push, OPCODE.push);
        opcode_map.insert(INSTRUCTIONS.pop, OPCODE.pop);
        opcode_map.insert(INSTRUCTIONS.ld, OPCODE.ld);
        opcode_map.insert(INSTRUCTIONS.st, OPCODE.st);
        opcode_map.insert(INSTRUCTIONS.jmp, OPCODE.jmp);
        opcode_map.insert(INSTRUCTIONS.jgr, OPCODE.jgr);
        opcode_map.insert(INSTRUCTIONS.jge, OPCODE.jge);
        opcode_map.insert(INSTRUCTIONS.jeq, OPCODE.jeq);
        opcode_map.insert(INSTRUCTIONS.jne, OPCODE.jne);
        opcode_map.insert(INSTRUCTIONS.jle, OPCODE.jle);
        opcode_map.insert(INSTRUCTIONS.jls, OPCODE.jls);

        for instruction in code {
            println!("{}", instruction);
        }
    }

    pub fn execute(&self) {}
}
