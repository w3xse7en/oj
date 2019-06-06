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
bool isprime(int t){
	bool flag=true;
	if(t==1||t==0){
		flag=false;
	}
	for(int i=2;i<=sqrt(t);i++){
		if(t%i==0){
			flag=false;
			break;
		}
	}
	return flag;
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,rank[10001]={0};
	cin>>N;
	for(int i=1;i<=N;i++){
		int t;cin>>t;
		rank[t]=i;
	}
	cin>>N;
	// -1 checkd;0 kidding;1-N
	for(int i=0;i<N;i++){
		int t;cin>>t;
		if(rank[t]==0){
			printf("%04d: Are you kidding?\n",t);
		}else if(rank[t]==1){
			printf("%04d: Mystery Award\n",t);
			rank[t]=-1;
		}else if(rank[t]==-1){
			printf("%04d: Checked\n",t);
		}else{
			if(isprime(rank[t])){
				printf("%04d: Minion\n",t);
			}else{
				printf("%04d: Chocolate\n",t);
			}
			rank[t]=-1;
		}
	}
	return 0;
}


