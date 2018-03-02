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
	int N,K;
	cin>>K;
	for(int i=0;i<K;i++){
		cin>>N;
		vector<int>v(N*2,0),q(N+1,0);
		bool flag=true;
		for(int j=1;j<=N;j++){
			int y;cin>>y;
			q[y]++;
			v[j+y-1]++;
			if(q[y]==2||v[j+y-1]==2){
				flag=false;
			}
		}
		if(flag){
			cout<<"YES"<<endl;
		}else{
			cout<<"NO"<<endl;
		}
	}
	
	return 0;
}


