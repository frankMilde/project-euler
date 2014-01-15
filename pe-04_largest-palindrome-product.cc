#include <iostream>

bool CheckNumberForPalidrom( long long number );


using namespace std;

int main()
{

	long long product					= 0;
	bool      number_is_palindrom		= false;
	long long max_palindrom				= 0;
	int		  max_three_digit_number	= 999;

	for (int number_1 = max_three_digit_number; number_1 != 0; --number_1){
		for (int number_2 = number_1; number_2 != 0; --number_2){
		
			product				= number_1 * number_2;
			number_is_palindrom = CheckNumberForPalidrom( product );
			
			if (	number_is_palindrom == true 
				&&	product > max_palindrom){
				cout << number_1 << " x " << number_2 <<
				" = " << product << " is palindrom: " << 
				std::boolalpha << CheckNumberForPalidrom( product ) 
				<< endl;
				max_palindrom = product;
			}
		}
	}

}

bool CheckNumberForPalidrom( long long number )
{
	int ones						= 0;

	int tens						= 0;
	int tens_remainder				= 0;

	int hundreds					= 0;
	int hundreds_remainder			= 0;

	int thousands					= 0;
	int thousands_remainder			= 0;

	int ten_thousands				= 0;
	int ten_thousands_remainder		= 0;

	int hundred_thousands			= 0;
	int hundred_thousands_remainder = 0;


	hundred_thousands			= number/100000;
	hundred_thousands_remainder = number%100000;
	
	ten_thousands			= (hundred_thousands_remainder)/10000;
	ten_thousands_remainder = (hundred_thousands_remainder)%10000;
	
	thousands				= (ten_thousands_remainder)/1000;
	thousands_remainder		= (ten_thousands_remainder)%1000;
	
	hundreds				= (thousands_remainder)/100;
	hundreds_remainder		= (thousands_remainder)%100;
	
	tens					= (hundreds_remainder)/10;
	tens_remainder			= (hundreds_remainder)%10;
	
	ones					= (tens_remainder);
	
	//cout << "Number: " << number << endl;
	//cout << "        " << hundred_thousands << ten_thousands << thousands << hundreds << tens << ones << endl;
	
	if ((ones==hundred_thousands) 
		&& (tens == ten_thousands)
		&& (hundreds == thousands))
		return true;
	else
		return false;
}