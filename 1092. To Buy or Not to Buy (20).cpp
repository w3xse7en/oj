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
	char c;
	int eva[128]={0},shop[128]={0},shoplen=0,evalen=0;
	while(cin.get(c)&&c!='\n'){
		shop[c]++;
		shoplen++;
	}
	while(cin.get(c)&&c!='\n'){
		eva[c]++;
		evalen++;
	}
	int cnt=0;
	for(int i=0;i<128;i++){
		if(eva[i]>shop[i]){
			cnt+=eva[i]-shop[i];
		}
	}
	if(cnt==0){
		cout<<"Yes "<<shoplen-evalen<<endl;;
	}else{
		cout<<"No "<<cnt<<endl;
	}
	return 0;
}


