"""
Created on Sat Jul  29 12:18:26 2017

@author: Dmitry White
"""

# TODO: Take 2 strings s1 and s2 including only letters from a to z.
# Return a new sorted string, the longest possible, containing distinct letters,
# each taken only once - coming from s1 or s2.

def longest(s1, s2):
    my_list = []
    for char in s1:
        if char not in my_list:
            my_list.append(char)
    for char in s2:
        if char not in my_list:
            my_list.append(char)
        my_list.sort()
    return "".join(my_list)