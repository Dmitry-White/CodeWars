"""
Created on Wed Aug 09 23:45:41 2017

@author: Dmitry White
"""

# A pangram is a sentence that contains every single letter of the alphabet at least once.
# For example, the sentence "The quick brown fox jumps over the lazy dog" is a pangram,
# because it uses the letters A-Z at least once (case is irrelevant).
# Given a string, detect whether or not it is a pangram. Return True if it is,
# False if not. Ignore numbers and punctuation.

import string
def is_pangram(s):
    alpha = {}
    for item in list(string.ascii_lowercase):
        alpha[item] = 0
    for char in s.lower():
        if char in alpha:
            alpha[char] += 1
    return all(alpha[item] > 0 for item in alpha)
