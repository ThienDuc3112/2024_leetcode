namespace JobScheduling;
public class Solution
{
    (int end, int profit)[] memo;
    int MaxProfit(List<(int start, int end, int profit)> jobs)
    {
        for (int i = 0; i < jobs.Count; i++)
        {
            int profit = jobs[i].profit;
            int lastJob = BinarySearch(jobs, i);
            if (lastJob != -1) profit += memo[lastJob].profit;
            if (i != 0) profit = Math.Max(profit, memo[i - 1].profit);
            memo[i] = (jobs[i].end, profit);
        }
        return memo[memo.Length - 1].profit;
    }
    public int JobScheduling(int[] startTime, int[] endTime, int[] profit)
    {
        memo = new (int, int)[startTime.Length];
        List<(int start, int end, int profit)> jobs = new();
        for (int i = 0; i < startTime.Length; i++)
        {
            jobs.Add((startTime[i], endTime[i], profit[i]));
        }
        jobs.Sort((a, b) => a.end.CompareTo(b.end));
        return MaxProfit(jobs);
    }
    // The only function that I copied
    // Cuz my own function keep not working?
    int BinarySearch(List<(int start, int end, int profit)> jobs, int index)
    {
        int start = 0;
        int end = index - 1;
        while (start <= end)
        {
            int mid = (start + end) / 2;
            if (jobs[mid].end <= jobs[index].start)
            {
                if (jobs[mid + 1].end <= jobs[index].start)
                    start = mid + 1;
                else
                    return mid;
            }
            else
                end = mid - 1;
        }
        return -1;
    }
}