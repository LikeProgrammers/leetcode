# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/reverse-integer/#/description
# Author: Wang Hao
# Date  : 2017-06-27

'''

Reverse Integer -- Easy

Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

    Have you thought about this?
    Here are some good questions to ask before coding.
    Bonus points for you if you have already thought through this!

    If the integer's last digit is 0, what should the output be?
    ie, cases such as 10, 100.

    Did you notice that the reversed integer might overflow?
    Assume the input is a 32-bit integer,
    then the reverse of 1000000003 overflows.
    How should you handle such cases?

    For the purpose of this problem,
    assume that your function returns 0
    when the reversed integer overflows.

    Note:
    The input is assumed to be a 32-bit signed integer.
    Your function should return 0 when the reversed integer overflows.

'''

class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        s = cmp(x, 0)
        r = int(repr(s*x)[::-1])
        return s*r * (r < pow(2, 31) - (s==1))

    def reverseTest(self):
        tDict = {
            1: {'in': -123, 'out': -321},
            2: {'in': 123, 'out': 321},
            3: {'in': 100, 'out': 1},
            4: {'in': 1000000003, 'out': 0},
        }

        for tNo, case in tDict.iteritems():
            print tNo,
            print 'pass' if case['out'] == self.reverse(case['in']) else 'fail'

Solution().reverseTest()
