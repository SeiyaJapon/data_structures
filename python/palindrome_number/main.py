# Palindrome Number
#
# Given an integer x, return true if x is a palindrome, and false otherwise.
#
#
#
# Example 1:
#
# Input: x = 121
# Output: true
# Explanation: 121 reads as 121 from left to right and from right to left.
# Example 2:
#
# Input: x = -121
# Output: false
# Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
# Example 3:
#
# Input: x = 10
# Output: false
# Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
#
#
# Constraints:
#
# -231 <= x <= 231 - 1

if __name__ == "__main__":
    def num_pal(x: int) -> bool:
        if x < 0:
            return False

        original = x
        reverse = 0
        while x > 0:
            digit = x % 10
            reverse = reverse * 10 + digit
            x //= 10

        return original == reverse

    print("121", num_pal(121))  # True
    print("-121", num_pal(-121))  # False
    print("10", num_pal(10))  # False
    print("12321", num_pal(12321))  # True
    print("0", num_pal(0))  # True
    print("123456", num_pal(123456))  # False
    print("1221", num_pal(1221))  # True
    print("1001", num_pal(1001))  # True
    print("1234321", num_pal(1234321))  # True
    print("1234567899", num_pal(1234567899))  # False