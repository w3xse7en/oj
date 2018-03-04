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
	vector<int>v;
	for(int i=0;i<N;i++){
		int t;
		cin>>t;
		v.push_back(t);
	}
	sort(v.begin(),v.end());
	int t=v[0];
	for(int i=1;i<v.size();i++){
		t=(t+v[i])/2;
	}
	cout<<t<<endl;; 
	return 0;
}


