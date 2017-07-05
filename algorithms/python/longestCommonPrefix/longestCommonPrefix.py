# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/longest-common-prefix/#/description
# Author: Wang Hao
# Date  : 2017-07-05

'''

14. Longest Common Prefix -- Easy

Write a function to find the longest common prefix string
amongst an array of strings.

'''

class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        r = ''

        for cList in zip(*strs):
            c = cList[0]
            if cList.count(c) == len(cList):
                r += c
            else:
                break

        return r

    def longestCommonPrefixTest(self):
        tDict = {
            1: {'i': [], 'o': ''},
            2: {'i': ['abcwang', 'abchao'], 'o': 'abc'},
            3: {'i': ['hao', 'haok', 'haokw'], 'o': 'hao'},
        }

        for tNo, t in tDict.iteritems():
            print tNo,
            print 'ok' if t['o'] == self.longestCommonPrefix(t['i']) else 'fail'


if __name__ == '__main__':
    Solution().longestCommonPrefixTest()
