package main

import (
    "math"
    "strings"
)

var base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
        result += string(base36[digit])
    }

    if isNegative {
        return "-" + result
    }
    return result
}
