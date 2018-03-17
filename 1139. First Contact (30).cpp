#include <cstdio>
#include <algorithm>
#include <vector>
#include <iostream>
#include <cstring>
#include <cctype>
#include <map>
#include <iterator>
#include <cmath>
#include <list>
#include <set>
using namespace std;
struct Info{
	int c,d;
};
vector<int> fv(vector<int> v,bool man){
	vector<int> t;
	for(int j=0;j<v.size();j++){
		if(man){
			if(v[j]>=0){
				t.push_back(v[j]);	
			}
		}else{
			if(v[j]<=0){
				t.push_back(v[j]);
			}
		}
	}
	return t;
}
bool f(vector<int>v,int t){
	for(int i=0;i<v.size();i++){
		if(v[i]==t){
			return true;
		}
	}
	return false;
}
bool cmp(Info &a,Info &b){
	if(a.c==b.c){
		return a.d<b.d;
	}else{
		return a.c<b.c;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,M;
	cin>>N>>M;
	map<int,vector<int> > mp;
	for(int i=0;i<M;i++){
		int a,b;cin>>a>>b;
		mp[a].push_back(b);
		mp[b].push_back(a);
	}
	map<int,vector<int> >::iterator it;
//	for(it=mp.begin();it!=mp.end();it++){
//		printf("%d:",it->first);
//		vector<int> v=it->second;
//		for(int j=0;j<v.size();j++){
//			printf("%d ",v[j]);
//		}
//		printf("\n");
//	}
	int K;cin>>K;
	for(int i=0;i<K;i++){
		int a,b;cin>>a>>b;
		vector<Info>fs;
		vector<int> vc=mp[a];
		bool cman,dman;
		if(a<0&&b<0){
			cman=dman=false;
		}else if(a>=0&&b>=0){
			cman=dman=true;
		}else if(a<0&&b>=0){
			cman=false,dman=true;
		}else if(a>=0&&b<0){
			cman=true,dman=false;
		}
		
		vc=fv(vc,cman);
		for(int j=0;j<vc.size();j++){
			int c=vc[j];
			vector<int>vd=mp[c];
			vd=fv(vd,dman);
			for(int k=0;k<vd.size();k++){
				int d=vd[k];
				vector<int>vb=mp[d];
				int stus = f(vb,b);
				if(stus){
					set<int> s;
					s.insert(a);
					s.insert(b);
					s.insert(c);
					s.insert(d);
					if(s.size()==4){
	//					printf("%d:%d %d %d %d %d\n",i,a,b,c,d,stus);
						Info info;
						info.c=abs(c);
						info.d=abs(d);
						fs.push_back(info);
					}
				}
			}
		}
		sort(fs.begin(),fs.end(),cmp);
		cout<<fs.size()<<endl;
		for(int j=0;j<fs.size();j++){
			Info info =fs[j];
			printf("%04d %04d\n",info.c,info.d);
		}
	}
	
	return 0;
}


