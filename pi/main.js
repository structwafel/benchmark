"use strict";

let fs = require("fs");
let rounds = parseInt(fs.readFileSync("./rounds.txt", "utf8"));

let x = 1.0;
let pi = 1.0;
let stop = rounds + 2

for (let i = 2; i < stop; i++) {
    x = -x;
    pi += x / (2 * i - 1);
}

pi *= 4;

console.log(pi);