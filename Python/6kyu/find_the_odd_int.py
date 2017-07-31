"""
Created on Tue Jul  25 17:15:26 2017

@author: Dmitry White
"""

# TODO:Given an array, find the int that appears an odd number of times.
# There will always be only one integer that appears an odd number of times.


def find_it(seq):
    for item in seq:
        times = 0
        for check in seq:
            if check == item:
                times += 1
        if times % 2 != 0:
            break
    return item
