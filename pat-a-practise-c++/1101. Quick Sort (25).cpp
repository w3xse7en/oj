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
struct Info{
	int t,mint,maxt;
};
int main(){
//	freopen("in.txt","r",stdin); 
	int N;cin>>N;
	vector<Info>v,vt;
	vector<int>result;
	int maxt=-1;
	for(int i=0;i<N;i++){
		int t;cin>>t;
		Info info;
		info.maxt=maxt;
		info.t=t;
		v.push_back(info);
		if(t>maxt){
			maxt=t;
		}
	}
	int mint=pow(10,9)+1;
	
	for(int i=v.size()-1;i>=0;i--){
		v[i].mint=mint;
		if(mint>v[i].t){
			mint=v[i].t;
		}
		if(v[i].t>v[i].maxt&&v[i].t<v[i].mint){
			result.push_back(v[i].t);
		}
//		cout<<v[i].maxt<<"---"<<v[i].t<<"---"<<v[i].mint<<endl;
	}
	if(result.size()!=0){
		cout<<result.size()<<endl;
		for(int i=result.size()-1;i>=0;i--){
			cout<<result[i];
			if(i!=0){
				cout<<" ";
			}
		}
	}else{
		cout<<"0"<<endl;
		cout<<""<<endl;
	}
	return 0;
}


