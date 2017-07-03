# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/container-with-most-water/#/description
# Author: Wang Hao
# Date  : 2017-07-02

'''

11. Container With Most Water -- Mdeium

Given n non-negative integers a1, a2, ..., an,
where each represents a point at coordinate (i, ai).
n vertical lines are drawn such
that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

'''

class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        left, right, maxArea = 0, len(height) - 1, 0

        while left < right:
            minHeight = min(height[left], height[right])
            maxArea = max(maxArea, (right - left) * minHeight)

            while left < right and height[left] <= minHeight:
                left += 1
            while left < right and height[right] <= minHeight:
                right -= 1

        # while left < right:
        #     maxArea = max(maxArea, (right - left) * min(height[left], height[right]))
        #     if height[left] < height[right]:
        #         while left < right:
        #             left += 1
        #             if height[left-1] < height[left]:
        #                 break
        #     else:
        #         while left < right:
        #             right -= 1
        #             if height[right+1] < height[right]:
        #                 break
        return maxArea

    def maxAreaTest(self):
        tDict = {
            1: {'in': [10, 20, 5, 30, 2], 'out': 40},
        }

        for tNo, case in tDict.iteritems():
            print tNo,
            print 'ok' if case['out'] == self.maxArea(case['in']) else 'fail'

if __name__ == '__main__':
    Solution().maxAreaTest()


