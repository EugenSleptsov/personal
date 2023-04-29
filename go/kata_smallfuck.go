package kata

func Interpreter(code, tape string) string {
  var pointer int
  var jmpStack []int

  tapeByte := []byte(tape)
  for i := 0; i < len(code); i++ {
    r := rune(code[i])

    switch r {
      case '>':
        pointer++
        if pointer == len(tape) {
          return string(tapeByte)
        }
      case '<':
        pointer--
        if pointer < 0 {
          return string(tapeByte)
        }
      case '*':
        if tapeByte[pointer] == byte('0') {
          tapeByte[pointer] = byte('1')
        } else {
          tapeByte[pointer] = byte('0')
        }
      case '[':
        if tapeByte[pointer] == byte('1') {
          jmpStack = append(jmpStack, i)
        } else {
          cnt := 1
          // ignoring all inner [] and searching closing ]
          for j := i + 1; j < len(code); j++ {
            if rune(code[j]) == '[' {
              cnt++
            } else if rune(code[j]) == ']' {
              cnt--
              if cnt == 0 {
                i = j
                break
              }
            }
          }
        }
      case ']':
        if tapeByte[pointer] == byte('1') {
          i = jmpStack[len(jmpStack)-1] - 1
        }
        jmpStack = jmpStack[:len(jmpStack)-1]
    }
  }
  
  return string(tapeByte)
}
