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
bool cmp(int &a,int &b){
	return a>b;
}
vector<vector<int> > g;
void pf(){
		for(int i=0;i<g.size();i++){
			for(int j=0;j<g[i].size();j++){
				printf("%d",g[i][j]);
				if(j!=g[i].size()-1){
					printf(" ");
				}
			}
			printf("\n");
		}
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N;
	cin>>N;
	vector<int>v;
	int m,n,maxt=N;
	for(int i=1;i<=N;i++){
		int a=i,b=N/i;
		if(abs(a-b)<maxt&&a*b==N&&b>=a){
			maxt=abs(a-b);
			n=a,m=b;
		}
		int t;cin>>t;
		v.push_back(t);
	}
	for(int i=0;i<m;i++){
		vector<int>t(n,0);
		g.push_back(t);
	}
//	printf("%d %d\n",m,n);
	sort(v.begin(),v.end(),cmp);
	int x=0,y=0,xl=n,yl=m-1,index=0;
	while(index<v.size()){
		if(index<v.size()){
			for(int i=0;i<xl;i++){
				g[y][x++]=v[index++];
			}
			xl--;
			x--;
			y++;
		}
		if(index<v.size()){
			for(int i=0;i<yl;i++){
				g[y++][x]=v[index++];
			}
			yl--;
			y--;
			x--;
		}
		if(index<v.size()){
			for(int i=0;i<xl;i++){
				g[y][x--]=v[index++];
			}
			xl--;
			x++;
			y--;
		}
		if(index<v.size()){
			for(int i=0;i<yl;i++){
				g[y--][x]=v[index++];
			}
			yl--;
			x++;
			y++;
		}
	}
//	printf("%d %d : %d %d\n",x,y,xl,yl);
	pf();
	
	return 0;
}


