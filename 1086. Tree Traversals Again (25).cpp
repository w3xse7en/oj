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
#include <list>
#include <queue>
using namespace std;
struct Node{
	int l=-1,r=-1;
};
vector<Node> tree;
void buildtree(vector<int> pre,vector<int> in){
	if(pre.size()){
		int root=pre[0];
		int index=-1;
		vector<int>lpre,lin,rpre,rin;
		int p=1;
		for(int i=0;i<in.size();i++){
			if(root==in[i]){
				index=i;
				continue;
			}
			if(index==-1){
				lpre.push_back(pre[p++]);
				lin.push_back(in[i]);
			}else{
				rpre.push_back(pre[p++]);
				rin.push_back(in[i]);
			}
		}
		if(lpre.size()){
			tree[root].l=lpre[0];
		}
		if(rpre.size()){
			tree[root].r=rpre[0];
		}
		buildtree(lpre,lin);
		buildtree(rpre,rin);
//		for(int i=0;i<lpre.size();i++){
//			cout<<lpre[i]<<"--l--"<<lin[i]<<endl;
//		}
//		cout<<""<<endl;
//		for(int i=0;i<rpre.size();i++){
//			cout<<rpre[i]<<"--r--"<<rin[i]<<endl;
//		}
	}
}
bool first=true;
void dsf(int root){
	if(root!=-1){
		dsf(tree[root].l);
		dsf(tree[root].r);
		if(first){
			cout<<root;
			first=false;
		}else{
			cout<<" "<<root;
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	tree.resize(N+1);
	vector<int>q;
	vector<int> pre,in;
	for(int i=0;i<N*2;i++){
		string t;
		cin>>t;
		if(t.compare("Push")==0){
			int n;cin>>n;
			q.push_back(n);
			pre.push_back(n);
		}else{
			in.push_back(q[q.size()-1]);
			q.pop_back();
		}
	}
	buildtree(pre,in);
	dsf(pre[0]);
//	for(int i=0;i<tree.size();i++){
//		cout<<i<<endl;
//		cout<<tree[i].l<<"  "<<tree[i].r<<endl;
//	}
	return 0;
}


