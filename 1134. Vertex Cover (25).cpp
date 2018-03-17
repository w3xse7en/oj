#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
#include <list>
#include <set>
using namespace std;
int main(){
//	freopen("in.txt","r",stdin);
	int N,M;
	cin>>N>>M;
	vector<int>ori(N,-1),vis;
	vector<vector<int> >e(N);
	for(int i=0;i<M;i++){
		int a,b;
		cin>>a>>b;
		e[a].push_back(b);
		e[b].push_back(a);
	}
	int Nv;
	cin>>Nv;
	for(int i=0;i<Nv;i++){
		int K;cin>>K;
		vis.clear();
		vis.resize(N,0);
		int cnt=0;
		for(int j=0;j<K;j++){
			int t;cin>>t;
				vis[t]=1;
			for(int k=0;k<e[t].size();k++){
				if(vis[e[t][k]]==0){
//					cout<<"!"<<t<<" "<<e[t][k]<<endl;
					cnt++;
				}
			}
		}
		if(cnt==M){
			cout<<"Yes"<<endl;
		}else{
			cout<<"No"<<endl;
		}
	}
	return 0;
}


