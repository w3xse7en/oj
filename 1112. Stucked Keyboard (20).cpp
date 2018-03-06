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
struct Info{
	char key;
	int index;
};
bool cmp(Info &a,Info &b){
	return a.index<b.index;
}
int main(){
//	freopen("in.txt","r",stdin); 
	int N;
	cin>>N;
	char str[1010]={0},result[1010]={0};
	int key[1010]={0};
	cin>>str;
	for(int i=0;i<strlen(str);i++){
		int ta=str[i];
		if(key[ta]!=-1&&i!=strlen(str)-1){
			bool flag=true;
			for(int j=i+1,k=0;j<strlen(str)&&k<N-1;j++,k++){
				int tb=str[j];
//				printf("%c %c %d %d\n",ta,tb,i,j);
				if(ta!=tb){
					flag=false;
					key[ta]=-1;
					break;
				}
			}
			if(flag){
				if(key[ta]==0){
					key[ta]=i+1;
				}
				i=i+N-1;
			}
		}
	}
	vector<Info>v;
	for(int i=0;i<1010;i++){
		if(key[i]>0){
//			printf("%c %d %d\n",i,i,key[i]);
			Info info;
			info.key=i;
			info.index=key[i];
			v.push_back(info);
		}
	}
	sort(v.begin(),v.end(),cmp);
	int rlen=0;
	for(int i=0;i<strlen(str);i++){
		int t=str[i];
		result[rlen]=t;
		rlen++;
		if(key[t]>0){
			i=N+i-1;
		}
	}
	result[rlen]=0;
	for(int i=0;i<v.size();i++){
		cout<<v[i].key;
	}
	printf("\n");
	cout<<result<<endl;
	return 0;
}


