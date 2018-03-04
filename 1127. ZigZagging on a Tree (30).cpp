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
vector<vector<int> >tree;
void buildtree(vector<int> por,vector<int> ior,int cnt){
//	for(int i=0;i<ior.size();i++){
//		printf("%d ",ior[i]);
//	}
//	printf("\n");
//	for(int i=0;i<por.size();i++){
//		printf("%d ",por[i]);
//	}
//	printf("\n");
	if(por.size()){
	int rt=por[por.size()-1],leftlen=0,rightlen=0;
	tree[cnt].push_back(rt);
	for(int i=0;i<ior.size();i++){
		if(rt==ior[i]){
			leftlen=i,rightlen=ior.size()-i-1;
			vector<int> leftior,leftpor,rightior,rightpor;
			if(leftlen){
				leftior.insert(leftior.end(),ior.begin(),ior.begin()+i);
				leftpor.insert(leftpor.end(),por.begin(),por.begin()+i);
				buildtree(leftpor,leftior,cnt+1);
			}
			if(rightlen){
				rightior.insert(rightior.end(),ior.begin()+i+1,ior.end());
				rightpor.insert(rightpor.end(),por.begin()+i,por.end()-1);
				buildtree(rightpor,rightior,cnt+1);
			}
		}
	}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	tree.resize(N);
	vector<int> por,ior;
	for(int i=0;i<N;i++){
		int t;cin>>t;
		ior.push_back(t);
	}
	for(int i=0;i<N;i++){
		int t;cin>>t;
		por.push_back(t);
	}
	buildtree(por,ior,0);
//	cout<<"----------------"<<endl;
	bool flag=false,first=true;
	for(int i=0;i<tree.size();i++){
		if(flag){
			for(int j=0;j<tree[i].size();j++){
				printf(" %d",tree[i][j]);
			}
			flag=false;
		}else{
			for(int j=tree[i].size()-1;j>=0;j--){
				if(first){
					printf("%d",tree[i][j]);		
					first=false;
				}else{
					printf(" %d",tree[i][j]);
				}
			}
			flag=true;
		}
//		printf("\n");
	}
	return 0;
}


