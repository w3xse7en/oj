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
bool isp(int n){
	if(n==1||n==0){
		return false;
	}
	for(int i=2;i<=sqrt(n);i++){
		if(n%i==0){
			return false;
		}
	}
	return true;
}

int main(){
//	freopen("in.txt","r",stdin);
	int N,M;cin>>N>>M;
	if(isp(N)==0){
		while(isp(++N)==0);
	}
	vector<int>v(N+1,0);
	int t,index;cin>>t;
	index=t%N;
	v[index]=1;
	cout<<index;
	for(int i=1;i<M;i++){
		cin>>t;
		index=t%N;
		if(v[index]==0){
			cout<<" "<<index;
			v[index]=1;
		}else{
			int flag = 0;
            for(int step = 1; step < N; step++) {
                index = (t + step * step) % N;
                if(v[index] == 0) {
                    v[index] = 1;
                    flag = 1;
                    cout<<" "<<index;
                    break;
                }
            }
            if(flag == 0) {
                cout<<" -";
            }
		}
	}
	return 0;
}


