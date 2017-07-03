# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/integer-to-roman/#/description
# Author: Wang Hao
# Date  : 2017-07-03

'''

12. Integer to Roman -- Mdeium

Given an integer, convert it to a roman numeral.
Input is guaranteed to be within the range from 1 to 3999.

'''

class Solution(object):
    def intToRoman(self, num):
        '''
        :type num: int
        :rtype: str
        '''
        value = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        symbol = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
        ret = []
        for v, s in zip(value, symbol):
            while num >= v:
                num -= v
                ret += s

        return ''.join(ret)

    def intToRomanTest(self):
        tDict = {
            1: {'in': 189, 'out': 'CLXXXIX'},
        }
        for tNo, case in tDict.iteritems():
            print tNo,
            print 'ok' if case['out'] == self.intToRoman(case['in']) else 'fail'

if __name__ == '__main__':
    Solution().intToRomanTest()
