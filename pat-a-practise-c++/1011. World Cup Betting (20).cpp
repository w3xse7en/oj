#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
using namespace std;
int main(){
//	freopen("in.txt","r",stdin);
	double w,t,l,sum=1;
	vector<double>v;
	for(int i=0;i<3;i++){
		double tm=0;
		cin>>w>>t>>l;
		if(w>tm){
			tm=w;
		}
		if(t>tm){
			tm=t;
		}
		if(l>tm){
			tm=l;
		}
		if(tm==w){
			printf("W");
		}else if(tm==t){
			printf("T");
		}else if(tm==l){
			printf("L");
		}
		printf(" ");
		sum*=tm;
	}
	double s = (sum*0.65-1)*2;
	printf("%.2f",s);

	return 0;
}
