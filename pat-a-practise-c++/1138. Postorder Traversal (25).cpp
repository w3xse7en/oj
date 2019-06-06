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

bool first=true;
int pre[50001],in[50001];
void bd(int prest,int inst,int ined){
	if(first==false){
		return;
	}
	if(ined>inst){
//		for(int i=prest,cnt=0;cnt<ined-inst;i++,cnt++){
//			printf("%d ",pre[i]);
//		}
//		printf("\n");
//		for(int i=inst,cnt=0;cnt<ined-inst;i++,cnt++){
//			printf("%d ",in[i]);
//		}
//		printf("\n");
		int root=pre[prest];
		int i=inst;
		while(in[i]!=root){
			i++;
		}
		bd(prest+1,inst,i);
		bd(prest+1+i-inst,i+1,ined);
		if(first){
			cout<<root<<endl;
			first=false;
		}
	}
	
	
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	scanf("%d",&N);
	for(int i=0;i<N;i++){
		int t;scanf("%d",&t);
		pre[i]=t;
	}
	for(int i=0;i<N;i++){
		int t;scanf("%d",&t);
		in[i]=t;
	}
	bd(0,0,N-1);
	return 0;
}


