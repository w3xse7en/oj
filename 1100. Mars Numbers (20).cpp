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
	int N,dlen=13,tlen=12;
	cin>>N;
	string dig[]={"tret","jan","feb","mar","apr","may","jun","jly","aug","sep","oct","nov","dec"};
	string ten[]={"tam","hel","maa","huh","tou","kes","hei","elo","syy","lok","mer","jou"};
	map<string,int>mp;
	for(int i=0;i<dlen;i++){
		mp[dig[i]]=i;
		if(i<dlen-1){
			mp[ten[i]]=(i+1)*13;
		}
//		cout<<mp[dig[i]]<<":"<<dig[i]<<"!!"<<mp[ten[i]]<<":"<<ten[i]<<endl;
	}
	cin.get();
	for(int i=0;i<N;i++){
		string str;getline(cin,str);
		int sum=0;
		if(isdigit(str[0])){
			int t=atoi(str.c_str());
			if(t/13!=0&&t%13==0){
				cout<<ten[t/13-1]<<endl;
			}else if(t/13!=0&&t%13!=0){
				cout<<ten[t/13-1]<<" "<<dig[t%13]<<endl;
			}else if(t/13==0){
				cout<<dig[t%13]<<endl;
			}
		}else{
			str+=' ';
			int len=0;
			for(int j=0;j<str.length();j++){
				if(' '==str[j]){
					string s=str.substr(j-len,len);
					sum+=mp[s];
					len=0;
				}else{
					len++;
				}
			}
			cout<<sum<<endl;
		}
	}
	return 0;
}


