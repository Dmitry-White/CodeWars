"""
Created on Sat Jul  29 12:18:26 2017

@author: Dmitry White
"""

# TODO: Your task is to make a function that can take any non-negative integer as a argument
# and return it with it's digits in descending order. Essentially, rearrange the digits to create the highest possible number.

def Descending_Order(num):
    return int("".join(map(str,list(reversed(sorted(list(map(int,str(num)))))))))
