#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include<cmath>
using namespace std;
struct City{
	int mx=0,store,dist=100000,pbmc=0;
	bool vis=false;
	vector<int> pos;
};
vector<City> citys;
vector<vector<int> >roads;
vector<int>v;
int mx,N,target,M,hmx;
void dij(){
	while(1){
		
	int index=-1,tmin=10000;
	for(int i=0;i<N;i++){
		if(citys[i].vis==false&&tmin>citys[i].dist){
			tmin=citys[i].dist;
			index=i;
		}
	}
	if(index==-1){
		break;
	}
//	printf("! city:%d v:%d d:%d pbmc:%d st:%d mx:%d\n",index,citys[index].vis,citys[index].dist,citys[index].pbmc,citys[index].store,citys[index].mx);
	citys[index].vis=true;
	for(int i=0;i<N;i++){
		int r=roads[index][i],d=citys[index].dist;
		if(citys[i].vis==false&&r!=0){
		if(r+d<citys[i].dist){
			citys[i].dist=r+d;
			citys[i].pos.clear();
			citys[i].pos.push_back(index);
//			printf("@ city:%d v:%d d:%d pbmc:%d st:%d mx:%d\n",i,citys[i].vis,citys[i].dist,citys[i].pbmc,citys[i].store,citys[i].mx);
		
		}else if(r+d==citys[i].dist){
			citys[i].pos.push_back(index);
//			printf("# city:%d v:%d d:%d pbmc:%d st:%d mx:%d\n",i,citys[i].vis,citys[i].dist,citys[i].pbmc,citys[i].store,citys[i].mx);
		
		}
		}
	}
	}
}
int tmin=10000,tpbmc;
void sch(vector<int> p,vector<int>path,int pbmc,int pi){
	if(p[0]==-1){
		if(tmin>abs(pbmc)){
			tmin=abs(pbmc);
			tpbmc=pbmc;
			v=path;
		}else if(tmin==abs(pbmc)){
			if(tpbmc<0){
				tpbmc=pbmc;
				v=path;
			}
		}
		return;
	}
	for(int i=0;i<p.size();i++){
		path.push_back(p[i]);
		sch(citys[p[i]].pos,path,pbmc+citys[p[i]].pbmc,p[i]);
		path.pop_back();
	}
}
int main() {
//	freopen("in.txt","r",stdin);
	cin>>mx>>N>>target>>M;
	N++;
	hmx=mx/2;
	for(int i=0;i<N;i++){
		City city;
		if(i!=0){
			city.mx=mx;
			int store;
			cin>>store;
			city.store=store;
			city.pbmc=hmx-store;
		}else{
			city.dist=0;
			city.pos.push_back(-1);
		}
		citys.push_back(city);	
	}
	for(int i=0;i<M;i++){
		vector<int>t(N);
		roads.push_back(t);
	}
	for(int i=0;i<M;i++){
		int a,b,c;
		cin>>a>>b>>c;
		roads[a][b]=c;
		roads[b][a]=c;
	}
	dij();
//	for(int i=0;i<N;i++){
//		for(int j=0;j<citys[i].pos.size();j++){
//			printf("%d ",citys[i].pos[j]);
//		}
//		printf("\n");
//	}
	vector<int>pt;
	pt.push_back(target);
	sch(citys[target].pos,pt,citys[target].pbmc,target);
	if(tpbmc>0){
		printf("%d ",tpbmc);
	}else{
		printf("0 ");
	}
	for(int i=v.size()-1;i>=0;i--){
		printf("%d",v[i]);
		if(i!=0){
			printf("->");
		}
	}
	
	if(tpbmc<0){
		printf(" %d",-tpbmc);
	}else{
		printf(" 0");
	}
//	printf("\n");
//	for(int i=0;i<N;i++){
//		printf("city:%d v:%d p:%d d:%d pbmc:%d st:%d mx:%d\n",i,citys[i].vis,citys[i].pos,citys[i].dist,citys[i].pbmc,citys[i].store,citys[i].mx);
//		
//	}
	return 0;
}
