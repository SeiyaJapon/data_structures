from typing import Optional

from python.min_absolute_diff_bst.bst import TreeNode, BST


class Solution:
    def get_minimum_difference(self, bst: BST) -> int:
        root = bst.get_root()
        return self._get_min_diff(root, root.val, -1)

    def _get_min_diff(self, node: Optional[TreeNode], prev_val: int, min_val: int) -> int:
        if node.left is None and node.right is None:
            return min(min_val, abs(node.val - prev_val)) if min_val != -1 else abs(node.val - prev_val)

        min_left = None
        if node.left is not None:
            min_left = self._get_min_diff(node.left, prev_val, min_val)

        min_right = None
        if node.right is not None:
            min_right = self._get_min_diff(node.right, node.val, min_val)

        return min(min_val, min_left, min_right) if min_val != -1 else min(min_left, min_right)