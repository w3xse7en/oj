#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
#include <set>
using namespace std;
struct Info{
	int l=-1,r=-1;
};
vector<Info>v;
vector<vector<int> >tree;
vector<int>node;
int cnt=0;
void dsf(int root,int lv){
	if(root!=-1){
		dsf(v[root].l,lv+1);
		tree[lv].push_back(node[cnt++]);
		dsf(v[root].r,lv+1);
	}
}
int main(){
	freopen("in.txt","r",stdin); 
	int N;
	cin>>N;
	v.resize(N);
	tree.resize(N);
	vector<int>rt(N,0);
	for(int i=0;i<N;i++){
		Info info;
		cin>>info.l>>info.r;
		v[i]=info;
		if(info.l!=-1){
			rt[info.l]=1;
		}
		if(info.r!=-1){
			rt[info.r]=1;
		}
	}
	for(int i=0;i<N;i++){
		int t;cin>>t;
		node.push_back(t);
	}
	sort(node.begin(),node.end());
	for(int i=0;i<N;i++){
		if(rt[i]==0){
			dsf(i,0);
		}
	}
	bool first=true;
	for(int i=0;i<tree.size();i++){
		for(int j=0;j<tree[i].size();j++){
			if(first){
			printf("%d",tree[i][j]);
				first=false;
			}else{
			printf(" %d",tree[i][j]);
			}
		
		}
	}
	return 0;
}


