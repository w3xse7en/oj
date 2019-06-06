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
	int id;
	vector<int>memb;
	double Asets=0.0,Aarea=0.0;
};
bool cmp(Info &a,Info &b){
	if(a.Aarea==b.Aarea){
		return a.id<b.id;
	}else{
		return a.Aarea>b.Aarea;
	}
}
vector<Info>v;
vector<int>ids(10001,0);
void add(int id,int index,vector<int>& memb,set<int> &relate){
	if(id!=-1){
		if(ids[id]==0){
			memb.push_back(id);
		}else{
			relate.insert(ids[id]);
		}
		ids[id]=index;
	}
}
int main(){
	freopen("in.txt","r",stdin); 
	int N;
	cin>>N;
	v.resize(N+1);
	for(int i=1;i<=N;i++){
		vector<int>memb;
		set<int>relate;
		int id,fa,mo,nchild,estate,area;
		cin>>id>>fa>>mo>>nchild;
		add(id,i,memb,relate);add(fa,i,memb,relate);add(mo,i,memb,relate);
		for(int j=0;j<nchild;j++){
			int tchild;cin>>tchild;
			add(tchild,i,memb,relate);
		}
		cin>>estate>>area;
		v[i].Aarea=area;
		v[i].Asets=estate;
		v[i].memb=memb;
		set<int>::iterator it;
		for(it=relate.begin();it!=relate.end();it++){
			vector<int> f=v[*it].memb;
			for(int j=0;j<f.size();j++){
				ids[f[j]]=i;
				v[i].memb.push_back(f[j]);
			}
			v[i].Aarea+=v[*it].Aarea;
			v[i].Asets+=v[*it].Asets;
			v[*it].memb.clear();
			v[*it].Aarea=0.0;
			v[*it].Asets=0.0;
		}
	}
	int cnt=0;
	for(int i=0;i<v.size();i++){
		sort(v[i].memb.begin(),v[i].memb.end());
		int num=v[i].memb.size();
		if(num){
			cnt++;
			v[i].id=v[i].memb[0];
			v[i].Asets/=num;
			v[i].Aarea/=num;
		}
	}
	sort(v.begin(),v.end(),cmp);
	cout<<cnt<<endl;
	for(int i=0;i<cnt;i++){
		int num=v[i].memb.size();
		if(num){
			printf("%04d ",v[i].id);
			printf("%d %.3f %.3f\n",num,v[i].Asets,v[i].Aarea);
		}
	}
	return 0;
}


