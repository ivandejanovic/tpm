// The MIT License (MIT)
//
// Copyright (c) 2016-2023 Ivan Dejanovic
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

pub struct Instructions {
    pub halt: &'static str,
    pub in_n: &'static str,
    pub out_n: &'static str,
    pub in_s: &'static str,
    pub out_s: &'static str,
    pub add: &'static str,
    pub sub: &'static str,
    pub mul: &'static str,
    pub div: &'static str,
    pub con: &'static str,
    pub push: &'static str,
    pub pop: &'static str,
    pub ld: &'static str,
    pub st: &'static str,
    pub jmp: &'static str,
    pub jgr: &'static str,
    pub jge: &'static str,
    pub jeq: &'static str,
    pub jne: &'static str,
    pub jle: &'static str,
    pub jls: &'static str,
}

pub const INSTRUCTIONS: Instructions = Instructions {
    halt: "HALT",
    in_n: "IN_N",
    out_n: "OUT_N",
    in_s: "IN_S",
    out_s: "OUT_S",
    add: "ADD",
    sub: "SUB",
    mul: "MUL",
    div: "DIV",
    con: "CON",
    push: "PUSH",
    pop: "POP",
    ld: "LD",
    st: "ST",
    jmp: "JMP",
    jgr: "JGR",
    jge: "JGE",
    jeq: "JEQ",
    jne: "JNE",
    jle: "JLE",
    jls: "JLS",
};
