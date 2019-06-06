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
struct Cust{
	int id=-1,p=0,ed=100000;
};
int main(){
//	freopen("in.txt","r",stdin);
	int N,M,K,Q;
	cin>>N>>M>>K>>Q;
	vector<Cust>result;
	vector<vector<Cust> > w;
	for(int i=0;i<N;i++){
		vector<Cust>m(M);
		w.push_back(m);
	}
	bool first=true;
	int i=0,nt=0,mt=0;
	while(1){
		int p;
		Cust cust;
		if(i<K){
			cin>>p;
			cust.id=i+1;
			cust.p=p;
			cust.ed=p;
		}
		if(i<N*M){
			if(mt==0){
				cust.p=0;
				w[nt][mt]=cust;
			}else{
				cust.p=w[nt][mt-1].ed;
				cust.ed+=w[nt][mt-1].ed;
				w[nt][mt]=cust;
			}
			if(nt==N-1){
				nt=0;
				mt++;
			}else{
				nt++;
			}
			i++;
			continue;
		}
		int mint=100000,index=0,cnt=0;
		for(int j=0;j<N;j++){
			if(mint>w[j][0].ed){
				mint=w[j][0].ed;
				index=j;
			}
			if(w[j][0].id==-1){
				cnt++;
			}
//			printf("%d id:%d %d %d\n",j,w[j][0].id,w[j][0].p,w[j][0].ed);
		}
		if(cnt==N){
			break;
		}
//		printf("index:%d\n",index);
		if(w[index][0].id!=-1){
			result.push_back(w[index][0]);
		}
		for(int j=0;j<M-1;j++){
			w[index][j]=w[index][j+1];
		}
		cust.p=w[index][M-1].ed;
		cust.ed+=w[index][M-1].ed;
		w[index][M-1]=cust;
		i++;
	}
	for(int i=0;i<Q;i++){
		int q;
		cin>>q;
		for(int j=0;j<result.size();j++){
			if(result[j].id==q){
//			printf("!!!%d %d %d\n",result[j].id,result[j].p,result[j].ed);
			if(result[j].p<540){
				int t = result[j].ed;
				printf("%02d:%02d\n",(t + 480) / 60, (t + 480) % 60);
			}else{
				printf("Sorry\n");
			}
			}
		}
	}
	return 0;
}
