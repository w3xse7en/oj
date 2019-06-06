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
	vector<int> v;
	for(int i=0;i<N;i++){
		int t;cin>>t;
		v.push_back(t);
	}
	sort(v.begin(),v.end());
	cout<<N%2<<" ";
	long long a=0,b=0;
	for(int i=0;i<v.size()/2;i++){
		a+=v[i];
	}
	for(int i=v.size()/2;i<v.size();i++){
		b+=v[i];
	}
	cout<<b-a<<endl;
	return 0;
}


