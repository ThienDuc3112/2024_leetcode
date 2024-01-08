import java.util.HashSet;
import java.util.List;

class Solution {
    HashSet<Integer> visited;

    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        visited = new HashSet<Integer>();
        visit(rooms, 0);
        return visited.size() >= rooms.size();
    }

    void visit(List<List<Integer>> rooms, int room) {
        visited.add(room);
        for (int key : rooms.get(room)) {
            if (!visited.contains(key)) {
                visit(rooms, key);
            }
        }
    }
}