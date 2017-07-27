"""
Created on Thu Jul  27 15:54:33 2017

@author: Dmitry White
"""

# TODO: Given a string, you have to return a string in which each character (case-sensitive) is repeated once.

def double_char(s):
    list = []
    for char in s:
        list.append(char*2)
    return "".join(list)
