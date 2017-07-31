"""
Created on Mon Jul  31 17:30:56 2017

@author: Dmitry White
"""

# TODO:A string is considered to be in title case if each word in the string is
# either (a) capitalised (that is, only the first letter of the word is in upper case)
# or (b) considered to be an exception and put entirely into lower case unless
# it is the first word, which is always capitalised.
# Write a function that will convert a string into title case,
# given an optional list of exceptions (minor words).
# The list of minor words will be given as a string with each word separated by a space.
# Your function should ignore the case of the minor words string --
# it should behave in the same way even if the case of the minor word string is changed.


def title_case(title, minor_words=""):
    start = False
    new_title = []
    title = title.split(" ")
    minor_words = minor_words.split(" ")
    minor_words = [x.lower() for x in minor_words]
    for word in title:
        word = word.lower()
        if start is False:
            new_title.append(word.capitalize())
            start = True
        elif word not in minor_words:
            new_title.append(word.capitalize())
        else:
            new_title.append(word)
    if all((x == "")or(x == " ") for x in new_title):
        del new_title[:]
    return " ".join(new_title)
