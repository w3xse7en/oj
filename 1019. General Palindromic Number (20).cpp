#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include<cmath>
using namespace std;
int main() {
//	freopen("in.txt","r",stdin);
	long long n,radix;
	cin>>n>>radix;
	vector<long long> v;
	while(n){
		v.push_back(n%radix);
		n/=radix;
	}
	bool flag=true;
	int vs=v.size();
	for(int i=0;i<vs/2;i++){
		if(v[i]!=v[vs-i-1]){
			flag=false;
		}
	}
	if(flag){
		cout<<"Yes"<<endl;
	}else{
		cout<<"No"<<endl;
	}
	for(int i=vs-1;i>=0;i--){
		cout<<v[i];
		if(i!=0){
			cout<<" ";
		}
	}
	if(vs==0){
		cout<<"0";
	}
	return 0;
}
