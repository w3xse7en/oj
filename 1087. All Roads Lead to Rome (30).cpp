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
int INF=100000;
struct Info{
	int happy=0,sum=0,dist=INF,path=-1,cnt=0;
	vector<int>route;
	bool vis=false;
};
int N,K,diff=0;
map<string,int>ctoi;
map<int,string>itoc;
vector<vector<int> >g;
vector<Info>v;
void dij(){
	while(1){
	int mint=INF;
	int st=-1;
	for(int i=0;i<v.size();i++){
		if(v[i].vis==false&&mint>v[i].dist){
			mint=v[i].dist;
			st=i;
		}
	}
	if(st==-1){
		break;
	}
	v[st].vis=true;
	for(int i=0;i<g[st].size();i++){
		int ed=i;
		if(v[ed].vis==false&&g[st][ed]!=0){
			int d=v[st].dist+g[st][ed];
			if(d<v[ed].dist){
				v[ed].dist=d;
				v[ed].path=st;
				v[ed].sum=v[ed].happy+v[st].sum;
				v[ed].cnt=v[st].cnt+1;
				v[ed].route.clear();
				v[ed].route.push_back(st);
			}else if(d==v[ed].dist){
				v[ed].route.push_back(st);
				int sum=v[st].sum+v[ed].happy;
				if(v[ed].sum<sum){
					v[ed].sum=sum;
					v[ed].path=st;
					v[ed].cnt=v[st].cnt+1;
				}
				else if(v[ed].sum==sum){
					int cnt = v[st].cnt+1;
					if(cnt<v[ed].cnt){
						v[ed].path=st;
						v[ed].cnt=cnt;
					}
				}
			}
		}
	}
	}
}
vector<int> p;
void dsf(int root){
	p.push_back(root);
//	int i=root;
//	Info info=v[root];
//	printf("%d:%s h:%d d:%d p:%d cnt:%d \n",i,itoc[i].c_str(),info.sum,info.dist,info.path,info.cnt);
	if(v[root].path!=-1){
		dsf(v[root].path);
	}
}
void dsfdiff(int root){
//	cout<<root<<endl;
	if(v[root].route.size()){
		for(int i=0;i<v[root].route.size();i++){
			dsfdiff(v[root].route[i]);
		}
	}else{
		diff++;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	string st,ed="ROM";
	cin>>N>>K>>st;
	v.resize(N);
	ctoi[st]=0;
	itoc[0]=st;
	for(int i=0;i<N;i++){
		vector<int>t(N,0);
		g.push_back(t);
	}
	for(int i=1;i<N;i++){
		string cname;cin>>cname;
		int t;cin>>t;
		ctoi[cname]=i;
		itoc[i]=cname;
		v[i].happy=t;
	}
	for(int i=0;i<K;i++){
		string aname,bname;
		int l;
		cin>>aname>>bname;
		cin>>l;
		int a=ctoi[aname],b=ctoi[bname];
		g[a][b]=g[b][a]=l;
	}
	v[0].dist=0;
	dij();
	dsf(ctoi[ed]);
	dsfdiff(ctoi[ed]);
	Info info=v[ctoi[ed]];
	printf("%d %d %d %d\n",diff,info.dist,info.sum,info.sum/info.cnt);
	for(int i=p.size()-1;i>=0;i--){
		cout<<itoc[p[i]];
		if(i!=0){
			cout<<"->";
		}
	}
//	printf("\n");
//	for(int i=0;i<v.size();i++){
//		Info info=v[i];
//		printf("%d:%s h:%d d:%d p:%d cnt:%d \n",i,itoc[i].c_str(),info.sum,info.dist,info.path,info.cnt);
//	}
	return 0;
}


