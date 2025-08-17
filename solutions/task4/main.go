package main

import "fmt"

func firstUniqChar(s string) int {
    result := -1
    numsSet := make(map[rune]int)
    
    for _, val := range s{
        _, ok := numsSet[val]
        if ok {
            numsSet[val]++
            continue
        }
        numsSet[val] = 1
    }
    
    for i, val := range s{
        if numsSet[val] == 1{
            result = i
            break
        }
    } 
    return result
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
