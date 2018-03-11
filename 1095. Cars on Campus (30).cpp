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
struct Info{
	string id;
	int time=-1,status;
};
vector<Info> v,w;
map<string,Info>plate;
map<string,int>stay;
vector<string>lest;
bool cmp(Info &a,Info &b){
	return a.time<b.time;
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,M;
	cin>>N>>M;
	for(int i=0;i<N;i++){
		int hh,mm,ss,time,stus;
		string id,status;
		cin>>id;
		scanf("%d:%d:%d",&hh,&mm,&ss);
		time=hh*60*60+mm*60+ss;
		cin>>status;
		if(status.compare("in")==0){
			stus=1;
		}else{
			stus=-1;
		}
		Info info;
		info.id=id,info.status=stus,info.time=time;
		v.push_back(info);
		stay[id]=0;
	}
	sort(v.begin(),v.end(),cmp);
	vector<Info>::iterator it,temp;
	for(it=v.begin();it!=v.end();it++){
		Info info = plate[it->id];
		if(info.time==-1){
			plate[it->id]=*it;
		}
	}
	int maxtime=0;
	for(int i=0;i<v.size();i++){
		Info info = plate[v[i].id];
		if(v[i].status==1){
			plate[v[i].id]=v[i];
		}else if(info.status==1&&v[i].status==-1){
			plate[v[i].id]=v[i];
			w.push_back(info);
			w.push_back(v[i]);
			stay[info.id]+=(v[i].time-info.time);
			int t=stay[info.id];
			if(t>maxtime){
				maxtime=t;
				lest.clear();
				lest.push_back(info.id);
			}else if(t==maxtime){
				lest.push_back(info.id);
			}
		}
//		printf("%s %02d:%02d:%02d %d %d\n",v[i].id.c_str(),v[i].time/3600,v[i].time/60%60,v[i].time%60,v[i].time,v[i].status);
	}
//	printf("----------\n");
	sort(w.begin(),w.end(),cmp);
//	for(int i=0;i<w.size();i++){
//			printf("%s %02d:%02d:%02d %d %d\n",w[i].id.c_str(),w[i].time/3600,w[i].time/60%60,w[i].time%60,w[i].time,w[i].status);
//	}
	int cnt=0,j=0;
	for(int i=0;i<M;i++){
		int hh,mm,ss,time;
		scanf("%d:%d:%d",&hh,&mm,&ss);
		time=hh*60*60+mm*60+ss;
		for(;j<w.size();j++){
			if(w[j].time<=time){
				cnt+=w[j].status;
			}else{
				break;
			}
		}
		cout<<cnt<<endl;
	}
	sort(lest.begin(),lest.end());
	for(int i=0;i<lest.size();i++){
		cout<<lest[i]<<" ";
	}
	printf("%02d:%02d:%02d",maxtime/3600,maxtime/60%60,maxtime%60);
	return 0;
}


