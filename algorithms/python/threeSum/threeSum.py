# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/3sum/#/description
# Author: Wang Hao
# Date  : 2017-07-12

'''

15. 3Sum -- Mdeium

Given an array S of n integers,
are there elements a, b, c in S such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

'''

class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        for i in xrange(len(nums)-2):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            l, r = i+1, len(nums)-1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if s < 0:
                    l +=1
                elif s > 0:
                    r -= 1
                else:
                    res.append([nums[i], nums[l], nums[r]])
                    l += 1; r -= 1
                    while l < r and nums[l] == nums[l-1]:
                        l += 1
                    while l < r and nums[r] == nums[r+1]:
                        r -= 1
        return res

    def threeSumTest(self):
        tDict = {
            1: {'i': [1, 0, -1, -2, 2], 'o': [[-2, 0, 2], [-1, 0, 1]]},
            2: {'i': [3, 2, 1], 'o': []},
            3: {'i': [-1, 0, 1, 2, -1, -4], 'o': [[-1, -1, 2], [-1, 0, 1]]}
        }

        for tNo, t in tDict.iteritems():
            print tNo,
            r = self.threeSum(t['i'])
            print r
            print 'ok' if t['o'] == r else 'fail'


if __name__ == '__main__':
    Solution().threeSumTest()
