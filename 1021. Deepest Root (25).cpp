#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include<cmath>
#include<set>
using namespace std;
struct _tree{
vector<int>node;
bool visit=false;
};
vector<_tree>root;
vector<int>deep;
const int MAX=10010;
int n,K=0,hmax=0;
void dfs(int v,int high){
    if(root[v].visit){
        return;
    }
    root[v].visit=true;
    if(hmax<high){
    hmax=high;
    deep.clear();
    deep.push_back(v);
    }
    else if(hmax==high){
        deep.push_back(v);
    }
    for(int i=0;i<root[v].node.size();i++){
        if(!root[root[v].node[i]].visit)
        dfs(root[v].node[i],high+1);
    }
    return ;
}
int main(){
//	freopen("in.txt","r",stdin);
    cin>>n;
    root.resize(n+2);
    for(int i=0;i<n-1;i++){
        int a,b;
        cin>>a>>b;
        root[a].node.push_back(b);
        root[b].node.push_back(a);
    }
    for(int i=1;i<=n;i++){
        if(!root[i].visit){
            dfs(i,0);
            K++;
        }
    }
    if(K!=1){
        printf("Error: %d components",K);
        return 0;
    }
    int t=deep[0];
    set<int>dp;
    for(int i=0;i<deep.size();i++){
    	dp.insert(deep[i]);
	}
    for(int j=0;j<=n;j++)
        root[j].visit=false;
    deep.clear();
	dfs(t,0);
    for(int i=0;i<deep.size();i++){
    	dp.insert((deep[i]));
	}
	set<int>::iterator it;
    for(it=dp.begin();it!=dp.end();it++){
        printf("%d\n",*it);
    }

    return 0;
}

