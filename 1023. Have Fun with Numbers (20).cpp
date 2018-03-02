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
	char dig[30];
	cin>>dig;
	vector<int> ori,doub,tori(10),tdoub(10);
	int digs=strlen(dig),up=0;
	for(int i=digs-1;i>=0;i--){
		int odig=dig[i]-'0';
		tori[odig]++;
		int db=odig*2+up;
		doub.push_back(db%10);
		tdoub[db%10]++;
		up=db/10;
//		printf("odig:%d db:%d up:%d\n",odig,db,up);
	}
	if(up!=0){
		doub.push_back(up);
	}
	int flag=true;
	for(int i=0;i<10;i++){
		if(tori[i]!=tdoub[i]){
			flag=false;
			break;
		}
	}
	if(flag){
		cout<<"Yes"<<endl;
	}else{
		cout<<"No"<<endl;
	}
	for(int i=doub.size()-1;i>=0;i--){
		cout<<doub[i];
	}
    return 0;
}

