namespace IsValid;
public class Solution
{
    public bool IsValid(string s)
    {
        Stack<char> stack = new();
        foreach (char c in s)
        {
            bool success = stack.TryPeek(out char peek);
            switch (c)
            {
                case ')':
                    if (!success || peek != '(') return false;
                    stack.Pop();
                    break;
                case ']':
                    if (!success || peek != '[') return false;
                    stack.Pop();
                    break;
                case '}':
                    if (!success || peek != '{') return false;
                    stack.Pop();
                    break;
                default:
                    stack.Push(c);
                    break;
            }
        }
        return stack.Count == 0;
    }
}