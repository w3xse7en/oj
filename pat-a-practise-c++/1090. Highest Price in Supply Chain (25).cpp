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
vector<vector<int> >v;
double P,R,maxp;
int cnt=0;
void dsf(int root,double p){
//	cout<<p<<endl;
	if(v[root].size()==0){
		if(p>maxp){
			maxp=p;
			cnt=1;
		}else if(p==maxp){
			cnt++;
		}
	}
	for(int i=0;i<v[root].size();i++){
		dsf(v[root][i],p*(1+R/100));
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N;
	cin>>N>>P>>R;
	maxp=P;
	v.resize(N);
	int root;
	for(int i=0;i<N;i++){
		int t;cin>>t;
		if(t==-1){
			root=i;
		}else{
			v[t].push_back(i);
		}
	}
	dsf(root,P);
	printf("%.2f %d",maxp,cnt);
	return 0;
}


