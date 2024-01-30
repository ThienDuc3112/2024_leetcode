public class MyQueue
{
    Stack<int> stack1;
    Stack<int> stack2;
    public MyQueue()
    {
        stack1 = new();
        stack2 = new();
    }

    public void Push(int x)
    {
        stack1.Push(x);
    }

    public int Pop()
    {
        // while (stack1.Any())
        // {
        //     stack2.Push(stack1.Pop());
        // }
        // int res = stack2.Pop();
        // while (stack2.Any())
        // {
        //     stack1.Push(stack2.Pop());
        // }
        // return res;

        // Amortized O(1)
        if (!stack2.Any())
        {
            while (stack1.Any())
            {
                stack2.Push(stack1.Pop());
            }
        }
        return stack2.Pop();
    }

    public int Peek()
    {
        // while (stack1.Any())
        // {
        //     stack2.Push(stack1.Pop());
        // }
        // int res = stack2.Peek();
        // while (stack2.Any())
        // {
        //     stack1.Push(stack2.Pop());
        // }
        // return res;

        // Amortized O(1)
        if (!stack2.Any())
        {
            while (stack1.Any())
            {
                stack2.Push(stack1.Pop());
            }
        }
        return stack2.Peek();
    }

    public bool Empty()
    {
        return !stack1.Any() && !stack2.Any();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue obj = new MyQueue();
 * obj.Push(x);
 * int param_2 = obj.Pop();
 * int param_3 = obj.Peek();
 * bool param_4 = obj.Empty();
 */