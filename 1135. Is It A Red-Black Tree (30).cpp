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
bool isrbtree=true;
typedef struct node{
	int data;
	bool red=false;
	node* l=NULL;
	node* r=NULL;
}Node;
void build(Node* root,Node* node){
	if(node->data>root->data){
		if(root->r==NULL){
			root->r=node;
		}else{
			build(root->r,node);
		}
	}else{
		if(root->l==NULL){
			root->l=node;
		}else{
			build(root->l,node);
		}
	}
}
int fbcnt=-1;
void dsf(Node* root,int bcnt){
	if(root){
		bool lb=false,rb=false;
		if(root->red){	
			if(root->l){
				lb=root->l->red;
			}
			if(root->r){	
				rb=root->r->red;
			}
		}
		if(lb||rb){
			isrbtree=false;
		}
		int cnt;
		root->red?cnt=0:cnt=1;
		dsf(root->l,bcnt+cnt);
		dsf(root->r,bcnt+cnt);
	}else{
		if(fbcnt==-1){
			fbcnt=bcnt;
		}else{
			if(fbcnt!=bcnt){
				isrbtree=false;
			}
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;cin>>N;
	for(int i=0;i<N;i++){
		int M,rt;cin>>M>>rt;
		Node* root=new Node;
		root->data=rt;
		for(int j=1;j<M;j++){
			int n;cin>>n;
			Node* node = new Node;
			if(n<0){
				node->red=true;
			}
			node->data=abs(n);
			build(root,node);	
		}
		isrbtree=true;
		if(root->data<0){
			cout<<"No"<<endl;
			continue;
		}
		fbcnt=-1;
		dsf(root,0);
		if(isrbtree){
			cout<<"Yes"<<endl;
		}else{
			cout<<"No"<<endl;
		}
	}
	return 0;
}


