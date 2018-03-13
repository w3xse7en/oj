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
	string a,b;
	cin>>a>>b;
	vector<char> alp;
	for(int i=0,j=0;i<a.length();i++){
		if(a[i]!=b[j]){
			if(a[i]>='a'&&a[i]<='z'){
				a[i]=a[i]-'a'+'A';
			}
			bool flag=true;
			for(int j=0;j<alp.size();j++){
				if(alp[j]==a[i]){
					flag=false;
					break;
				}
			}
			if(flag){
				cout<<a[i];
				alp.push_back(a[i]);
			}
		}else{
			j++;
		}
	}
	
	return 0;
}


