"""
Created on Sat Sep 09 11:59:47 2017
@author: Dmitry White
"""

# TODO: You have an array of numbers.
# Your task is to sort ascending odd numbers but even numbers must be on their places.
# Zero isn't an odd number and you don't need to move it. If you have an empty array, you need to return it.


def sort_array(numbers):
    evens = []
    odds = []
    for a in numbers:
        if a % 2:
            odds.append(a)
            evens.append(None)
        else:
            evens.append(a)
    odds = iter(sorted(odds))
    return [next(odds) if b is None else b for b in evens]