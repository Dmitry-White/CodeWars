"""
Created on Thu Jul  27 16:03:02 2017

@author: Dmitry White
"""

# TODO: Return the number (count) of vowels in the given string.
# Consider a, e, i, o, and u as vowels for this Kata.

def getCount(inputStr):
    num_vowels = 0
    vowels = ['a','e','i','o','u']
    for char in inputStr:
        if char in vowels: num_vowels +=1
    return num_vowels
