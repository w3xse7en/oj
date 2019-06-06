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
typedef struct _node{
	int data;
	_node* left=NULL;
	_node* right=NULL;
}Node;
vector<vector<int> >v; 
int high=0;
void add(Node* root,Node* n){
	if(n->data<=root->data){
		if(root->left){
			add(root->left,n);
		}else{
			root->left=n;
		}
	}else{
		if(root->right){
			add(root->right,n);
		}else{
			root->right=n;
		}
	}
}
void dsf(Node* root,int cnt){
	if(root){
		if(high<cnt){
			high=cnt;
		}
		v[cnt].push_back(root->data);
		dsf(root->left,cnt+1);
		dsf(root->right,cnt+1);
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N,t;
	cin>>N>>t;
	v.resize(N);
	Node* root =new Node;
	root->data=t;
	for(int i=0;i<N-1;i++){
		cin>>t;
		Node* n=new Node;
		n->data=t;
		add(root,n);
	}
	dsf(root,0);
	int a,b;
	if(high>1){
		a=v[high].size(),b=v[high-1].size();
	}else{
		a=v[high].size(),b=0;
	}
	cout<<a<<" + "<<b<<" = "<<a+b<<endl;
//	for(int i=0;i<v.size();i++){
//		for(int j=0;j<v[i].size();j++){
//			printf("%d ",v[i][j]);
//		}
//		printf("\n");
//	}
	return 0;
}


