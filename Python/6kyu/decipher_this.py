"""
Created on Thu Sep 07 16:04:29 2017

@author: Dmitry White
"""

# TODO: Decipher this!
# You are given several secret messages you need to decipher. Here are the conditions:
#    The first letter corresponds to ASCII character code (case sensitive)
#    The second letter needs to be switched to the last letter
#    The last letter needs to be switched to the second letter
#    If it only has one letter, it will be unchanged
#    If it only has two letters, you will just need to convert the ASCII character code to a letter
#    Keeping it simple -- there are no special characters
# Example:
#   decipherThis('72olle 103doo 100ya'); // 'Hello good day'
#   decipherThis('82yade 115te 103o'); // 'Ready set go'


def decipher_this(string):
    l = string.split(' ')
    first = []
    second = []
    num, alpha, final = "", "", ""
    for word in l:
        for letter in word:
            if letter.isdigit():
                first.append(letter)
            if letter.isalpha():
                second.append(letter)
        num = chr(int("".join(first)))
        if second != []:
            if len(second) > 1:
                temp = second[0]
                second[0] = second[-1]
                second[-1] = temp
            alpha = "".join(second)
        final = final + num + alpha + ' '
        num = ''
        alpha = ''
        del first[:]
        del second[:]
    return final[0:len(final)-1]