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
int N,M,K;
vector<vector<int> >g;
vector<int> vis;
int main(){
//	freopen("in.txt","r",stdin);
	cin>>N>>M;
	vis.resize(N+1,0);
	for(int i=0;i<=N;i++){
		g.push_back(vis);
	}
	for(int i=0;i<M;i++){
		int a,b;cin>>a>>b;
		g[a][b]=g[b][a]=1;
	}
	cin>>K;
	for(int i=0;i<K;i++){
		int n,st,ed;
		vis.clear();
		vis.resize(N,0);
		cin>>n>>st;
		bool flag=true;
		for(int j=0;j<n-1;j++){
			cin>>ed;
			if(g[st][ed]==1&&vis[ed-1]==0){
				vis[ed-1]=1;
			}else{
				flag=false;
			}
			st=ed;
		}
		for(int j=0;j<vis.size();j++){
			if(vis[j]==0){
				flag=false;
				break;
			}
//			printf("%d ",vis[j]);
		}
//		printf("\n");
		if(flag){
			cout<<"YES"<<endl;
		}else{
			cout<<"NO"<<endl;
		}
	}
}


