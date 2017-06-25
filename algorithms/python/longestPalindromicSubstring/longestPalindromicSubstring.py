# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/longest-palindromic-substring/#/description
# Author: Wang Hao
# Date  : 2017-06-25

'''

Longest Palindromic Substring -- Medium

Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example:
Input: "cbbd"
Output: "bb"

'''

class Solution(object):
    def findPalindrome(self, s, lPaIndex, rPaIndex):
        rStr = ''
        sLen = len(s)
        while True:
            lPaIndex -= 1
            rPaIndex += 1
            if lPaIndex < 0 or rPaIndex >= sLen or s[lPaIndex] != s[rPaIndex]:
                break

        rStr = s[lPaIndex+1:rPaIndex]
        return rStr

    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        sLen = len(s)
        rStr=''
        for i in range(sLen):
            if i <= 0:
                rStr = s[0:1]
                continue

            if s[i] == s[i-1]:
                paStr = self.findPalindrome(s, i-1, i)
                rStr = paStr if len(paStr) > len(rStr) else rStr

            paStr = self.findPalindrome(s, i, i)
            rStr = paStr if len(paStr) > len(rStr) else rStr

        return rStr

    def longestPalindromeTest(self):
        testDict = {
            1: ['babad', 'bab'], # aba
            2: ['cbbd', 'bb'],
            3: ['', ''],
            4: ['ab', 'a'],
            5: ['a', 'a'],
        }

        for tNo, case in testDict.iteritems():
            print tNo,
            print 'pass' if case[1] == self.longestPalindrome(case[0]) else 'fail'

Solution().longestPalindromeTest()
