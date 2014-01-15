#include <iostream>
#include <vector>
#include <algorithm>

long long FindNewPrime();
bool CheckIfIsNewPrime(long long number);
void AddNewPrime ( long long prime);

std::vector<long long> v_primes;

using namespace std;

int main()
{
//long long number = 600851475143LL;

long long number_of_prime = 10001;

cout << number_of_prime << "'s prime:";


v_primes.push_back(2);
v_primes.push_back(3);

for (int i =2; i<10001; ++i)
{
	AddNewPrime(FindNewPrime());
}

cout << v_primes.back()<<endl;


}

//==============================
// Main function implementations
//==============================
long long FindNewPrime()
{
	long long counter = v_primes.back();
	
	while( true){
		counter+=2;
		if(CheckIfIsNewPrime(counter) == true){
			return counter;
		}
	}
}

void AddNewPrime ( long long prime)
{
	v_primes.push_back(prime);
}

bool CheckIfIsNewPrime(long long number)
{
	std::vector<long long>::const_iterator it = v_primes.begin();
	std::vector<long long>::const_iterator it_end = v_primes.end();
	
	bool check_if_number_is_new_prime=false;
	long long remainder = 0;
	
	for(; it != it_end; ++it){
		remainder=number%(*it);
		if (remainder == 0) {
			check_if_number_is_new_prime = false;
			break;
		}
		else {
			check_if_number_is_new_prime = true;
		}
	}
return check_if_number_is_new_prime;
}
