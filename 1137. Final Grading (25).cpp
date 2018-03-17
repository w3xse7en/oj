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
struct Info{
	string id;
	int Gp=-1,Gm=-1,Gf=-1,G=-1;
};
bool cmp(Info &a,Info &b){
	if(a.G==b.G){
		return a.id<b.id;
	}else{
		return a.G>b.G;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int P,M,N;
	cin>>P>>M>>N;
	vector<Info>v;
	map<string,Info>mp;
	for(int i=0;i<P;i++){
		string id;int score;
		cin>>id>>score;
		mp[id].Gp=score;
		mp[id].id=id;
	}
	for(int i=0;i<M;i++){
		string id;int score;
		cin>>id>>score;
		mp[id].Gm=score;
		mp[id].id=id;
	}
	for(int i=0;i<N;i++){
		string id;int score;
		cin>>id>>score;
		mp[id].Gf=score;
		mp[id].id=id;
	}
	map<string,Info>::iterator it;
	for(it=mp.begin();it!=mp.end();it++){
		Info info = it->second;
		if(info.Gp>=200){
			if(info.Gm>info.Gf){
				info.G=round(info.Gm*0.4+info.Gf*0.6);
			}else{
				info.G=info.Gf;
			}
			if(info.G>=60){
				v.push_back(info);
			}
		}
//		printf("%s %d %d %d %d\n",info.id.c_str(),info.Gp,info.Gm,info.Gf,info.G);
	}
	sort(v.begin(),v.end(),cmp);
	for(int i=0;i<v.size();i++){
		Info info = v[i];
		printf("%s %d %d %d %d\n",info.id.c_str(),info.Gp,info.Gm,info.Gf,info.G);
		
	}
	return 0;
}


