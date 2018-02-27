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
struct Cust{
	int t,p;
};
bool cmp(Cust &a,Cust &b){
	return a.t<b.t;
}
int Open = 8*60*60,Close=17*60*60;
int main(){
//	freopen("in.txt","r",stdin);
	int n,m;
	cin>>n>>m;
	Cust cust;
	vector<Cust> v;
	vector<int> w(m,Open);
	for(int i=0;i<n;i++){
		int hh,mm,ss,p;
		scanf("%d:%d:%d %d",&hh,&mm,&ss,&p);
//		printf("%d:%d:%d %d\n",hh,mm,ss,p);
		int tm=hh*60*60+mm*60+ss;
		if(p>60){
			p=60;
		}
		if(tm>Close){
			continue;
		}
			cust.t=tm;
			cust.p=p*60;
			v.push_back(cust);
	}
	sort(v.begin(),v.end(),cmp);
	double sum=0;
	int cnt=0;
	for(int i=0;i<v.size();i++){
//		printf("%d %d\n",v[i].t,v[i].p);
		int index,minT=25*60*60;
		for(int j=0;j<m;j++){
			if(minT>w[j]){
				minT=w[j];
				index=j;
			}
		}
		if(w[index]>Close){
			break;
		}
		if(w[index]>v[i].t){
			sum+=(w[index]-v[i].t);
			w[index]+=v[i].p;
		}else{
			w[index]=v[i].t+v[i].p;
		}
		cnt++;
//		printf("%.0f @%d %d %d ",sum,index,w[index],v[i].t);
//		printf(" !%d %d\n",index,w[index]);
	}
	if(cnt==0){
		printf("0.0");
	}else{
		printf("%.1f",sum/cnt/60);	
	}
	return 0;
}
