package main

import (
    "math"
    "strings"
)

var base36 = map[int]string{
    0: "0",
    1: "1",
    2: "2",
    3: "3",
    4: "4",
    5: "5",
    6: "6",
    7: "7",
    8: "8",
    9: "9",
    10: "A",
    11: "B",
    12: "C",
    13: "D",
    14: "E",
    15: "F",
    16: "G",
    17: "H",
    18: "I",
    19: "J",
    20: "K",
    21: "L",
    22: "M",
    23: "N",
    24: "O",
    25: "P",
    26: "Q",
    27: "R",
    28: "S",
    29: "T",
    30: "U",
    31: "V",
    32: "W",
    33: "X",
    34: "Y",
    35: "Z",
}

func Converter(num float64, precision int, base float64) string {
    if num == 0 {
        return "0." + strings.Repeat("0", precision)
    }
    result := ""
    isNegative := false
    if num < 0 {
        num *= -1
        isNegative = true
    }
    k := int(math.Floor(math.Log(num) / math.Log(base))) + 1

    for i := k - 1; i > -precision-1; i-- {
        if len(result) == k {
            result += "."
        }
        digit := int(math.Mod(math.Floor((num / math.Pow(base, float64(i)))), base))
        num -= float64(digit) * math.Pow(base, float64(i))
        result += base36[digit]
    }

    if isNegative {
        return "-" + result
    }
    return result
}
