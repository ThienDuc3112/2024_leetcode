namespace SumSubarrayMins;
// Hackathon day + question was hard af
public class Solution
{
    const int MOD = 1000000007;
    public int SumSubarrayMins(int[] arr)
    {
        // validations
        if (arr == null || arr.Length == 0) return 0;

        long sum = 0;
        Stack<int> stack = new Stack<int>();

        // traverse the array
        for (int i = 0; i <= arr.Length; i++)
        {
            // check if stack is > 0
            // if so, check if current value is smaller than top of stack
            while (stack.Count > 0 && arr[stack.Peek()] > (i == arr.Length ? Int32.MinValue : arr[i]))
            {
                // if arr[i] < arr[stack.Peek()], we need to pop 
                // until we find an index where the value was smaller
                // each time we pop, recalc the min for the subarray
                int mid = stack.Pop();

                // left bound is previous stack entry (or beginning of array)
                int left = stack.Count == 0 ? -1 : stack.Peek();

                // right bound is our current index (or end of array)
                int right = i;

                // multiple the min value by number of subarrays that include it
                sum += (long)(arr[mid]) * (mid - left) * (right - mid);
            }
            stack.Push(i);
        }

        return (int)(sum % 1000000007);
    }
}