"""
Created on Mon Aug  05 23:51:03 2017

@author: Dmitry White
"""

# TODO: Given two arrays a and b write a function comp(a, b)
# that checks whether the two arrays have the "same" elements,
# with the same multiplicities. "Same" means, here, that the elements
# in b are the elements in a squared, regardless of the order.


def comp(array1, array2):
    try:
        for item in array2:
            if item**0.5 not in array1:
                return False
        return True
    except:
        return False
