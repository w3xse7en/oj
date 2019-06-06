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
vector<vector<int> >v,family;
void dsf(int fa,int lv){
	family[lv].push_back(fa);
	if(v[fa].size()){
		for(int i=0;i<v[fa].size();i++){
			dsf(v[fa][i],lv+1);
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,M;
	cin>>N>>M;
	v.resize(101);
	family.resize(N);
	for(int i=0;i<M;i++){
		int fa,n,child;
		cin>>fa>>n;
		for(int j=0;j<n;j++){
			cin>>child;
			v[fa].push_back(child);
		}
	}
	dsf(1,0);
	int maxsize=0,lv=0;
	for(int i=0;i<family.size();i++){
		if(maxsize<family[i].size()){
			maxsize=family[i].size();
			lv=i+1;
		}
//		for(int j=0;j<family[i].size();j++){
//			printf("%d ",family[i][j]);
//		}
//		printf("\n");
	}
	cout<<maxsize<<" "<<lv<<endl;
	return 0;
}


