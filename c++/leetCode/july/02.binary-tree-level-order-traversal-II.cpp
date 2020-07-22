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
    
    void traverse(TreeNode *root, int depth, vector<vector<int>> &levelOrder) {
        if(root == NULL)
            return;
        
        if(levelOrder.size() <= depth) {
            vector<int> v;
            levelOrder.push_back(v);
        }
        
        levelOrder[depth].push_back(root->val);
        
        traverse(root->left, 1 + depth, levelOrder);
        traverse(root->right, 1 + depth, levelOrder);
        
    }
    
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
            
        vector<vector<int>> levelOrder;
        traverse(root, 0, levelOrder);
        
        reverse(levelOrder.begin(), levelOrder.end());
        return levelOrder;
    }
};
