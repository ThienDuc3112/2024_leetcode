import java.util.HashSet;

class RandomizedSet {
    private HashSet<Integer> set;

    public RandomizedSet() {
        set = new HashSet<>();
    }

    public boolean insert(int val) {
        boolean res = set.contains(val);
        if (res)
            return false;
        set.add(val);
        return true;
    }

    public boolean remove(int val) {
        boolean res = set.contains(val);
        if (!res)
            return false;
        set.remove(val);
        return true;
    }

    public int getRandom() {
        int random = (int) (Math.random() * set.size());
        var iterator = set.iterator();
        for (int i = 0; i < random - 1; i++) {
            iterator.next();
        }
        return iterator.next();
    }
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * RandomizedSet obj = new RandomizedSet();
 * boolean param_1 = obj.insert(val);
 * boolean param_2 = obj.remove(val);
 * int param_3 = obj.getRandom();
 */