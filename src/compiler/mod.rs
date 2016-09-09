use common::fetch_msg1;
use std::fs::File;
use std::io::*;

pub fn cpl(mut codefile: File) -> Vec<String> {
    let mut code: String = String::new();
    let result = codefile.read_to_string(&mut code);

    match result {
        Ok(_) => compile(code),
        Err(e) => panic!("Error reading souce file: {}", e),
    }
}

fn compile(code: String) -> Vec<String> {
    println!("{}", code);
    println!("{}", fetch_msg1());
    Vec::new()
}
