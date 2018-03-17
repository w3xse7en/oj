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
bool ispal(string str){
	for(int i=0;i<str.size()/2;i++){
		if(str[i]!=str[str.size()-1-i]){
			return false;
		}
	}
	return true;
}
string pls(string stra,string strb){
	int up=0;
	char sum[1024]={0};
	for(int i=0;i<stra.size();i++){
		int a=atoi(stra.substr(i,1).c_str()),b=atoi(strb.substr(i,1).c_str());
		int t=a+b+up;
		sum[i]=t%10+'0';
		up=t/10;
	}
	if(up!=0){
		sum[strlen(sum)]=up+'0';
	}
	string str=sum;
	reverse(str.begin(),str.end());
	return str;
}
int main(){
//	freopen("in.txt","r",stdin);
	string a,b;cin>>a;
	bool flag=false;
	for(int i=0;i<10;i++){
		b=a;
		reverse(b.begin(),b.end());
		string c = pls(a,b);
		if(ispal(a)){
			flag=true;
			break;	
		}else{
			cout<<a<<" + "<<b<<" = "<<c<<endl;
		}
		a=c;
	}
	if(flag){
		cout<<a<<" is a palindromic number."<<endl;
	}else{
		cout<<"Not found in 10 iterations."<<endl;	
	}
	return 0;
}


