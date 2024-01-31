import java.util.Stack;

// Read solution
// Look up monotonic stack more
class Solution {
    public int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> monoStack = new Stack<>();
        int[] res = new int[temperatures.length];

        for (int i = temperatures.length - 1; i >= 0; i--) {
            while (!monoStack.empty() && temperatures[i] >= temperatures[monoStack.peek()]) {
                monoStack.pop();
            }
            if (!monoStack.empty()) {
                res[i] = monoStack.peek() - i;
            }
            monoStack.push(i);
        }
        return res;
    }
}