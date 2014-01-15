#include <iostream>
#include <vector>
#include <algorithm>

long long FindLargestPrimefactorOf(const long long number, bool & are_primes_suffient);
long long FindNewPrime();
long long FindNewFactor(const long long number, bool & are_primes_suffient);
bool CheckIfIsNewPrime(long long number);
void AddNewPrime ( long long prime);
void AddNewFactor ( long long factor);

void DisplayVector(const std::vector<long long> & v);


std::vector<long long> v_primes;
std::vector<long long> v_factors;

using namespace std;

int main()
{
//long long number = 600851475143LL;

long long number = 385;

cout << "Factorize "<< number << ":" << endl;


v_primes.push_back(2);
v_primes.push_back(3);
bool primes_are_suffient = true;
long long new_factor = number;

while ( new_factor > v_primes.back() ){
	new_factor = FindNewFactor(new_factor, primes_are_suffient);
	std::cout << "new factor: " << new_factor << std::endl;
	std::cout << "primes_are_suffient: " <<std::boolalpha<< primes_are_suffient << std::endl;	
	std::cout << "highes prime: " <<v_primes.back() << std::endl;	

	if ( primes_are_suffient == false){
	AddNewPrime(FindNewPrime());
	}

}

v_factors.push_back(v_primes.back());

 DisplayVector(v_factors);
 
 
 long long check = 1;
 
 std::vector<long long>::const_iterator citer;
	
	for (citer = v_factors.begin(); citer != v_factors.end(); citer++) 
	check*=(*citer);
	
		std::cout << "Check: " << check << std::endl;

  DisplayVector(v_primes);


}

//==============================
// Main function implementations
//==============================

long long FindNewFactor(const long long number, bool & primes_are_suffient)
{
	std::vector<long long>::const_iterator it = v_primes.begin();
	std::vector<long long>::const_iterator it_end = v_primes.end();	
		
	long long remainder = 0;
	long long quotient = 0;
	
	for(; it != it_end; ++it){
		remainder=number%(*it);
		if (remainder == 0) {
			AddNewFactor((*it));
			primes_are_suffient = true;
			return number/(*it);
		}
	}
	AddNewPrime(FindNewPrime());
	primes_are_suffient = false;
	return number;
}


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

void AddNewFactor ( long long factor)
{
	v_factors.push_back(factor);
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



// ===  FUNCTION  ==============================================================
//         Name:  DisplayVector
//  Description:  displays a vector of doubles
// =============================================================================
void DisplayVector(const std::vector<long long> & v)
{
	std::vector<long long>::const_iterator citer;
	
	for (citer = v.begin(); citer != v.end(); citer++) 
		std::cout << *citer << std::endl;
	
	std::cout << std::endl;
}   // -----  end of function DisplayVector  -----
		
		
