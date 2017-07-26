"""
Created on Wed Jul  26 17:23:57 2017

@author: Dmitry White
"""

# TODO:Given an integral number, determine if it's a square number.

def is_square(n):
    if n < 0: return False
    return pow(n, 0.5).is_integer()
