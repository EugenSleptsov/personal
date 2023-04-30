package main

import (
	"bytes"
	"fmt"
)

func MakeJumpMap(code string) map[int]int {
	jmpMap := make(map[int]int)
	var stack []int

	for i := range code {
		switch code[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			j := stack[len(stack)-1]
			jmpMap[j] = i
			jmpMap[i] = j
			stack = stack[:len(stack)-1]
		}
	}

	return jmpMap
}

func Interpreter(code, inputString string) string {
	var pointer int
	var inputPointer int
	tape := make([]byte, 30000)
	jmpMap := MakeJumpMap(code)
	var output bytes.Buffer

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			pointer++
			if pointer >= len(tape) {
				pointer = 0 // Wrap around when reaching the end
			}
		case '<':
			pointer--
			if pointer < 0 {
				pointer = len(tape) - 1 // Wrap around when reaching the beginning
			}
		case '+':
			tape[pointer]++
		case '-':
			tape[pointer]--
		case '.':
			output.WriteByte(tape[pointer])
		case ',':
			if inputPointer < len(inputString) {
				tape[pointer] = inputString[inputPointer]
				inputPointer++
			} else {
				tape[pointer] = 0
			}
		case '[':
			if tape[pointer] == 0 {
				i = jmpMap[i]
			}
		case ']':
			if tape[pointer] != 0 {
				i = jmpMap[i]
			}
		}
	}

	return output.String()
}

func main() {
	code := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."

	fmt.Println(Interpreter(code, "ABC"))
}
