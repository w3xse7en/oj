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
	string name;
	int tall;
};
bool cmp(Info &a,Info &b){
	if(a.tall==b.tall){
		return a.name>b.name;
	}else{
		return a.tall<b.tall;
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N,K;
	cin>>N>>K;
	vector<Info> v;
	for(int i=0;i<N;i++){
		Info info;
		cin>>info.name>>info.tall;
		v.push_back(info);
	}
	sort(v.begin(),v.end(),cmp);
//	for(int i=0;i<v.size();i++){
//		printf("%s %d\n",v[i].name.c_str(),v[i].tall);
//	}
	int extra = N%K,sz=N/K;
	for(int i=0;i<K;i++){
		int index=(sz+extra)/2,cnt=1,flag=-1;
		vector<Info>t(sz+extra);
		for(int j=0;j<sz+extra;j++){
//			printf("%d %d %d %d \n",index,cnt,flag,sz+extra);
			t[index]=v[v.size()-1];
			v.pop_back();
			index+=(cnt*flag);
			flag*=-1;
			cnt++;
		}
		extra=0;
		for(int j=0;j<t.size();j++){
			cout<<t[j].name;
			if(j!=t.size()-1){
				cout<<" ";
			}
		}
		cout<<""<<endl;
	}
	return 0;
}


