/*
    There are a total of n courses you have to take, labeled from 0 to n-1.
    Some courses may have prerequisites, for example to take course 0
    you have to first take course 1, which is expressed as a pair: [0,1]
    Given the total number of courses and a list of prerequisite pairs,
    return the ordering of courses you should take to finish all courses.
    There may be multiple correct orders, you just need to return one of them.
    If it is impossible to finish all courses, return an empty array.
*/

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
