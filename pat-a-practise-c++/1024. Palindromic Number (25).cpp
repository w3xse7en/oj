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
bool check(string n){
//	cout<<"!"<<n<<endl;
	int len=n.size();
	for(int i=0;i<len;i++){
		if(n[i]!=n[len-1-i]){
			return false;
		}
	}
	return true;
}
string stradd(string a,string b){
	char sum[10000];
	string strsum;
	int up=0,len=a.size();
	for(int i=0;i<len;i++){
		int ai=atoi(a.substr(i,1).c_str());
		int bi=atoi(b.substr(i,1).c_str());
		int abi=ai+bi+up;
		sum[i]=abi%10+'0';
		up=abi/10;
	}
	if(up!=0){
		sum[len]=up+'0';
	}
	strsum=sum;
	reverse(strsum.begin(),strsum.end());
	return strsum;
}
int main(){
//	freopen("in.txt","r",stdin);
	int k;
	string a,b;
	cin>>a>>k;
	for(int i=0;i<k;i++){
		if(check(a)){
			cout<<a<<endl;
			cout<<i<<endl;
			return 0;
		}
		b=a;
		reverse(b.begin(),b.end());
		a=stradd(a,b);
//		cout<<"@"<<a<<endl; 
	}
	cout<<a<<endl;
	cout<<k<<endl;
    return 0;
}

