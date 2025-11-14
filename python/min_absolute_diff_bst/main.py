from python.min_absolute_diff_bst.bst import BST
from python.min_absolute_diff_bst.solution import Solution


def main() -> None:
    # arr = [4, 2, 6, 1, 3]
    arr = [7, 0, 48, 12, 79]
    bst = BST()
    for v in arr:
        bst.insert(v)

    solution = Solution()
    result = solution.get_minimum_difference(bst)
    print("\nMinimum Absolute Difference:", result)


if __name__ == "__main__":
    main()
