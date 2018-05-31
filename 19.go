package main

import "fmt"

func main() {
    year := 1900
    result := 0
    daysToEndWeek := 0
    
    daysCount := [12]int{31,28,31,30,31,30,31,31,30,31,30,31}
    
    for year != 2001 {
        delimiter := 4
        
        if year % 100 == 0 {
            delimiter = 400
        }
        
        if year % delimiter == 0 {
            daysCount[1] = 29
        } else {
            daysCount[1] = 28
        }
        
        for i := 0; i < len(daysCount); i++ {
            daysToEndWeek = (daysToEndWeek + daysCount[i]) % 7
            
            if (year != 1900 && daysToEndWeek + 1 == 7) {
                result++;
            }
        }
        
        year++
    }
    
    fmt.Println(result)
}