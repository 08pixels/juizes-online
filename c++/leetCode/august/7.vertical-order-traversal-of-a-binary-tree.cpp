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

bool compare(pair<int,int>&a, pair<int,int>&b) {
    if( a.first ^ b.first )
        return a.first > b.first;
    
    return a.second < b.second;
}

class Solution {
public:
    map<int, vector<pair<int,int>>> compute;    

    void solve(int u, int v, TreeNode *root) {
        
        if( root == NULL )
            return ;
        
        compute[u].push_back( {v, root->val} );
        solve(u - 1, v - 1, root->left );
        solve(u + 1, v - 1, root->right);
    }
    
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        solve(0,0, root);
        
        vector<vector<int>> result;
        
        for(auto &k: compute) {
            int n = k.second.size();
            sort(k.second.begin(), k.second.end(), compare);
            vector<int> v(n);
            
            for(int i = 0; i < n; ++i)
                v[i] = k.second[i].second;

            result.push_back(v);
        }
        
        return result;
    }
};
