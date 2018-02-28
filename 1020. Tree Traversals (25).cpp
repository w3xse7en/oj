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
vector<vector<int> >tree;
void proot(vector<int> iord,vector<int> pord,int cnt){
	int pt=0,is = iord.size(),ps=pord.size();
//	for(int i=0;i<is;i++){
//		printf("c:%d %d %d\n",cnt,iord[i],pord[i]);
//	}
	if(ps!=0){
		pt = pord[ps-1];
		tree[cnt].push_back(pt);
	}
	for(int i=0;i<is;i++){
		if(pt==iord[i]){
			vector<int>li,lp,ri,rp;
			li.insert(li.end(),iord.begin(),iord.begin()+i);
			ri.insert(ri.end(),iord.begin()+i+1,iord.end());
			lp.insert(lp.end(),pord.begin(),pord.begin()+i);
			rp.insert(rp.end(),pord.begin()+i,pord.end()-1);
			proot(li,lp,cnt+1);
			proot(ri,rp,cnt+1);
			
		}
	}
}
int main() {
//	freopen("in.txt","r",stdin);
	int n;
	cin>>n;
	vector<int> iord,pord;
	for(int i=0;i<n;i++){
		vector<int>v;
		tree.push_back(v);
		int t;cin>>t;
		pord.push_back(t);
	}
	for(int i=0;i<n;i++){
		int t;cin>>t;
		iord.push_back(t);
	}
	proot(iord,pord,0);
	bool first =true;
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
