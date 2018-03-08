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
int hob[1001]={0};
bool cmp(int &a,int &b){
	return a>b;
}
int main() {
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	vector<int>v(N+1,0),cnt(N+1,1),result;
	for(int i=1;i<=N;i++){
		int M;scanf("%d: ",&M);set<int> same;
		for(int j=0;j<M;j++){
			int t;cin>>t;
			if(hob[t]!=0){
//				printf("%d %d %d\n",t,hob[t],i);
				same.insert(hob[t]);
			}
			hob[t]=i;
		}
		set<int>::iterator it;
		for(it=same.begin();it!=same.end();it++){
			cnt[i]+=cnt[*it];
		}
		for(int j=0;j<1001;j++){
			for(it=same.begin();it!=same.end();it++){
				if(hob[j]==*it){
					hob[j]=i;
				}
			}
		}
	}
	for(int i=0;i<1001;i++){
		if(hob[i]!=0){
			v[hob[i]]++;
//			printf("%d ---- %d\n",i,hob[i]);
		}
	}
	for(int i=0;i<v.size();i++){
		if(v[i]!=0){
//			cout<<i<<":"<<v[i]<<endl;
//			cout<<i<<"--"<<cnt[i]<<endl;
			result.push_back(cnt[i]);
		}
	}
	sort(result.begin(),result.end(),cmp);
	cout<<result.size()<<endl;
	for(int i=0;i<result.size();i++){
			cout<<result[i];
		if(i!=result.size()-1){
			cout<<" ";
		}
	}
    return 0;
}
