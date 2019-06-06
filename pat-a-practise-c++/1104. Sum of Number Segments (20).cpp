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
	int N;
	cin>>N;
	double sum=0,s,t;
	for(int i=1;i<=N;i++){
		cin>>t;
		s=t*i*(N+1-i);
		sum+=s;
	}
	printf("%.2f",sum);
	
	return 0;
}


