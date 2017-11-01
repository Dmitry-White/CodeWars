"""
Created on Wed Nov 01 21:02:41 2017

@author: Dmitry White
"""

# TODO:Given two strings s1 and s2, we want to visualize how different the two
# strings are. We will only take into account the lowercase letters (a to z).
# First let us count the frequency of each lowercase letters in s1 and s2.
# s1 = "A aaaa bb c"
# s2 = "& aaa bbb c d"
# s1 has 4 'a', 2 'b', 1 'c'
# s2 has 3 'a', 3 'b', 1 'c', 1 'd'
# So the maximum for 'a' in s1 and s2 is 4 from s1; the maximum for 'b' is 3 from s2.
# In the following we will not consider letters when the maximum of their
# occurrences is less than or equal to 1.
# We can resume the differences between s1 and s2 in the following string:
# "1:aaaa/2:bbb" where 1 in 1:aaaa stands for string s1 and aaaa because
# the maximum for a is 4. In the same manner 2:bbb stands for string s2 and
# bbb because the maximum for b is 3.
# The task is to produce a string in which each lowercase letters of s1 or s2
# appears as many times as its maximum if this maximum is strictly greater than 1;
# these letters will be prefixed by the number of the string
# where they appear with their maximum value and :.
# If the maximum is in s1 as well as in s2 the prefix is =:.
# In the result, substrings (a substring is for example 2:nnnnn or 1:hhh;
# it contains the prefix) will be in decreasing order of their length and
#  when they have the same length sorted in ascending lexicographic order
# (letters and digits - more precisely sorted by codepoint);
# the different groups will be separated by '/'. See examples and "Example Tests".
# s1 = "my&friend&Paul has heavy hats! &"
# s2 = "my friend John has many many friends &"
# mix(s1, s2) --> "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

def mix(s1,s2):
	alpha = 'abcdefghijklmnopqrstuvwxyz'
	count1 = []
	count2 = []
	max1 = []
	max2 = []
	equal = []
	form1 = []
	form2 = []
	forme = []

	for letter in alpha:
		lets = list(filter(lambda x: x == letter, s1))
		count1.append(lets)
	for letter in alpha:
		lets = list(filter(lambda x: x == letter, s2))
		count2.append(lets)

	for i in range(0,len(count1)):
		if len(count1[i]) == len(count2[i]) and len(count1[i]) > 1:
			equal.append(count1[i])
		elif len(count1[i]) > len(count2[i]) and len(count1[i]) > 1:
			max1.append(count1[i])
		elif len(count1[i]) < len(count2[i]) and len(count2[i]) > 1:
			max2.append(count2[i])

	for items in max1:
		form1.append('1:' + items[0]*len(items))
	for items in max2:
		form2.append('2:' + items[0]*len(items))
	for items in equal:
		forme.append('=:' + items[0]*len(items))

	final = form1 + form2 + forme
	final.sort(key=len, reverse=True)
	return '/'.join(final)
