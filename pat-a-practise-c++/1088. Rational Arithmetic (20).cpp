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
void formate(char f[100],long a,long b){
//	cout<<a<<"---"<<b<<endl;
	long ta=a,tb=b;
	while(ta%tb){
		long m=ta%tb;
		ta=tb;
		tb=m;
	}
	a/=tb,b/=tb;
	double t=a/(double)b;
	if(a%b==0){
		if(t>=0){
			sprintf(f,"%ld",a/b);
		}else{
			sprintf(f,"(%ld)",a/b);
		}
	}else{
		if(t>0&&t<1){
			sprintf(f,"%ld/%ld",a,b);
		}
		else if(t>1){
			sprintf(f,"%d %ld/%ld",(int)t,a%b,b);
		}
		else if(t<0&&t>-1){
			sprintf(f,"(-%ld/%ld)",abs(a),abs(b));
		}
		else if(t<-1){
			sprintf(f,"(%d %ld/%ld)",(int)t,abs(a%b),abs(b));
		}
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	long a,b,c,d;
	scanf("%ld/%ld %ld/%ld",&a,&b,&c,&d);
//	printf("%ld/%ld %ld/%ld\n",a,b,c,d);
	char A[100],B[100],C[100];
	formate(A,a,b);
	formate(B,c,d);
	formate(C,a*d+c*b,b*d);
	printf("%s + %s = %s\n",A,B,C);
	formate(C,a*d-c*b,b*d);
	printf("%s - %s = %s\n",A,B,C);
	formate(C,a*c,b*d);
	printf("%s * %s = %s\n",A,B,C);
	if(c!=0){
		formate(C,a*d,c*b);
		printf("%s / %s = %s\n",A,B,C);
	}else{
		printf("%s / %s = Inf\n",A,B);
	}
	
	return 0;
}


