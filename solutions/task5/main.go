package main

import "fmt"

func isAnagram(s string, t string) bool {
    sLetters := make(map[rune]int)
    tLetters := make(map[rune]int)
    
    for _, val := range s{
        _, ok := sLetters[val]
        if ok {
            sLetters[val]++
            continue
        }
        sLetters[val] = 1
    }
    
    for _, val := range t{
        _, ok := tLetters[val]
        if ok {
            tLetters[val]++
            continue
        }
        tLetters[val] = 1
    }
    
    if len(sLetters) != len(tLetters){
        return false
    }
    
    for key, val := range sLetters{
        _, ok := tLetters[key]
        if !ok || tLetters[key] != val {
            return false
        }
    }
    return true
}

func main(){
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}