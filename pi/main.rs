// rustc -C opt-level=3 main.rs -o pi_rs
use std::fs::File;
use std::io::prelude::*;

fn main() {
    let mut file = File::open("./rounds.txt").expect("file not found");

    let mut contents = String::new();
    file.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let rounds = contents.trim().parse::<u32>().unwrap() + 2;

    let mut x: f64 = 1.0;
    let mut pi: f64 = 1.0;
    for i in 2..=rounds {
        x = -x;
        pi += x / (2 * i - 1) as f64;
    }

    println!("{:.16}", pi * 4.0);
}
