public class Trie
{
    private Node head;
    public Trie()
    {
        head = new();
    }

    public void Insert(string word)
    {
        Node cur = head;
        for (int i = 0; i < word.Length; i++)
        {
            if (cur.next[word[i] - 'a'] == null)
            {
                cur.next[word[i] - 'a'] = new();
            }
            cur = cur.next[word[i] - 'a'];
            if (i == word.Length - 1)
            {
                cur.end = true;
            }
        }
    }

    public bool Search(string word)
    {
        Node cur = head;
        for (int i = 0; i < word.Length; i++)
        {
            if (cur.next[word[i] - 'a'] == null)
            {
                return false;
            }
            cur = cur.next[word[i] - 'a'];
        }
        return cur.end;
    }

    public bool StartsWith(string prefix)
    {
        Node cur = head;
        for (int i = 0; i < prefix.Length; i++)
        {
            if (cur.next[prefix[i] - 'a'] == null)
            {
                return false;
            }
            cur = cur.next[prefix[i] - 'a'];
        }
        return true;
    }

    private class Node
    {
        public Node[] next = new Node[26];
        public bool end = false;

        public Node() { }
        public Node(bool end)
        {
            this.end = end;
        }
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.Insert(word);
 * bool param_2 = obj.Search(word);
 * bool param_3 = obj.StartsWith(prefix);
 */