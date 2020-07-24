class Solution {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int> dependencies(numCourses, 0);
        vector<int> graph[numCourses];
        
        for(auto& task: prerequisites) {
            ++dependencies[task[0]];
            graph[task[1]].push_back(task[0]);
        }
        
        vector<int> tasks;
        
        for(int i=0; i<numCourses; ++i) {
            if(!dependencies[i])
                tasks.push_back(i);
        }
        
        
        int task = 0;
        while(task < tasks.size()) {
            int current = tasks[task];
            
            for(auto &e: graph[current]) {
                --dependencies[e];
                
                if(dependencies[e] == 0) {
                    tasks.push_back(e);
                }
            }
            ++task;
        }
        
        if(task == numCourses)
            return tasks;
        else
            return {};
        
    }
};
