# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/median-of-two-sorted-arrays/#/description
# Author: Wang Hao
# Date  : 2017-06-24

'''

Median of Two Sorted Arrays -- hard

There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

'''

class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        i = 0
        ii = 0
        ret = None
        mList = []
        n1Len = len(nums1)
        n2Len = len(nums2)
        sLen = n1Len + n2Len
        if sLen <= 0:
            return None

        while len(mList) <= sLen // 2:
            if i < n1Len and (ii >= n2Len or nums1[i] <= nums2[ii]):
                mList.append(nums1[i])
                i += 1
            elif ii < n2Len and (i >= n1Len or nums1[i] >= nums2[ii]):
                mList.append(nums2[ii])
                ii += 1
            else :
                break

        ret = float(mList[-1]) if sLen % 2 != 0 else float(mList[-1] + mList[-2]) / 2
        return ret

    def findMedianSortedArraysTest(self):
        testDict = {
            2.0: [[1, 3], [2]],
            2.5: [[1, 2], [3, 4]],
            2.5: [[1, 3], [2, 4]],
            3.0: [[3, 4, 5], [0, 1]],
            0.0: [[0], [0]],
            10.0: [[10], [10]],
            None: [[],[]],
        }

        caseNo = 1
        for ret, case in testDict.iteritems():
            print caseNo,
            print 'pass' if ret == self.findMedianSortedArrays(case[0], case[1]) else 'fail'
            caseNo += 1

Solution().findMedianSortedArraysTest()
