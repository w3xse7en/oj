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
int N,K,P,maxt=0;
vector<int>seq,result,v;
void dsf(int index,int sum,int sumseq,int cnt){
	if(sum>N||cnt>K){
		return;
	}
	if(sum==N&&cnt==K){
		if(sumseq>maxt){
			maxt=sumseq;
			result=seq;
		}
		return;
	}
	for(int i=index;i>=0;i--){
		seq.push_back(i+1);
		dsf(i,sum+v[i],sumseq+i+1,cnt+1);
		seq.pop_back();
	}
}
int main(){
//	freopen("in.txt","r",stdin); 
	cin>>N>>K>>P;
	for(int i=1;i<=N;i++){
		int t=pow(i,P);
		if(t<=N){
			v.push_back(t);
		}else{
			break;
		}
	}
	dsf(v.size()-1,0,0,0);
	if(maxt==0){
		cout<<"Impossible"<<endl;
		return 0;
	}
	printf("%d = ",N);
	for(int i=0;i<result.size();i++){
		if(i != 0) printf(" + ");
		printf("%d^%d",result[i],P);
	}
	return 0;
}


