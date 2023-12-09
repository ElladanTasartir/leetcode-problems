func isPalindrome(s string) bool {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")

    cleanText := strings.ToLower(re.ReplaceAllString(s, ""))

    if cleanText == "" {
        return true
    }

    for i, j := 0, len(cleanText) - 1; i < j; i, j = i+1, j-1 {
        if cleanText[i] != cleanText[j] {
            return false
        }
    }
    
    return true
}
