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
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N;
	vector<int> couple(100001,-1);
	for(int i=0;i<N;i++){
		int a,b;
		cin>>a>>b;
		couple[a]=b;
		couple[b]=a;
	}
	cin>>N;
	for(int i=0;i<N;i++){
		int t;
		cin>>t;
		int b=couple[t];
		if(couple[t]!=-1){
			if(couple[b]!=-2){
				couple[t]=-2;
			}else{
				couple[t]=couple[b]=-3;
			}
		}else{
			couple[t]=-2;
		}
	}
	vector<int> v;
	for(int i=0;i<100001;i++){
		if(couple[i]==-2){
			v.push_back(i);
		}
	}
	cout<<v.size()<<endl;
	for(int i=0;i<v.size();i++){
		printf("%05d",v[i]);
		if(i!=v.size()-1){
			cout<<" ";
		}
	}
}


