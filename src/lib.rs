mod common;
mod compiler;
mod vm;

use std::env::args;
use std::env::Args;
use std::fs::File;

use self::compiler::cpl;
use self::vm::vm;

const MINUS: &'static str = "-";
const VERSION_FLAG_SHORT: &'static str = "-v";
const VERSION_FLAG_LONG: &'static str = "--version";
const HELP_FLAG_SHORT: &'static str = "-h";
const HELP_FLAG_LONG: &'static str = "--help";

const HELP_MESSAGE: &'static str = "Usage: tpm <codefilename>\n\nOptions:\n  -h, --help       \
                                    Prints help\n  -v, --version    Prints version";
const VERSION_MESSAGE: &'static str = "TPM interpreter version 0.1.0";
const INVALID_USAGE_MESSAGE: &'static str = "Invalid usage. For correct usage examples please \
                                             try: tpm -h";

pub fn tpm() {
    if !handle_flags(args()) {
        return;
    }

    let mut arguments = args();

    if arguments.len() < 2 {
        invalid_usage_msg();
        return;
    }

    match arguments.nth(1) {
        Some(codefile) => {
            let filename: &str = codefile.as_ref();
            let result = File::open(filename);
            match result {
                Ok(file) => {
                    let code = cpl(file);
                    vm(code);
                }
                Err(e) => println!("{}", e),
            }

        }
        None => invalid_usage_msg(),
    }

}

fn handle_flags(args: Args) -> bool {
    for argument in args {
        let arg: &str = argument.as_ref();

        if arg.starts_with(MINUS) {
            match arg {
                VERSION_FLAG_SHORT |
                VERSION_FLAG_LONG => println!("{}", VERSION_MESSAGE),
                HELP_FLAG_SHORT | HELP_FLAG_LONG => println!("{}", HELP_MESSAGE),
                _ => invalid_usage_msg(),
            }

            return false;
        }
    }

    true
}

fn invalid_usage_msg() {
    println!("{}", INVALID_USAGE_MESSAGE)
}
