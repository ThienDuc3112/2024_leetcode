package sortedSquares
class Solution {
    fun sortedSquares(nums: IntArray): IntArray {
        val squared: MutableList<Int> = nums.map { num -> num * num }.toMutableList()
        squared.sort()
        return squared.toIntArray()
    }
}
