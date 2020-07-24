/*
    Given a directed, acyclic graph of N nodes.
    Find all possible paths from node 0 to node N-1,
    and return them in any order.
    The graph is given as follows:
    the nodes are 0, 1, ..., graph.length - 1.
    graph[i] is a list of all nodes j for which the edge (i, j) exists
*/
class Solution {
public:
    vector<int> path;
    vector<vector<int>> paths;
    
    void findAllPaths(int node, vector<int> *parent) {
        path.push_back(node);

        if(node == 0)
            paths.push_back(path);
        
        for(int i=0; i<parent[node].size(); ++i) {
            findAllPaths(parent[node][i], parent);
            path.pop_back();
        }
    }
    
    vector<vector<int>> allPathsSourceTarget(vector<vector<int>>& graph) {
        int N = graph.size();
        vector<int> parent[N];
        
        
        int count = 0;
        for(auto &adj: graph) {
            for(auto &e: adj)
                parent[e].push_back(count);
            ++count;
        }
        findAllPaths(N-1, parent);
        
        for(int i=0; i<paths.size(); ++i)
            sort(paths[i].begin(), paths[i].end());
        
        return paths;
    }
};
