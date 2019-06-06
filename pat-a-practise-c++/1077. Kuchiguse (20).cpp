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
#include <list>
#include <queue>
using namespace std;

int main(){
//	freopen("in.txt","r",stdin);
	int N;cin>>N;
	vector<string> v;
	cin.get();
	for(int i=0;i<N;i++){
		string t;
		getline(cin,t);
		reverse(t.begin(),t.end());
		v.push_back(t);
	}
	string str=v[0];
	for(int i=1;i<v.size();i++){
		string t=v[i];
		int len=str.length()>t.length()?t.length():str.length();
		for(int j=0;j<len;j++){
			if(str[j]!=t[j]){
				str=str.substr(0,j);
				break;
			}
		}
	}
	reverse(str.begin(),str.end());
	if(str.length()!=0){
		cout<<str<<endl;
	}else{
		cout<<"nai"<<endl;
	}
	return 0;
}


