"""
Created on Wed Jul  26 17:10:18 2017

@author: Dmitry White
"""

# TODO:Given a list lst and a number N, create a new list that contains
# each number of lst at most N times without reordering.
# For example if N = 2, and the input is [1,2,3,1,2,1,2,3], you take [1,2,3,1,2],
# drop the next [1,2] since this would lead to 1 and 2 being in the result 3 times,
# and then take 3, which leads to [1,2,3,1,2,3].

def delete_nth(order,max_e):
    list = []
    for item in order:
        m = list.count(item)
        if m < max_e:
            list.append(item)
    return list
