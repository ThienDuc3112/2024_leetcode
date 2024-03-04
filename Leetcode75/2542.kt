import java.util.PriorityQueue
package maxScore;
// Check from answer
// Ain't no way I would think of this
// Especially in Kotlin
class Solution {
    fun maxScore(nums1: IntArray, nums2: IntArray, k: Int): Long {

        val ind = nums1.indices.sortedWith(compareByDescending<Int> { nums2[it] }.thenByDescending { nums1[it] })
        var score = 0L
        var sum = 0L
        val pq = PriorityQueue<Int>(compareBy{nums1[it]})
        ind.forEach{
            sum += nums1[it].toLong()
            pq.add(it)
            if(pq.size > k) sum -= nums1[pq.poll()].toLong()
            if(pq.size == k) score = maxOf(score, sum * nums2[it].toLong())
        }
        return score
    }
}