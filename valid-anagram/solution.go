func isAnagram(s string, t string) bool {
    sMap := map[rune]int{}
    tMap := map[rune]int{}

    if len(s) != len(t) {
        return false
    }

    for i := 0; i < len(s); i++ {
        if number, ok := sMap[rune(s[i])]; ok {
            sMap[rune(s[i])] = number + 1;
        } else {
            sMap[rune(s[i])] = 1;
        }

        if number, ok := tMap[rune(t[i])]; ok {
            tMap[rune(t[i])] = number + 1;
        } else {
            tMap[rune(t[i])] = 1;
        }
    }

    for key, value := range sMap {
        number, ok := tMap[key]
        if !ok {
            return false
        }

        if number != value {
            return false
        } 
    }

    return true
}
