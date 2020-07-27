/**
 *      Given inorder and postorder traversal of a tree, construct the binary tree.
 *      Note:
 *      You may assume that duplicates do not exist in the tree.
 */


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
    int curr;
    TreeNode* recover(int i, int j, vector<int>& inorder, vector<int>& postorder) {
        if ( i > j )
            return NULL;
        
        int ind = find( inorder.begin() + i, inorder.end(), postorder[--curr] ) - inorder.begin();
        TreeNode *node = new TreeNode(inorder[ind]);
        
        if ( i ^ j ) {
            node->right = recover( ind + 1, j, inorder, postorder );
            node->left  = recover( i, ind - 1, inorder, postorder );
        }
        
        return node;
    }
    
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        curr = inorder.size();
        return recover( 0, inorder.size()-1, inorder, postorder );
    }
};