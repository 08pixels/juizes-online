/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    vector<ListNode *> nodes;
    
    void traverse(ListNode* head) {
        if ( head == NULL )
            return;

        nodes.push_back(head);
        return traverse(head->next);
    }
    
    void reorderList(ListNode* head) {
        traverse(head);
        
        int n = nodes.size();
        
        for (int i = 0; i < (n >> 1); ++i) {
            nodes[n - i - 1]->next = nodes[i]->next;
            nodes[i]->next = nodes[n - i - 1];
        }
        
        if ( n ^ 0 )
            nodes[n - (n >> 1) - (n & 1)]->next = NULL;
    }
};
