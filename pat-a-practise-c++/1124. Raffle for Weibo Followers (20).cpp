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
	map<string,bool> mp;
	int N,M,Skip,next;
	cin>>M>>Skip>>next;
	bool flag=true;
	for(int i=1;i<=M;i++){
		string str;
		cin>>str;
		if(i==next){
			if(mp[str]){
				next++;
			}else{
				mp[str]=true;
				next+=Skip;
				flag=false;
				cout<<str<<endl;
			}
		}
	}
	if(flag){
		cout<<"Keep going..."<<endl;
	}
	return 0;
}


