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
bool check(int n){
	if(n==1){
		return false;
	}
	bool isp=true;
	for(int i=2;i<=sqrt(n);i++){
		if(n%i==0){
			isp=false;
			break;
		}
	}
//	cout<<"is:"<<isp<<" n:"<<n<<" sq:"<<sqrt(n)<<endl;
	return isp;
}
int main(){
//	freopen("in.txt","r",stdin);
	int n,m;
	while(cin>>n){
		if(n<0){
			break;
		}
		cin>>m;
		if(check(n)==false){
			cout<<"No"<<endl;
			continue;
		}
		char str[10000];
		int i=0;
		while(1){
			if(n!=0){
				str[i++]=n%m+'0';
			}else{
				break;
			}
			n/=m;
		}
		str[i]='\0';
		int ren =0,pw;
		for(int i=strlen(str)-1;i>=0;i--){
			pw = pow(m,strlen(str)-i-1);
			ren += (str[i]-'0')*pw;
//			printf("!%d %d %d\n",ren,str[i]-'0',pw);
		}
		if(check(ren)){
			cout<<"Yes"<<endl;
		}else{
			cout<<"No"<<endl;
		}
//		cout<<n<<" "<<str<<endl;
	}
	return 0;
}
