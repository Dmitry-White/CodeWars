"""
Created on Tue Aug  08 23:30:45 2017

@author: Dmitry White
"""

# TODO: You need to return a string that displays a diamond shape on the screen using asterisk ("*") characters.
# Please see provided test cases for exact output format.
# The shape that will be returned from print method resembles a diamond, where the number provided as
# input represents the number of *’s printed on the middle line.
# The line above and below will be centered and will have 2 less *’s than the middle line.
# This reduction by 2 *’s for each line continues until a line
# with a single * is printed at the top and bottom of the figure. Return null if input is even number or negative
# (as it is not possible to print diamond with even number or negative number).


def diamond(n):
    if (n % 2 ==0) or (n < 0):
        return None
    s = []
    m = 1
    if n > 1:
        while m < n:
            for itme in range((n-m)//2):
                s.append(" ")
            for item in range(m):
                s.append("*")
            m += 2
            s.append("\n")
        while m >= 1:
            for itme in range((n-m)//2):
                s.append(" ")
            for item in range(m):
                s.append("*")
            s.append("\n")
            m-=2
        return "".join(s)
print(diamond(4))