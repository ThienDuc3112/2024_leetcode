public class StockSpanner {
    int day;
    Stack<(int price, int day)> monoStack;
    public StockSpanner() {
        day = 0;
        monoStack = new();
    }
    
    public int Next(int price) {
        day++;
        while(monoStack.Any() && monoStack.Peek().price <= price){
            monoStack.Pop();
        }
        int res = !monoStack.Any()? day : day - monoStack.Peek().day;
        monoStack.Push((price, day));
        return res;
    }
}
