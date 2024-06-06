package main

import (
    "strconv"
    "strings"
)


func addBinary(a string, b string) string {
    var sb strings.Builder
    carry := 0
    i, j := len(a)-1, len(b)-1

    for i >= 0 || j >= 0 || carry == 1 {
        if i >= 0 {
            carry += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            carry += int(b[j] - '0')
            j--
        }
        sb.WriteString(strconv.Itoa(carry % 2))
        carry /= 2
    }

    result := sb.String()
    return reverse(result)
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}