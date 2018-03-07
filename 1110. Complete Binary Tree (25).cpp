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
struct Info{
	int left=-1,right=-1;
};
vector<Info>tree;
vector<vector<int> >v;
void dsf(int root,int cnt){
	v[cnt].push_back(root);
	if(root!=-1){
		dsf(tree[root].left,cnt+1);
		dsf(tree[root].right,cnt+1);
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N;
	cin>>N;
	tree.resize(N);
	vector<int>t;
	for(int i=0;i<=N;i++){
		v.push_back(t);
	}
	t.resize(N,0);
	for(int i=0;i<N;i++){
		char l[10],r[10];
		cin>>l>>r;
		int left=-1,right=-1;
		if(isdigit(l[0])){
			left=atoi(l);
			t[left]=1;
		}
		if(isdigit(r[0])){
			right=atoi(r);
			t[right]=1;
		}
		Info info;
		info.left=left;
		info.right=right;
		tree[i]=info;
	}
	int last=-1,index=-1;
	for(int i=0;i<t.size();i++){
		if(t[i]==0){
			index=i;
			dsf(i,0);
			break;
		}
	}
	bool flag=true;
	for(int i=v.size()-1;i>=0;i--){
		for(int j=v[i].size()-1;j>=0;j--){
			if(last==-1){
				last=v[i][j];
			}else if(v[i][j]==-1){
				flag=false;
			}
//			printf("%d ",v[i][j]);
		}
//		printf("\n");
	}
	if(flag){
		cout<<"YES "<<last<<endl;
	}else{
		cout<<"NO "<<index<<endl;
	}
	return 0;
}


