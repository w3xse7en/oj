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
struct Node{
	bool vis=false;
	int rcnt=0,ccnt=0,line=0,stop;
};
map<int,vector<int> >mp;
int line[10001][10001]={0},minrcnt,minccnt;
vector<Node>route,result;
void dsf(int st,int stop,map<int,Node> &node,int cnt){
	if(st==stop){
		if(minrcnt>cnt){
			minrcnt=cnt;
			node[st].rcnt=cnt;
			result=route;
		}else if(minrcnt==cnt){
			int acnt=0,bcnt=0;
			for(int i=1;i<result.size();i++){
				if(route[i].line!=route[i-1].line){
					acnt++;
				}
				if(result[i].line!=result[i-1].line){
					bcnt++;
				}
			}
			if(acnt<bcnt){
				result=route;
			}
		}
		return;
	}
	for(int i=0;i<mp[st].size();i++){
		int ed=mp[st][i];
		if(node[ed].vis==false){
			node[ed].stop=ed;
			if(line[st][ed]!=node[st].line){
				node[ed].line=line[st][ed];
			}else{
				node[ed].line=node[st].line;
			}
			route.push_back(node[ed]);
			node[ed].vis=true;
			dsf(ed,stop,node,cnt+1);
			node[ed].vis=false;
			route.pop_back();
		}
	}
	return;
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,M;
	cin>>N;
	for(int i=1;i<=N;i++){
		int a,b;
		cin>>M>>a;
		for(int j=0;j<M-1;j++){
			cin>>b;
			mp[a].push_back(b);
			mp[b].push_back(a);
			line[a][b]=line[b][a]=i;
			a=b;
		}
	}
	cin>>N;
	for(int i=0;i<N;i++){
		int a,b;
		cin>>a>>b;
		route.clear();
		map<int,Node> nd;
		nd[a].stop=a;
		route.push_back(nd[a]);
		minccnt=10001,minrcnt=10001;
		dsf(a,b,nd,0);
		cout<<nd[b].rcnt<<endl;;
		route.clear();
//		reverse(result.begin(),result.end());
		for(int j=0;j<result.size()-1;j++){
//			Node node = result[j];
//			printf("%d %d\n",node.stop,node.line);
			if(result[j].line!=result[j+1].line){
				route.push_back(result[j]);
			}
		}
		route.push_back(result[result.size()-1]);
		for(int j=1;j<route.size();j++){
			printf("Take Line#%d from %04d to %04d.\n",route[j].line,route[j-1].stop,route[j].stop);
		}
	}
	return 0;
}


