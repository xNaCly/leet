def reverse(x: int) -> int:
    # reverse string with string notation
    s = str(x)[::-1]
    # check if last character is -
    if s[-1] == "-":
        # "-" all chars in string but the last
        s = f"-{s[:-1]}"
    s = int(s)
    # return s if s is in the scope of a 32bit unsigned int
    return s if (-pow(2, 31)) < s < (pow(2, 31) - 1) else 0


print(reverse(-123))
