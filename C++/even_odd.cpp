/*
Created on Wed Aug  23 23:18:11 2017

@author: Dmitry White
*/

/*
    TODO: Create a function that takes an integer as an argument and
    returns "Even" for even numbers or "Odd" for odd numbers.
*/


std::string even_or_odd(int number)
{
  if(number%2 == 0)
  {
    return "Even";
  }
  else
  {
    return "Odd";
  }
}