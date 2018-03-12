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
struct Node{
	int x,y,z;
};
int g[70][1300][130]={0},vis[70][1300][130];
int X[6] = {1, 0, 0, -1, 0, 0};
int Y[6] = {0, 1, 0, 0, -1, 0};
int Z[6] = {0, 0, 1, 0, 0, -1};
int N,M,L,T;
bool judge(int x, int y, int z) {
//	printf("%d %d %d\n",x,y,z);
    if(x < 0 || x >= L || y < 0 || y >= N || z < 0 || z >= M || g[x][y][z] == 0 || vis[x][y][z] == true){
		return false;
	} 
    return true;
}
int check(int x,int y,int z){
	int cnt=0;
	Node node;
	node.x=x,node.y=y,node.z=z;
	list<Node> l;
	l.push_back(node);
	vis[x][y][z]=true;
	while(!l.empty()){
		Node top=l.front();
		l.pop_front();
		cnt++;
		for(int i=0;i<6;i++){
			int tx=top.x+X[i],ty=top.y+Y[i],tz=top.z+Z[i];
			if(judge(tx,ty,tz)){
				vis[tx][ty][tz]=true;
				node.x=tx,node.y=ty,node.z=tz;
				l.push_back(node);
			}
		}
	}
	if(cnt>=T){
		return cnt;
	}else{
		return 0;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	cin>>N>>M>>L>>T;
	for(int i=0;i<L;i++){
		int cnt=0;
		for(int j=0;j<N;j++){
			for(int k=0;k<M;k++){
				int t;cin>>t;
				g[i][j][k]=t;
			}
		}
	}
	int cnt=0;
	for(int i=0;i<L;i++){
		for(int j=0;j<N;j++){
			for(int k=0;k<M;k++){
//				printf("%d ",g[i][j][k]);
				if(g[i][j][k]!=0&&vis[i][j][k]==false){
					cnt+=check(i,j,k);
				}
			}
//			printf("\n");
		}
//		printf("\n");
	}
	cout<<cnt<<endl;
	
	return 0;
}


