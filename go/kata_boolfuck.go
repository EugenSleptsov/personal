package kata

func Boolfuck(code, input string) string {
  code = Sanitize(code)
  jmpMap := MakeJumpMap(code)
  inputBits := InputBits(input)
  
  outputBits := []bool{}

  tape := map[int]bool{0:false}
  
  pointer := 0
  for i := 0; i < len(code); i++ {
    curBit := tape[pointer]

    switch code[i] {
      case '+':
        tape[pointer] = !tape[pointer]
      case ',':
        newBit := false
        if len(inputBits) > 0 {
          newBit = inputBits[0]
          inputBits = inputBits[1:]
        }
        tape[pointer] = newBit
      case ';':
        outputBits = append(outputBits, curBit)
      case '<':
        pointer--
        if _,ok := tape[pointer]; !ok {
          tape[pointer] = false
        }
      case '>':
        pointer++
        if _,ok := tape[pointer]; !ok {
          tape[pointer] = false
        }
      case '[':
        if !curBit {
          i = jmpMap[i]
        }
      case ']':
        if curBit {
          i = jmpMap[i]
        }
    }
  }

  return OutputString(outputBits)
}

func InputBits(input string) (bits []bool) {
  for i := range input {
    bits = append(bits, ByteToBits(input[i])...)
  }
  
  return
}

func OutputString(output []bool) string { 
  var bytes []byte
  for len(output) > 8 {
    bytes = append(bytes, BitsToByte(output[:8]))
    output = output[8:]
  }

  bytes = append(bytes, BitsToByte(output))
  return string(bytes)
}

func ByteToBits(b byte) []bool {
  raw := []bool{
    b >= 128,
    b % 128 >= 64,
    b % 64 >= 32,
    b % 32 >= 16,
    b % 16 >= 8,
    b % 8 >= 4,
    b % 4 >= 2,
    b % 2 == 1,
  }

  return []bool{
    raw[7],
    raw[6],
    raw[5],
    raw[4],
    raw[3],
    raw[2],
    raw[1],
    raw[0],
  }
}

func BitsToByte(bits []bool) byte {
  b := 0
  d := 1
  for _,bit := range bits {
    if bit {
      b += d
    }
    d *= 2
  }
  
  return byte(b)
}

func Sanitize(code string) (sCode string) {
  for i := range code {
    switch code[i] {
      case '+',',',';','<','>','[',']':
        sCode += string(code[i])
    }
  } 
  return
}

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
