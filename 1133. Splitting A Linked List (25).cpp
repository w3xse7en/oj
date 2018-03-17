#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
#include <list>
#include <set>
using namespace std;
struct Info{
	int a=-1,data,b=-1;
};
vector<Info> v(100001);
vector<Info> l;
int rt,N,K;
void dsf(int rt,int mode){
	if(mode==-1){
		if(v[rt].data<0){
			l.push_back(v[rt]);
		}
	}
	if(mode==0){
		if(v[rt].data<=K&&v[rt].data>=0){
			l.push_back(v[rt]);
		}
	}
	if(mode==1){
		if(v[rt].data>K){
			l.push_back(v[rt]);
		}
	}
	if(v[rt].b!=-1){
		dsf(v[rt].b,mode);
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	cin>>rt>>N>>K;
	for(int i=0;i<N;i++){
		int a,b,data;
		cin>>a>>data>>b;
		v[a].data=data;
		v[a].a=a;
		v[a].b=b;
	}
	dsf(rt,-1);
	dsf(rt,0);
	dsf(rt,1);
	for(int i=0;i<l.size();i++){
		if(i+1<l.size()){
			printf("%05d %d %05d\n",l[i].a,l[i].data,l[i+1].a);
		}
	}
	printf("%05d %d -1\n",l[l.size()-1].a,l[l.size()-1].data);
	return 0;
}


