package numbers

import "math"

//As tested on leetcode

func myPow(x float64, n int) float64 {
    answer := 0.0
    if x == 0 || math.Abs(x) == 1{
        if math.Mod(float64(n), 2) == 0{
            return math.Abs(x)
        }
        return x
    }

    absoluteN := n
    if n == 0 {
        answer = 1
    }else if n > 0{
        answer = x
        absoluteN = n
    }else if n < 0 {
        answer = 1/x
        absoluteN = -n
    }
    
    for i := 1; i < absoluteN; i++{
        if n > 1{
            answer = answer * x  
        }
        if n < 1 {
            answer = answer * ( 1/x )
        }
    }
    
    return answer
}
