class Solution {
    fun isUgly(n: Int): Boolean {
        if( n <= 0) return false
        var a = n
        while(a % 2 == 0) a /= 2
        while(a % 3 == 0) a /= 3
        while(a % 5 == 0) a /= 5
        return a == 1
    }
}