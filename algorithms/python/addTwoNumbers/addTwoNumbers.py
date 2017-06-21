# -*- coding: utf-8 -*-

# Source: https://leetcode.com/problems/add-two-numbers/#/description
# Author: Wang Hao
# Date  : 2017-06-21

'''

Add Two Numbers -- Medium

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero,
except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8

'''

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        carry = 0
        retList = ListNode(0)
        curNode = retList
        list1 = l1
        list2 = l2

        while list1 or list2 or carry:
            add = (list1.val if list1 else 0) + (list2.val if list2 else 0) + carry
            carry = add / 10
            curNode.next = ListNode(add % 10)

            curNode = curNode.next
            list1 = list1.next if list1 else None
            list2 = list2.next if list2 else None

        return retList.next

    def createListByTuple(self, t):
        retList = ListNode(0)
        curNode = retList

        for val in t:
            curNode.next = ListNode(val)
            curNode = curNode.next

        return retList.next

    def compList(self, l1, l2):
        list1 = l1
        list2 = l2

        while list1 or list2:
            val1 = list1.val
            val2 = list2.val

            list1 = list1.next if list1 else -1
            list2 = list2.next if list2 else -1
            if val1 != val2:
                return False

        return True

    def printList(self, l):
        n = l
        while n:
            print n.val
            n = n.next if n else None

    def addTwoNumbersTest(self):
        l1 = self.createListByTuple((2, 4, 3))
        l2 = self.createListByTuple((5, 6, 4))
        rl = self.createListByTuple((7, 0, 8))

        ret = self.addTwoNumbers(l1, l2)
        if self.compList(ret, rl):
            print "pass"
        else:
            print "fail"

Solution().addTwoNumbersTest()
