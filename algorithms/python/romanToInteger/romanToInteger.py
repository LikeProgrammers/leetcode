# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/roman-to-integer/#/description
# Author: Wang Hao
# Date  : 2017-07-05

'''

13. Roman to Integer -- Easy

Given a roman numeral, convert it to an integer.
Input is guaranteed to be within the range from 1 to 3999.

'''

class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        sv = {'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}

        ret = 0
        pren = 0
        for c in iter(s[::-1]):
            curn = sv[c]
            ret += -curn if curn < pren else curn
            pren = curn

        return ret

    def romanToIntTest(self):
        tDict = {
            1: {'i': 'CLXXXIX', 'o': 189},
            2: {'i': 'IX', 'o': 9},
        }

        for tNo, t in tDict.iteritems():
            print tNo,
            print 'ok' if t['o'] == self.romanToInt(t['i']) else 'fail'


if __name__ == '__main__':
    Solution().romanToIntTest()
