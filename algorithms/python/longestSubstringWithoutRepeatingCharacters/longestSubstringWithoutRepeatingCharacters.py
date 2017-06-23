# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description
# Author: Wang Hao
# Date  : 2017-06-23

'''

Longest Substring Without Repeating Characters -- medium

Given a string,
find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3.
Note that the answer must be a substring,
"pwke" is a subsequence and not a substring.

'''

class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        subIndex, subLen, maxLen = 0, 0, 0

        for i, c in enumerate(s):
            curSub = s[subIndex:i]
            if c in curSub:
                maxLen = max(subLen, maxLen)
                subIndex += curSub.index(c) + 1
                subLen = i - subIndex + 1
            else:
                subLen += 1

        return max(maxLen, subLen)

    def lengthOfLongestSubstringTest(self):
        print '1', "pass" if 3 == self.lengthOfLongestSubstring('abcabcbb') else "fail"
        print '2', "pass" if 1 == self.lengthOfLongestSubstring('bbbbb') else 'fail'
        print '3', "pass" if 3 == self.lengthOfLongestSubstring('pwwkew') else 'fail'
        print '4', 'pass' if 0 == self.lengthOfLongestSubstring('') else 'fail'
        print '5', 'pass' if 1 == self.lengthOfLongestSubstring('a') else 'fail'
        print '6', 'pass' if 3 == self.lengthOfLongestSubstring('dvdf') else 'fail'

Solution().lengthOfLongestSubstringTest()
