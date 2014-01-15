#include <iostream>

int sum(int target);

using namespace std;

int main()
{
int max_number = 4000000;
int result = sum(max_number);
std::cout << "Sum: " << result << std::endl;
}

int sum(int target)
{
int sum            = 0;
int current        = 2;
int last           = 1;
int second_to_last = 0;

while(last < target-second_to_last)
{    

	current = last + second_to_last;
	if (current%2 == 0)
		sum+=current;
		
	second_to_last = last;
	          last = current;
			  
cout << current << endl;

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