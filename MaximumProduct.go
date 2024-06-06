package main

func maxProduct(words []string) int {
    if len(words) == 0 {
        return 0
    }
    
    max := 0

    for i := 0; i < len(words)-1; i++ {
        for j := i+1; j < len(words); j++ {
            if !hasCommonLetters(words[i], words[j]) {
                temp := len(words[i]) * len(words[j])
                if temp > max {
                    max = temp
                }
            }
        }
    }
    return max
}

func hasCommonLetters(word1, word2 string) bool {
    letters := make(map[rune]bool)
    
    for _, char := range word1 {
        letters[char] = true
    }
    
    for _, char := range word2 {
        if letters[char] {
            return true
        }
    }
    
    return false
}
