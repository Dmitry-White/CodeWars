"""
Created on Sun Jul  10 15:40:11 2017

@author: Dmitry White
"""

# TODO: Check to see if a string has the same amount of 'x's and 'o's.
# The method must return a boolean and be case insensitive. 
# The string can contains any char.

def xo(s):
    i = 0
    j = 0
    for char in s.lower():
        if char == 'x':
            i += 1
        if char == 'o':
            j += 1
    return i == j