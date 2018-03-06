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
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	int bird[10001]={0},maxb=0;
	vector<vector<int> >tree(N+1);
	vector<int>connect;
	for(int i=0;i<N;i++){
		int K;cin>>K;
		connect.clear();
		for(int j=0;j<K;j++){
			int t;cin>>t;
			if(maxb<t){
				maxb=t;
			}
			if(bird[t]!=0){
				connect.push_back(bird[t]);
			}
			tree[i+1].push_back(t);
			bird[t]=i+1;
		}
		for(int j=1;j<=maxb;j++){
			for(int k=0;k<connect.size();k++){
				if(bird[j]==connect[k]){
					bird[j]=i+1;
				}
			}
		}
	}
	set<int> s;
	for(int i=1;i<=maxb;i++){
		s.insert(bird[i]);
//		printf("%d %d\n",i,bird[i]);
	}
	printf("%d %d\n",s.size(),maxb);
	cin>>N;
	for(int i=0;i<N;i++){
		int a,b;
		cin>>a>>b;
		if(bird[a]!=bird[b]){
			printf("No\n");
		}else{
			printf("Yes\n");
		}
	}
	return 0;
}


