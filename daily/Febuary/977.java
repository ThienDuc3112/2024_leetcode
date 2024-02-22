import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

class Solution {
    private class Person {
        public int person;
        public Set<Integer> trust;
        public Set<Integer> trustBy;

        public Person(int id) {
            person = id;
            trust = new HashSet<>();
            trustBy = new HashSet<>();
        }
    }

    public int findJudge(int n, int[][] trust) {
        if (n == 1)
            return 1;
        HashMap<Integer, Person> map = new HashMap<>();
        for (int[] relation : trust) {
            if (!map.containsKey(relation[0]))
                map.put(relation[0], new Person(relation[0]));
            if (!map.containsKey(relation[1]))
                map.put(relation[1], new Person(relation[1]));
            map.get(relation[0]).trust.add(relation[1]);
            map.get(relation[1]).trustBy.add(relation[0]);
        }
        for (int i = 1; i <= n; ++i) {
            if (!map.containsKey(i))
                return -1;
            Person curPerson = map.get(i);
            if (curPerson.trust.size() == 0 && curPerson.trustBy.size() == n - 1)
                return curPerson.person;
        }
        return -1;
    }

    public int betterFindJudge(int n, int[][] trust) {
        int[][] trustMatrix = new int[n + 1][];
        for (int i = 0; i < trustMatrix.length; i++) {
            trustMatrix[i] = new int[n + 1];
        }
        for (int[] relation : trust) {
            trustMatrix[relation[0]][relation[1]] = 1;
        }
        for (int i = 0; i < trustMatrix.length; i++) {
            boolean trustNoone = true;
            boolean trustByEveryone = true;
            for (int j = 0; j < trustMatrix.length; j++) {
                if (i == j)
                    continue;
                if (trustMatrix[j][i] != 1)
                    trustByEveryone = false;
                if (trustMatrix[i][j] == 1)
                    trustNoone = false;
            }
            if (trustNoone && trustByEveryone)
                return i;
        }
        return -1;
    }

    public int evenBetterFindJudge(int n, int[][] trust) {
        int[] count = new int[n + 1];
        for (int[] relation : trust) {
            count[relation[0]] = -1;
            if (count[relation[1]] == -1)
                continue;
            count[relation[1]]++;
        }
        for (int i = 0; i < count.length; i++) {
            if (count[i] == n - 1)
                return i;
        }
        return -1;
    }
}