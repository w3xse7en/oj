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
	map<string,vector<string> > title,auther,publier,pubyear,keys;
    int n;
	cin>>n;
    cin.get();
    for(int i=0;i<n;i++){
		string id,ctitle,cauther,cpublier,cpubyear;
		getline(cin,id);
		getline(cin,ctitle);
		title[ctitle].push_back(id);
		
		getline(cin,cauther);
		auther[cauther].push_back(id);
		
		string tkeys;
		getline(cin,tkeys);
		tkeys+=" ";
		int st=0,len=0;
		string s;
		for(int i=0;i<tkeys.size();i++,len++){
			if(tkeys[i]==' '){
				s=tkeys.substr(st,len);
//				cout<<st<<" "<<len<<"!"<<s<<"!"<<endl;
				st=i+1;
				len=-1;
				keys[s].push_back(id);
			}
		}
		getline(cin,cpublier);
		publier[cpublier].push_back(id);
		
		getline(cin,cpubyear);
		pubyear[cpubyear].push_back(id);
	}
	cin>>n;
	cin.get();
	for(int i=0;i<n;i++){
		string s,ts,str;
		int st=0,len=0,index=0;
		getline(cin,ts);
		
		for(int i=0;i<ts.size();i++){
			if(ts[i]==':'){
				st=i+2;
				break;
			}
		}
		index = atoi(ts.substr(0,1).c_str());
		s=ts.substr(st,ts.size());
		cout<<ts<<endl;
		bool flag=false;
		if(index==1){
			sort(title[s].begin(),title[s].end());
			for(int i=0;i<title[s].size();i++){
				flag=true;
				cout<<title[s][i]<<endl;
			}
		}else if(index==2){
			sort(auther[s].begin(),auther[s].end());
			for(int i=0;i<auther[s].size();i++){
				flag=true;
				cout<<auther[s][i]<<endl;
			}
		}else if(index==3){
			sort(keys[s].begin(),keys[s].end());
			for(int i=0;i<keys[s].size();i++){
				flag=true;
				cout<<keys[s][i]<<endl;
			}
		}else if(index==4){
			sort(publier[s].begin(),publier[s].end());
			for(int i=0;i<publier[s].size();i++){
				flag=true;
				cout<<publier[s][i]<<endl;
			}
		}else if(index==5){
			sort(pubyear[s].begin(),pubyear[s].end());
			for(int i=0;i<pubyear[s].size();i++){
				flag=true;
				cout<<pubyear[s][i]<<endl;
			}
		}
		if(flag==false){
			cout<<"Not Found"<<endl;
		}
	}
    return 0;
}

