/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void traverse(TreeNode *root, int depth, vector<vector<int>> &zigzag) {
        if(root == NULL)
            return;
        
        if(zigzag.size() <= depth) {
            vector<int> v;
            zigzag.push_back(v);
        }
        
        if(depth&1)
            zigzag[depth].insert(zigzag[depth].begin(), root->val);
        else
            zigzag[depth].push_back(root->val);
        
        
        traverse(root->left,  depth + 1, zigzag);
        traverse(root->right, depth + 1, zigzag);
    }
    
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> zigzagOrder;
        traverse(root, 0, zigzagOrder);
        
        return zigzagOrder;
    }
};
