# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/regular-expression-matching/#/description
# Author: Wang Hao
# Date  : 2017-07-01

'''

10. Regular Expression Matching -- Hard

Implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true

'''

import re

class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        return True if re.findall('^' + p + '$', s) else False


    def isMatchTest(self):
        tDict = {
            1: {'in': ('aa', 'a'), 'out': False},
            2: {'in': ('aa', 'aa'), 'out': True},
            3: {'in': ('aaa', 'aa'), 'out': False},
            4: {'in': ('aa', 'a*'), 'out': True},
            5: {'in': ('aa', '.*'), 'out': True},
            6: {'in': ('ab', '.*'), 'out': True},
            7: {'in': ('aab', 'c*a*b'), 'out': True},
        }

        for tNo, case in tDict.iteritems():
            print tNo,
            print 'ok' if case['out'] == self.isMatch(case['in'][0], case['in'][1]) \
                    else 'fail'

if __name__ == '__main__':
    Solution().isMatchTest()
