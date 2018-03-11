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
	for(int i=12;i>=1;i--){
		for(int j=2;j<=sqrt(N);j++){
			long long sum=1;
			int len=0;
			for(int k=j;len<i;k++,len++){
				sum*=k;
//				printf("%02d ",k);
			}
//			printf("\n%d\n",sum);
			if(N%sum==0){
				printf("%d\n",len);
				for(int a=0;a<len;a++){
					printf("%d",j);
					if(a!=len-1){
						printf("*");
					}
					j++;
				}
				return 0;
			}
		}
	}
	printf("1\n%d",N);
	return 0;
}


