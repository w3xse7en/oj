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
bool cmp(int &a,int&b){
	return a>b;
}
int main(){
//	freopen("in.txt","r",stdin);
	int N, cnt=0;
	cin>>N;
	vector<int> ride;
	for(int i=1;i<=N;i++){
		int t;cin>>t;
		ride.push_back(t);
	}
	sort(ride.begin(),ride.end(),cmp);
	for(int i=0;i<ride.size();i++){
//		cout<<ride[i]<<endl;
		if(ride[i]>i+1){
			cnt++;
		}
	}
	cout<<cnt<<endl;
	return 0;
}


