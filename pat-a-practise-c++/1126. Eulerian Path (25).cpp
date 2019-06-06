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
vector<vector<int> >v;
vector<int>vis;
int N,M,odd=0,mxcnt=0;
void dsf(int st){
	vis[st]=true;
	mxcnt++;
	for(int i=0;i<v[st].size();i++){
		int ed = v[st][i];
		if(vis[ed]==false){
			dsf(ed);
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	cin>>N>>M;
	v.resize(N+1);
	vis.resize(N+1,0);
	for(int i=0;i<M;i++){
		int a,b;
		cin>>a>>b;
		v[a].push_back(b);
		v[b].push_back(a);
	}
	int mxindex=0,mxsize=0;
	for(int i=1;i<v.size();i++){
		cout<<v[i].size();
		if(v[i].size()%2!=0){
			odd++;
		}
		if(i!=v.size()-1){
			cout<<" ";
		}
	}
	cout<<""<<endl;
	dsf(1);
		if(mxcnt==N&&odd==0)
		cout<<"Eulerian"<<endl;
		else if(mxcnt==N&&odd==2)
		cout<<"Semi-Eulerian"<<endl;
		else
		cout<<"Non-Eulerian"<<endl;
	return 0;
}


