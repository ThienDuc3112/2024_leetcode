namespace EvalRPN;
public class Solution
{
	public int EvalRPN(string[] tokens)
	{
		Stack<int> stack = new();
		foreach (string token in tokens)
		{
			if (int.TryParse(token, out int cur))
			{
				stack.Push(cur);
			}
			else
			{
				(int num1, int num2) = (stack.Pop(), stack.Pop());
				switch (token)
				{
					case "+":
						stack.Push(num1 + num2);
						break;
					case "-":
						stack.Push(num2 - num1);
						break;
					case "*":
						stack.Push(num1 * num2);
						break;
					case "/":
						stack.Push(num2 / num1);
						break;
				}
			}
		}
		return stack.Peek();
	}
}
