"""
Created on Mon Sep 11 19:02:38 2017
@author: Dmitry White
"""

# TODO: You are given a string of n lines, each substring being n characters long:
# For example:
#   s = "abcd\nefgh\nijkl\nmnop"
# We will study some transformations of this square of strings.
#   Vertical mirror: vert_mirror (or vertMirror or vert-mirror)
#       vert_mirror(s) => "dcba\nhgfe\nlkji\nponm"
#   Horizontal mirror: hor_mirror (or horMirror or hor-mirror)
#       hor_mirror(s) => "mnop\nijkl\nefgh\nabcd"

def vert_mirror(strng):
    l = [(list(x)) for x in list(strng.split("\n"))]
    [x.reverse() for x in l]
    l = ["".join(x) for x in l]
    return "\n".join(l)


def hor_mirror(strng):
    l = list(strng.split("\n"))
    l.reverse()
    return "\n".join(l)


def oper(fct, s):
    return fct(s)