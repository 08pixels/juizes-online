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

    int traversal(TreeNode* root, int guess, int sum) {
        if ( !root )
            return 0;
        
        int count = (guess + root->val) == sum;
        
        int l = traversal(root->left, guess + root->val, sum);
        int r = traversal(root->right,guess + root->val, sum);
        
        return count + l + r;
    }
    
    int pathSum(TreeNode* root, int sum) {
        if ( !root )
            return 0;
        
        int curr = traversal(root, 0, sum);
        int l = pathSum(root->left,  sum);
        int r = pathSum(root->right, sum);
        
        return curr + l + r;
    }
};
