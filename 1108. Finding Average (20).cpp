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
	int N;
	double sum=0.0,cnt=0.0;
	cin>>N;
	for(int i=0;i<N;i++){
		char a[100],b[100];
		cin>>a;
		double d = atof(a);
		sprintf(b,"%.2f",d);
//		printf("%s | %f | %s\n",a,d,b);
		bool flag=true;
		for(int j=0;j<strlen(a);j++){
			if(a[j]!=b[j]){
				flag=false;
			}
		}
		if(d>1000.0||d<-1000.0){
			flag=false;
		}
		if(flag){
			cnt++;
			sum+=d;
		}else{
			printf("ERROR: %s is not a legal number\n",a);
		}
	}
	if(cnt==1){
		printf("The average of 1 number is %.2f",sum);
	}else if(cnt>1){
		printf("The average of %d numbers is %.2f",(int)cnt,sum/cnt);
	}else{
		printf("The average of 0 numbers is Undefined");
	}
	return 0;
}


