package kata

type Pointer struct {
  x int
  y int
}

func Interpreter(code string, iterations, width, height int) string {
  grid := GridInit(width, height)
  pointer := Pointer{
    x: 0,
    y: 0,
  }

  code = Sanitize(code)
  jmpMap := MakeJumpMap(code)
  
  var iterCnt int
  
  for i := 0; i < len(code); i++ {
    if iterCnt >= iterations {
      break 
    }
    
    switch code[i] {
      case 'n':
        if pointer.y == 0 {
          pointer.y = height - 1
        } else {
          pointer.y--
        }
      case 'e':
        if pointer.x == width - 1 {
          pointer.x = 0
        } else {
          pointer.x++
        }
      case 's':
        if pointer.y == height - 1 {
          pointer.y = 0  
        } else {
          pointer.y++
        }
      case 'w':
        if pointer.x == 0 {
          pointer.x = width - 1
        } else {
          pointer.x--
        }
      case '*':
        grid[pointer.y][pointer.x] = !grid[pointer.y][pointer.x]
      case '[':
        if !grid[pointer.y][pointer.x] {
          i = jmpMap[i]
        }
      case ']':
        if grid[pointer.y][pointer.x] {
          i = jmpMap[i]
        }
    }
    
    iterCnt++
  }
  
  return GridToString(grid)
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

// Sanitize code
func Sanitize(code string) (sCode string) {
  for i := range code {
    switch code[i] {
      case 'n','e','s','w','*','[',']':
        sCode += string(code[i])
    }
  }
  
  return
}

// GridInit helper
func GridInit(width, height int) [][]bool {
  // grid init
  var grid [][]bool
  for i := 0; i < height; i++ {
    grid = append(grid, make([]bool, width))
  }
  
  return grid
}

// GridToString helper to return string
func GridToString(grid [][]bool) (res string) {
  for i := range grid {
    for j := range grid[i] {
      if grid[i][j] {
        res += "1"
      } else {
        res += "0"
      }
    }

    if i != len(grid) - 1 {
      res += "\r\n"
    }
  }
  
  return
}
