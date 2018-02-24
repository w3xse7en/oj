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
vector<int> toll;
struct Info{
	string id;
	int mtime;
	bool status;
};
vector<Info> v;
bool cmp(Info &a,Info &b){
	if(a.id==b.id){
		return a.mtime<b.mtime;
	}
	return a.id<b.id;
}
int main(){
	map<int,string> ftime;
	freopen("in.txt","r",stdin);
	for(int i=0;i<24;i++){
		int cents;
		cin>>cents;
		for(int j=0;j<60;j++){
			toll.push_back(cents);
		}
	}
	int len = toll.size();
	for(int i=0;i<30;i++){
		for(int j=0;j<len;j++){
			toll.push_back(toll[j]);
		}
	}
	int n;
	cin>>n;
	for(int i=0;i<n;i++){
		char name[25],status[20],ftc[20];
		int MM,dd,hh,mm;
		scanf("%s %d:%d:%d:%d %s",&name,&MM,&dd,&hh,&mm,&status);
//		printf("%s %d:%d:%d:%d %s\n",name,MM,dd,hh,mm,status);
		
		bool flag=false;
		string astr = status,bstr="on-line",id=name,ft;
		if(astr.compare(bstr)==0){
			flag= true;
		}
		
		int index=-1,dtime=mm+60*hh+60*24*dd;
		sprintf(ftc,"%02d:%02d:%02d",dd,hh,mm);
		ftime[dtime]=string(ftc);
		sprintf(ftc," %02d",MM);
		
		Info info;
		info.id = id+string(ftc);
		info.mtime = dtime;
		info.status=flag;
		v.push_back(info);
	}
	sort(v.begin(),v.end(),cmp);
	string id;
	bool first = true;
	double total=0;
	for(int i=0;i<v.size();i++){
		if(v[i].status==0&&v[i-1].status==1&&v[i].id.compare(v[i-1].id)==0){
			if(id.compare(v[i].id)!=0){
				if(first){
					first=false;
				}else{
					printf("Total amount: $%.2f\n",total);
				}
				total=0;
				printf("%s\n",v[i].id.c_str(),v[i].mtime,v[i].status);
			}
			id = v[i].id;
			int st = v[i-1].mtime,ed=v[i].mtime;
			double sum = 0;
			for(int j=st;j<ed;j++){
				sum+=toll[j];
			}
			sum/=100;
			total+=sum;
			printf("%s %s %d $%.2f\n",ftime[st].c_str(),ftime[ed].c_str(),ed-st,sum);
		}
	}
	printf("Total amount: $%.2f\n",total);
	return 0;
}
