public class MinStack
{
    Stack<(int val, int minSoFar)> stack;
    public MinStack()
    {
        stack = new();
    }

    public void Push(int val)
    {
        if (stack.Any())
        {
            var peek = stack.Peek();
            stack.Push((val, Math.Min(val, peek.minSoFar)));
        }
        else
        {
            stack.Push((val, val));
        }
    }

    public void Pop()
    {
        stack.Pop();
    }

    public int Top()
    {
        return stack.Peek().val;
    }

    public int GetMin()
    {
        return stack.Peek().minSoFar;
    }
}