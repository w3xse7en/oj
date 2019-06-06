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
	int l,r;
};
vector<vector<int> >tree;
vector<Info>v;
void dsf(int root,int cnt){
	if(root!=-1){
		tree[cnt].push_back(root);
		dsf(v[root].l,cnt+1);
		dsf(v[root].r,cnt+1);
	}
}
bool fst=true;
void dsft(int root){
	if(root!=-1){
		dsft(v[root].l);
		if(fst){
			cout<<root;
			fst=false;
		}else{
			cout<<" "<<root;
		}
		dsft(v[root].r);
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N;cin>>N;
	tree.resize(N);
	v.resize(N);
	vector<int> rt(N,0);
	for(int i=0;i<N;i++){
		char a,b;
		int l=-1,r=-1;
		cin>>a>>b;
		if(isdigit(a)){
			l=a-'0';
			rt[l]=1;
		}
		if(isdigit(b)){
			r=b-'0';
			rt[r]=1;
		}
		Info info;
		info.l=l;info.r=r;
		v[i]=info;
		
	}
	int root;
	for(int i=0;i<rt.size();i++){
		if(rt[i]==0){
			root=i;
			dsf(i,0);
		}
	}
	bool first=true;
	for(int i=0;i<tree.size();i++){
		for(int j=tree[i].size()-1;j>=0;j--){
			if(first){
				printf("%d",tree[i][j]);
				first=false;
			}else{
				printf(" %d",tree[i][j]);
			}
		}
	}
	for(int i=0;i<N;i++){
		int t=v[i].l;
		v[i].l=v[i].r;
		v[i].r=t;
	}
	printf("\n");
	dsft(root);
	return 0;
}


