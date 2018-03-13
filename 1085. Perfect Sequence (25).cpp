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
vector<long>v;
int f(int l,int r,double data){
	if(l<=r){
		int index=(l+r)/2;
//		cout<<index<<"-"<<data<<endl;
		if(data>v[index]){
			l=index+1;
		}else if(data<v[index]){
			r=index-1;
		}else {
			return index;
		}
		f(l,r,data);
	}else{
		return r;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,P;
	cin>>N>>P;
	for(long i=0;i<N;i++){
		double t;
		cin>>t;
		v.push_back(t);
	}
	int maxcnt=0;
	sort(v.begin(),v.end());
		for(int i=0;i<v.size();i++){
//			cout<<i<<":"<<v[i]<<endl;
			int index = f(i,v.size(),(double)v[i]*P);
//			cout<<"----"<<index<<endl;
			if(index==v.size()){
				index--;
			}
			if(index-i+1>maxcnt){
				maxcnt=index-i+1;
			}
			if(index==v.size()-1){
				break;
			}
		}
		cout<<maxcnt<<endl;
	return 0;
}


