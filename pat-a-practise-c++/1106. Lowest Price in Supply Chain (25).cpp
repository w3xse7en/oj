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
vector<vector<int> >v;
int ncnt=0,mincnt;
void dsf(int root,int cnt){
//	printf("%d %d\n",root,cnt);
	if(v[root].size()==0){
		if(cnt<mincnt){
			mincnt=cnt;
			ncnt=1;
		}else if(cnt==mincnt){
			ncnt++;
		}
	}
	for(int i=0;i<v[root].size();i++){
		dsf(v[root][i],cnt+1);
	}
}
int main() {
	freopen("in.txt","r",stdin);
	int N;
	double p,r;
	cin>>N>>p>>r;
	v.resize(N);
	for(int i=0;i<N;i++){
		int M,t;
		cin>>M;
		for(int j=0;j<M;j++){
			cin>>t;
			v[i].push_back(t);
		}
	}
//	for(int i=0;i<N;i++){
//		for(int j=0;j<v[i].size();j++){
//			printf("%d ",v[i][j]);
//		}
//		printf("\n");
//	}
	mincnt=N;
	for(int i=0;i<v[0].size();i++){
		dsf(v[0][i],1);
	}
	double sum=p;
	if(ncnt!=0){
		for(int i=0;i<mincnt;i++){
			sum=sum*(1.0+r/100.0);
		}
	}else{
		ncnt=1;
	}
	printf("%.4f %d",sum,ncnt);
    return 0;
}
