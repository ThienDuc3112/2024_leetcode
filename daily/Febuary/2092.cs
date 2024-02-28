namespace FindAllPeople;
public class Solution
{
    // Soltuion is correct
    // TLE due to the UFDS implementation
    public IList<int> FindAllPeople(int n, int[][] meetings, int firstPerson)
    {
        HashSet<int> secretKnower = new()
        {
            0,
            firstPerson
        };
        List<int[]> sm = new(meetings.OrderBy(meeting => meeting[2]));
        int i = 0;
        while (i < sm.Count)
        {
            Dictionary<int, int> ufds = new();
            HashSet<int> parentSecretKnower = new();
            int timeframe = sm[i][2];
            while (i < sm.Count && sm[i][2] == timeframe)
            {
                if (!ufds.ContainsKey(sm[i][0])) ufds.Add(sm[i][0], sm[i][0]);
                if (!ufds.ContainsKey(sm[i][1])) ufds.Add(sm[i][1], sm[i][1]);
                if (secretKnower.Contains(sm[i][0]) || secretKnower.Contains(sm[i][1]))
                {
                    parentSecretKnower.Add(sm[i][0]);
                    parentSecretKnower.Add(sm[i][1]);
                }
                int firstParent = ufds[sm[i][0]];
                while (ufds[firstParent] != firstParent) firstParent = ufds[firstParent];
                int secondParent = ufds[sm[i][1]];
                while (ufds[secondParent] != secondParent) secondParent = ufds[secondParent];
                if (firstParent != secondParent)
                {
                    if (parentSecretKnower.Contains(secondParent)) parentSecretKnower.Add(firstParent);
                    ufds[secondParent] = firstParent;
                }
                i++;
            }
            foreach (var person in ufds)
            {
                int parent = person.Value;
                while (parent != ufds[parent]) parent = ufds[parent];
                if (parentSecretKnower.Contains(parent))
                {
                    secretKnower.Add(person.Key);
                }
            }
        }
        return secretKnower.ToList();
    }
}