from typing import Optional


class TreeNode:
    def __init__(self, val: int = 0, left: Optional['TreeNode'] = None, right: Optional['TreeNode'] = None):
        self.val: int = val
        self.left: Optional['TreeNode'] = left
        self.right: Optional['TreeNode'] = right

class BST:
    def __init__(self, root: Optional[TreeNode] = None):
        self.root: Optional[TreeNode] = root

    def get_root(self):
        return self.root

    def insert(self, val: int) -> None:
        if not self.root:
            self.root = TreeNode(val)
            return

        current: Optional[TreeNode] = self.root
        while True:
            if val < current.val:
                if current.left is None:
                    current.left = TreeNode(val)
                    return
                current = current.left
            elif val > current.val:
                if current.right is None:
                    current.right = TreeNode(val)
                    return
                current = current.right
            else:
                return  # Value already exists in the BST

    def find(self, val: int) -> Optional[TreeNode]:
        current: Optional[TreeNode] = self.root
        while current:
            if val < current.val:
                current = current.left
            elif val > current.val:
                current = current.right
            else:
                return current
        return None

    def inorder(self) -> list[int]:
        def _inorder(node: Optional[TreeNode], result: list[int]) -> None:
            if node:
                _inorder(node.left, result)
                result.append(node.val)
                _inorder(node.right, result)
        result = []
        _inorder(self.root, result)
        return result

    def display(self) -> None:
        def _display(node: Optional[TreeNode], indent: int) -> None:
            if node is None:
                return
            _display(node.right, indent + 4)
            print(" " * indent + str(node.val))
            _display(node.left, indent + 4)

        _display(self.root, 0)