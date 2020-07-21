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
    ListNode* removeElements(ListNode* head, int val) {
        
        while(head != NULL && head->val == val) {
            head = head->next;
        }
        
        ListNode* aux = head;
        ListNode* temp = head;
        
        while(aux != NULL) {
            if(aux->val == val) {
                temp->next = aux->next;
                aux = temp;
            }
            
            temp = aux;
            aux = aux->next;
        }
        
        return head;
    }
};
