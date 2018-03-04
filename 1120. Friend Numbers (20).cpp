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
	map<int,int>mp;
	for(int i=0;i<N;i++){
		char t[100];
		cin>>t;
		int sum=0;
		for(int j=0;j<strlen(t);j++){
			sum+=t[j]-'0';
		}
		mp[sum]=sum;
	}
	map<int,int>::iterator it;
	bool first=true;
	cout<<mp.size()<<endl;
	for(it=mp.begin();it!=mp.end();it++){
		if(first){
			cout<<it->second;
			first=false;
		}else{
			cout<<" "<<it->second;
		}
	}
}


