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
	string id;
	int frank,lnumb,lrank,score;
};
bool cmpsrank(int &a,int &b){
	return a>b;
}
bool cmp(Info &a,Info &b){
	if(a.frank==b.frank){
		return a.id<b.id;
	}else{
		return a.frank<b.frank;
	}
}
int main(){
//	freopen("in.txt","r",stdin);
	int N,K;
	string id;
	cin>>N;
	vector<Info> v,t;
	vector<int> srank,allrank;
	for(int i=0;i<N;i++){
		cin>>K;
		t.clear();
		srank.clear();
		for(int j=0;j<K;j++){
			Info info;
			info.lnumb=i+1;
			cin>>info.id>>info.score;
			srank.push_back(info.score);
			allrank.push_back(info.score);
			t.push_back(info);
		}
		sort(srank.begin(),srank.end(),cmpsrank);
		int sr[101]={0};
		for(int j=0;j<K;j++){
			int ts=srank[j];
			if(sr[ts]==0){
				sr[ts]=j+1;
			}
		}
		for(int j=0;j<K;j++){
			int ts=t[j].score;
			t[j].lrank=sr[ts];
		}
		v.insert(v.end(),t.begin(),t.end());
	}
	sort(allrank.begin(),allrank.end(),cmpsrank);
	int sr[101]={0};
	for(int j=0;j<allrank.size();j++){
  		int ts=allrank[j];
		if(sr[ts]==0){
			sr[ts]=j+1;
		}
	}
	int vs=v.size();
    for(int i=0;i<v.size();i++){
    	int ts=v[i].score;
		v[i].frank=sr[ts];
	}
	sort(v.begin(),v.end(),cmp);
	cout<<vs<<endl;
	for(int i=0;i<vs;i++){
		cout<<v[i].id<<" "<<v[i].frank<<" "<<v[i].lnumb<<" "<<v[i].lrank<<endl;
	}
	return 0;
}

