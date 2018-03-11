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
int main(){
//	freopen("in.txt","r",stdin);
	char c;
	long long p=0,pa=0,sum=0;
	while(cin.get(c)&&c!='\n'){
		if(c=='P'){
			p++;
		}
		if(c=='A'){
			pa+=p;
		}
		if(c=='T'){
			sum+=pa;
		}
	}
	cout<<sum%1000000007<<endl;
	return 0;
}


