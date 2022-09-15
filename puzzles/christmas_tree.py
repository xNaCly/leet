# task:
# should print a christmas tree like so:
#! n = 5 (amount of rows)
#! solution should be O(n)
#    * 1
#   *** 3
#  ***** 5
# ******* 7
#    *
def chris_tree(n: int) -> None:
  stars = 1
  max_stars = (n*2)-1
  for x in range(1,n+1):
    left_space = 1
    calc_left_space = (max_stars-stars)//2

    if x == n:
      left_space = 0
    else:
      left_space = calc_left_space


    print((" "*left_space)+("*"*stars))
    stars += 2
  print(" "*(n-1)+"*")

chris_tree(25)