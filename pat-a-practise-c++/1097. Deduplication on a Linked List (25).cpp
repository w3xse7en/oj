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
	int root,data,next;
};
vector<Info>v(100001),root,dup;
vector<int>cnt(10001,0),ord;
void dsf(int node){
	if(node!=-1){
		int d=abs(v[node].data);
		if(cnt[d]!=0){
			dup.push_back(v[node]);
		}else if(cnt[d]==0){
			root.push_back(v[node]);
		}
		cnt[d]++;
		dsf(v[node].next);
	}
}
void pf(vector<Info> root){
	for(int i=0;i<root.size();i++){
		if(i+1!=root.size()){
			printf("%05d %d %05d\n",root[i].root,root[i].data,root[i+1].root);
		}else{
			printf("%05d %d -1\n",root[i].root,root[i].data);
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int rt,N;
	cin>>rt>>N;
	for(int i=0;i<N;i++){
		int a,data,b;
		cin>>a>>data>>b;
		Info info;
		info.root=a;
		info.data=data;
		info.next=b;
		v[a]=info;
	}
	dsf(rt);
	pf(root);
	pf(dup);
	return 0;
}


