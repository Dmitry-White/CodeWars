"""
Created on Fri Nov  03 21:34:48 2017

@author: Dmitry White
"""

# TODO:The goal of this exercise is to convert a string to a new string
# where each character in the new string is '(' if that character appears
# only once in the original string, or ')' if that character appears
# more than once in the original string.
# Ignore capitalization when determining if a character is a duplicate.


from collections import Counter
def duplicate_encode(word):
    l = []
    s = ""
    word = word.lower()
    for letter in word:
        l.append(word.count(letter))
    for x in l:
        if x == 1:
            s += '('
        else:
            s += ')'
    return s
