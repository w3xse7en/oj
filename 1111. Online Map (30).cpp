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
int INF=0xffffff;
struct Info{
	int len=0,tim=0,dist=INF,path=-1,cnt=0;
	bool vis=false;
};
struct Point{
	int len=0,tim=0;
};
vector<Info>v;
vector<int>sest,fest;
vector<vector<Point> >g;
int N,M;
void dijshort(){
	while(1){
		int st=-1,mint=INF;
		for(int i=0;i<v.size();i++){
			if(v[i].vis==false&&mint>v[i].dist){
				st=i;	
			}
		}
		if(st==-1){
			break;
		}
		v[st].vis=true;
		for(int i=0;i<g[st].size();i++){
			int ed=i;
			Info sti=v[st],edi=v[ed];
			Point p=g[st][ed];
			if(p.len!=0&&sti.dist+p.len<edi.dist){
				v[ed].dist=sti.dist+p.len;
				v[ed].path=st;
				v[ed].cnt++;
				v[ed].tim=v[st].tim+p.tim;
//				printf("! %d-%d:%d %d %d %d\n",st,ed,v[ed].dist,v[ed].path,v[ed].tim,v[ed].cnt);
			}else if(p.len!=0&&sti.dist+p.len==edi.dist){
				if(v[ed].tim>v[st].tim+p.tim){
					v[ed].path=st;
					v[ed].tim=v[st].tim+p.tim;
				}
			}
		}
	}
	return;
}
void dijfast(){
	while(1){
		int st=-1,mint=INF;
		for(int i=0;i<v.size();i++){
			if(v[i].vis==false&&mint>v[i].dist){
				st=i;	
			}
		}
		if(st==-1){
			break;
		}
		v[st].vis=true;
		for(int i=0;i<g[st].size();i++){
			int ed=i;
			Info sti=v[st],edi=v[ed];
			Point p=g[st][ed];
			if(p.len!=0&&sti.dist+p.tim<edi.dist){
				v[ed].dist=sti.dist+p.tim;
				v[ed].path=st;
				v[ed].cnt=v[st].cnt+1;
//				printf("! %d-%d:%d %d %d %d\n",st,ed,v[ed].dist,v[ed].path,v[ed].tim,v[ed].cnt);
			}else if(p.len!=0&&sti.dist+p.tim==edi.dist){
				if(v[ed].cnt>v[st].cnt+1){
					v[ed].path=st;
					v[ed].cnt=v[st].cnt+1;
				}
			}
		}
	}
	return;
}
void dsf(int t,vector<int> &est){
	if(t!=-1){
		est.push_back(t);
		dsf(v[t].path,est);
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	cin>>N>>M;
	vector<Point>t(N);
	for(int i=0;i<N;i++){
		g.push_back(t);
	}
	for(int i=0;i<M;i++){
		int a,b,vt,len,tim;
		cin>>a>>b>>vt>>len>>tim;
		Point p;p.len=len,p.tim=tim;
		if(vt){
			g[a][b]=p;
		}else{
			g[a][b]=g[b][a]=p;
		}
	}
	int st,ed,D,T;
	cin>>st>>ed;
	//short
	v.resize(N);
	v[st].dist=0;
	dijshort();
	D=v[ed].dist;
	dsf(ed,sest);
	//fast
	v.clear();
	v.resize(N);
	v[st].dist=0;
	dijfast();
	T=v[ed].dist;
	dsf(ed,fest);
	
//	for(int i=0;i<v.size();i++){
//		printf("%d:%d %d %d %d\n",i,v[i].dist,v[i].path,v[i].tim,v[i].cnt);
//	}
	bool flag=true;
	if(sest.size()==fest.size()){
		for(int i=0;i<sest.size();i++){
			if(fest[i]!=sest[i]){
				flag=false;break;
			}
		}
	}else{
		flag=false;
	}
	if(flag){
		printf("Distance = %d; Time = %d: ",D,T);
		for(int i=sest.size()-1;i>=0;i--){
			printf("%d",sest[i]);
			if(i!=0){
				printf(" -> ");
			}
		}
	}else{
		printf("Distance = %d: ",D);
		for(int i=sest.size()-1;i>=0;i--){
			printf("%d",sest[i]);
			if(i!=0){
				printf(" -> ");
			}
		}
		printf("\n");
		printf("Time = %d: ",T);
		for(int i=fest.size()-1;i>=0;i--){
			printf("%d",fest[i]);
			if(i!=0){
				printf(" -> ");
			}
		}
	}
//	printf("\n");
//	for(int i=0;i<fest.size();i++){
//		printf("%d ",fest[i]);
//	}
	return 0;
}


