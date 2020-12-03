package main

import (
	"fmt"
    "log"
	"os/exec"
	"math"
)

const SUM_TO_STOLE = 1000000;


func main() {
	out, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%s", out)
	MODULE := math.Pow(2,32)
	fmt.Printf("%s", MODULE)
}
	