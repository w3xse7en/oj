#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
using namespace std;
bool cmp(int &a,int &b){
	return a>b;
}
struct Info{
	int id,a,c,m,e;
	vector<int> ranks;
};
int main(){
//	freopen("in.txt","r",stdin);
	int n,m;
	cin>>n>>m;
	vector<int> ar,cr,mr,er,ids(1000000,0);
	vector<Info>v;
	Info info;
	for(int i=0;i<n;i++){
		int a,c,m,e,id;
		cin>>id>>c>>m>>e;
		a=(c+m+e)/3;
		ar.push_back(a);
		cr.push_back(c);
		mr.push_back(m);
		er.push_back(e);
		ids[id]=1;
		info.a=a;
		info.c=c;
		info.m=m;
		info.e=e;
		info.id=id;
		v.push_back(info);
	}
	sort(ar.begin(),ar.end(),cmp);
	sort(cr.begin(),cr.end(),cmp);
	sort(mr.begin(),mr.end(),cmp);
	sort(er.begin(),er.end(),cmp);
	for(int i=0;i<v.size();i++){
		for(int j=0;j<ar.size();j++){
			if(ar[j]==v[i].a){
				v[i].ranks.push_back(j+1);
				break;
			}
		}
		for(int j=0;j<cr.size();j++){
			if(cr[j]==v[i].c){
				v[i].ranks.push_back(j+1);
				break;
			}
		}
		for(int j=0;j<mr.size();j++){
			if(mr[j]==v[i].m){
				v[i].ranks.push_back(j+1);
				break;
			}
		}
		for(int j=0;j<er.size();j++){
			if(er[j]==v[i].e){
				v[i].ranks.push_back(j+1);
				break;
			}
		}
	}
//	for(int i=0;i<v.size();i++){
//		printf("id:%d a:%d r:%d c%d r%d m%d r%d e%d r%d\n",
//		v[i].id,v[i].a,v[i].ranks[0],v[i].c,v[i].ranks[1],v[i].m,v[i].ranks[2],v[i].e,v[i].ranks[3]);
//	}
	
	for(int i=0;i<m;i++){
		int id;
		cin>>id;
		if(ids[id]==0){
			printf("N/A\n");
		}else{
			vector<int>ranks;
			for(int i=0;i<v.size();i++){
				if(v[i].id==id){
					ranks=v[i].ranks;
				}
			}
			int min = 10000,t;
			for(int j=0;j<4;j++){
				if(min>ranks[j]){
				   min = ranks[j];
				   t=j;
				}
			}
			if(t==0){
				printf("%d A\n",min);
			}
			if(t==1){
				printf("%d C\n",min);
			}
			if(t==2){
				printf("%d M\n",min);
			}
			if(t==3){
				printf("%d E\n",min);
			}
		}
		
	}
	return 0;
}
