"""
Created on Wed Aug 30 23:34:23 2017

@author: Dmitry White
"""

# TODO: You will be given a number and you will need to return it as a string in Expanded Form.
# For example:
# expanded_form(42), should return '40 + 2'
# expanded_form(70304), should return '70000 + 300 + 4'


def expanded_form(num):
    l = len(str(num))
    nulls = 10**(l-1)
    if l == 1:
        return str(num)
    num_new = (num//nulls)*nulls
    num_rest = num - num_new
    if num_rest == 0:
        return str(num_new)
    return str(num_new) + ' + ' + expanded_form(num_rest)