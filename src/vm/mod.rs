pub struct Vm {

}

impl Vm {
    pub fn new() -> Box<Vm> {
        Box::new(Vm {})
    }

    pub fn load_code(&self, code: Vec<String>) {
        println!("{}", code.len());
    }

    pub fn execute(&self) {}
}
