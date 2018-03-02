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
	int id=0,cnt=-1;
};
bool cmp(Info &a,Info &b){
	if(a.cnt==b.cnt){
		return a.id<b.id;
	}else{
		return a.cnt>b.cnt;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,K;
	cin>>N>>K;
	vector<int> vn(N+1);
	vector<Info>v(K);
	
	for(int i=0;i<N;i++){
		int t;
		cin>>t;
		if(i!=0){
			cout<<t<<":";
			for(int j=0;j<v.size();j++){
				if(v[j].cnt!=-1)
				cout<<" "<<v[j].id;
			}
			cout<<""<<endl;
		}
		int cnt=vn[t];
		if(cnt>=v[K-1].cnt){
			bool flag=true;
			for(int j=0;j<K;j++){
				if(t==v[j].id){
					v[j].cnt=cnt;
					flag = false;
				}
			}
			if(flag){
				Info info;
				info.id=t;
				info.cnt=cnt;
				v.push_back(info);
				sort(v.begin(),v.end(),cmp);
				v.pop_back();
			}
		}
		sort(v.begin(),v.end(),cmp);
		vn[t]++;
	}
	return 0;
}


