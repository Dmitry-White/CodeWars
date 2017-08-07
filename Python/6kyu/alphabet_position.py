"""
Created on Sat Aug  05 23:59:07 2017

@author: Dmitry White
"""

# TODO: Given a string, replace every letter with its position in the alphabet.
# If anything in the text isn't a letter, ignore it and don't return it.a being 1, b being 2, etc.

from string import ascii_lowercase


def alphabet_position(text):
    return ' '.join(str(ascii_lowercase.index(n.lower()) + 1) for n in text if n.isalpha())
