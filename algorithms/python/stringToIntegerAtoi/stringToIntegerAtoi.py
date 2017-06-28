# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/string-to-integer-atoi/#/description
# Author: Wang Hao
# Date  : 2017-06-28

'''

String to Integer (atoi) -- Medium

Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases.
If you want a challenge, please do not see below and
ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely
(ie, no given input specs).
You are responsible to gather all the input requirements up front.

Update (2015-02-10):
The signature of the C++ function had been updated.
If you still see your function signature accepts a const char * argument,
please click the reload button  to reset your code definition.

spoilers alert... click to show requirements for atoi.

Requirements for atoi:
The function first discards as many whitespace characters as necessary
until the first non-whitespace character is found.
Then, starting from this character,
takes an optional initial plus or minus sign followed
by as many numerical digits as possible,
and interprets them as a numerical value.

The string can contain additional characters after those
that form the integral number,
which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str
is not a valid integral number,
or if no such sequence exists because either str is empty
or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.
If the correct value is out of the range of representable values,
INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

'''

class Solution(object):
    INT_MAX = pow(2, 31) - 1
    INT_MIN = -pow(2, 31)

    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        if str == None or str == '':
            return 0

        for i, c in enumerate(str):
            if c.isdigit() or c == '+' or c == '-':
                str = str[i:]
                break
            if not c.isspace():
                return 0

        s = -1 if str[0] == '-' else 1
        str = str[1:] if str[0] == '+' or str[0] == '-' else str

        for i, c in enumerate(str):
            if not c.isdigit():
                str = str[:i]
                break

        if not str.isdigit():
            return 0

        num = int(str)
        num *= s
        if num < self.INT_MIN:
            ret = self.INT_MIN
        elif num > self.INT_MAX:
            ret = self.INT_MAX
        else:
            ret = num
        return ret

    def myAtoiTest(self):
        tDict = {
            1: {'in': "0001", 'out': 1},
            2: {'in': '-001', 'out': -1},
            3: {'in': '+002', 'out': 2},
            4: {'in': '9000000000', 'out': 2147483647},
            5: {'in': '', 'out': 0},
            6: {'in': '+023', 'out': 23},
            7: {'in': '   0023   ', 'out': 23},
            8: {'in': '+', 'out': 0},
            9: {'in': '+-2', 'out': 0},
            10: {'in': '2+', 'out': 2},
            11: {'in': "  -0012a42", 'out': -12},
            12: {'in': "2147483648", 'out': 2147483647},
            13: {'in': " b11228552307", 'out': 0},
        }

        for tNo, case in tDict.iteritems():
            print tNo,
            print 'ok' if case['out'] == self.myAtoi(case['in']) else 'fail'

def main():
    Solution().myAtoiTest()

if __name__ == '__main__':
    main()
