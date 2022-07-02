# Given an integer x, return true if x is palindrome integer.
# An integer is a palindrome when it reads the same backward as forward.
# For example, 121 is a palindrome while 123 is not.

from operator import truediv


number = 12344321

def isPalindromeSlow(x):
    print("x is: ", x)

    asStr = str(x)

    myLen = len(asStr)

    i = 0

    while i < myLen:
        j = (i + 1) * -1
        # print("i = ", i, "j = :", j)
        print(asStr[i], " ", asStr[j])

        if asStr[i] != asStr[j]:
            return False

        i += 1
    
    return True

# isPalindromeSlow(number)

def isPalindromeFast(x):
    print("x is: ", x)

    asStr = str(x)
    revStr = asStr[::-1]

    print(asStr)
    print(revStr)

    if asStr == revStr:
        return True
    else:
        return False
    print(asStr)
    print(revStr)

isPalindromeFast(number)
