/**
*   Design a HashSet without using any built-in hash table libraries.
*/

class MyHashSet {
public:
    int table[31260];
    
    MyHashSet() {
        for(int i=0; i<31260; ++i)
            table[i] = 0;
    }
    
    void add(int key) {
        int hash      = key >> 5;
        int collision = key  & 31;
        
        table[hash] |= (1 << collision);
    }
    
    void remove(int key) {
        int hash      = key >> 5;
        int collision = key  & 31;
        
        if(contains(key))
            table[hash] ^= (1 << collision);
    }
    
    /** Returns true if this set contains the specified element */
    bool contains(int key) {
        return table[key >> 5] & (1 << (key & 31));
    }
    
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet* obj = new MyHashSet();
 * obj->add(key);
 * obj->remove(key);
 * bool param_3 = obj->contains(key);
 */
