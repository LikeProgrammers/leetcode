# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/zigzag-conversion/#/description
# Author: Wang Hao
# Date  : 2017-06-26

'''

ZigZag Conversion -- Medium

The string "PAYPALISHIRING" is written in a zigzag pattern
on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string
and make this conversion given a number of rows:

string convert(string text, int nRows);
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

'''

'''
0     6     C rList[0]
1   5 7   B   rList[1]
2 4   8 A     rList[2]
3     9       rList[nRows]
'''

class Solution(object):
    def convert(self, s, nRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if nRows <= 1:
            return s

        rList = ['' for i in xrange(nRows)]
        period = 2 * (nRows - 1)
        rIndex = {
            i: i if i < nRows else (period - i) for i in xrange(period)
        }

        for i in xrange(len(s)):
            rList[rIndex[i % period]] += s[i]

        return ''.join(rList)

    def convertTest(self):
        tDict = {
            1: {'in': ['PAYPALISHIRING', 3], 'out': 'PAHNAPLSIIGYIR'},
            2: {'in': ['0123456789ABC', 4], 'out': '06C157B248A39'},
        }

        for tNo, case in tDict.iteritems():
            print tNo,
            print 'pass' if case['out'] == \
                self.convert(case['in'][0], case['in'][1]) else 'fail'

Solution().convertTest()
