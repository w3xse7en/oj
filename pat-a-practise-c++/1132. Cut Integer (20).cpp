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
	for(int i=0;i<N;i++){
		string Z;
		double z,a,b;
		int half;
		cin>>Z;
		half=Z.length()/2;
		string A=Z.substr(0,half),B=Z.substr(half,half);
		z=atof(Z.c_str()),a=atof(A.c_str()),b=atof(B.c_str());
		double result=z/(a*b);
		if(a==0||b==0){
			cout<<"No"<<endl;
		}else if(result-(int)result==0){
			cout<<"Yes"<<endl;
		}else{
			cout<<"No"<<endl;
		}
	}
	return 0;
}


