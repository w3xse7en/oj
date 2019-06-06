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
struct Node{
	char data[20];
	int left,right,sp=0;
	bool vis=true;
};
	vector<Node>v;
void dsf(int rt,int h){
	if(rt==-1){
//		cout<<h;
		return;
	}
	Node node = v[rt];
	if(node.sp)
	cout<<"(";
	dsf(node.left,h+1);
	cout<<node.data;
	dsf(node.right,h+1);
	if(node.sp)
	cout<<")";
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	v.resize(N+1);
	for(int i=1;i<=N;i++){
		Node node=v[i];
		int left,right;
		cin>>node.data>>left>>right;
		node.left=left,node.right=right;
		if(left!=-1){
			v[left].vis=false;
		}
		if(right!=-1){
			v[right].vis=false;
		}
		if(right!=-1||left!=-1){
			node.sp=1;
		}
		v[i]=node;
	}
		for(int j=1;j<=N;j++){
			Node node = v[j];
			if(v[j].vis){
//				printf("%d:%s %d %d %d\n",j,node.data,node.left,node.right,node.vis);
				v[j].sp=0;
				dsf(j,0);
			}
		}
	return 0;
}


