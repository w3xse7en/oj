#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include<cmath>
using namespace std;
vector<vector<int> >road;
vector<bool>visit;
	int n,m,k;
void dfs(int node) {
    visit[node] = true;
    for(int i = 1; i <= n; i++) {
        if(visit[i] == false && road[node][i] == 1)
            dfs(i);
    }
}
int main() {
	//freopen("in.txt","r",stdin);
	cin>>n>>m>>k;
	visit.resize(n+1,false);
	road.resize(n+1);
	for(int i=0;i<=n;i++){
		road[i].resize(n+1,0);
	}
	for(int i=0;i<m;i++){
		int a,b;
		cin>>a>>b;
		road[a][b]=1;
		road[b][a]=1;
	} 
	for(int i=0;i<k;i++){
		for(int j=0;j<visit.size();j++){
			visit[j]=false;
		}
		int city,cnt=0;
		cin>>city;
        visit[city] = true;
        for(int j = 1; j <= n; j++) {
            if(visit[j] == false) {
                dfs(j);
                cnt++;
            }
        }
		cout<<cnt-1<<endl;
	}
	
	return 0;
}
