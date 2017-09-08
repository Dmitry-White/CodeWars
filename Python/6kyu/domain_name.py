"""
Created on Fri Sep 08 13:00:08 2017

@author: Dmitry White
"""

# TODO:Write a function that when given a URL as a string,
# parses out just the domain name and returns it as a string.
# For example:
#   domain_name("http://github.com/carbonfive/raygun") == "github"
#   domain_name("http://www.zombie-bites.com") == "zombie-bites"
#   domain_name("https://www.cnet.com") == "cnet"

def domain_name(url):
    www = url.find("www")
    http = url.find("//")
    if www != -1:
        url = url[www+4:]
    else:
        url = url[http+2:]
    dot = url.find(".")
    return url[:dot]


