#include <iostream>

int sum(int target);

using namespace std;

int main()
{
std::cout << sum(1000) << std::endl;
}

int sum(int target)
{
int sum = 0;

for(int counter = 0; counter<target; ++counter)
{

	if(counter%3==0) {
		sum+=counter;
			cout << counter << endl;

	}
	if(counter%5==0 && !(counter%3==0)){
		sum+=counter;
			cout  << counter << endl;

	}

}
return sum;
}

/*
std::vector<int> fives;
for(int counter = 0; reach_target<target; ++counter)
 fives.push_back(5*counter);
 
std:vector<int> threes;
for(int counter = 0; reach_target<target; ++counter)
 fives.push_back(3*counter);

std::vector<int> accumulator;
 

for(int counter = 0; sum<target; ++counter)
{
 int threes = 3*counter;
int  fives = 5*counter;

std::cout<< " 3s: " << threes <<std::endl;
std::cout <<" 5s: " << fives << std::endl;
std::cout <<" Check" << threes%5 << std::endl;

if(threes%5)
sum+=fives;
else
sum+=threes+fives;
}

*/