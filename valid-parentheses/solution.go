func isValid(s string) bool {
    bracketStack := []rune{}
    bracketMap := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range s {
        if char == '(' || char == '[' || char == '{' {
            bracketStack = append(bracketStack, char)
            continue
        }

        if len(bracketStack) < 1 || bracketMap[char] != bracketStack[len(bracketStack) - 1] {
            return false
        }

        bracketStack = bracketStack[:len(bracketStack) - 1]
    }

    return len(bracketStack) == 0
}
