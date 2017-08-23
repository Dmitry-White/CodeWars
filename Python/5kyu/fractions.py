"""
Created on Wed Aug  23 23:24:01 2017

@author: Dmitry White
"""

# TODO: Given a string representing a simple fraction x/y, your function must return a string
# representing the corresponding mixed fraction in the following format:
#    [sign]a b/c
# where a is integer part and b/c is irreducible proper fraction.
# There must be exactly one space between a and b/c.
# Provide [sign] only if negative (and non zero) and only at the beginning of the number (both integer part
# and fractional part must be provided absolute).
# If the x/y equals the integer part, return integer part only.
# If integer part is zero, return the irreducible proper fraction only.
# In both of these cases, the resulting string must not contain any spaces.
# Division by zero should raise an error (preferably, the standard zero division error of your language).


def gcdRecur(a, b):
    if b == 0:
        return a
    else:
        return gcdRecur(b, a % b)


def mixed_fraction(s):
    l = s.split("/")
    num = int(l[0])
    denum = int(l[1])
    frac = []
    frac_sign = []

    if denum == 0:
        raise ZeroDivisionError

    if num < 0 and denum < 0:
        num = abs(num)
        denum = abs(denum)
    elif num < 0 or denum < 0:
        num = abs(num)
        denum = abs(denum)
        frac_sign.append('-')

    print(num, denum)
    whole = num // denum
    print(whole)
    new_num = num - whole * denum
    div = gcdRecur(new_num, denum)

    if new_num == 0 or new_num == denum:
        if whole != 0:
            frac_sign.append(str(whole))
            ans = "".join(frac_sign)
            return ans
        else:
            return str(whole)

    frac_temp = [str(new_num // div), str(denum // div)]
    print(whole)
    if whole != 0:
        frac.append(str(whole))
    print(frac)
    frac.append("/".join(frac_temp))
    digits = " ".join(frac)
    frac_sign.append(digits)
    ans = "".join(frac_sign)
    return ans