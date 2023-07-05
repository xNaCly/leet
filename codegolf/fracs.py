import sys;import math as m
for a in sys.argv[1:]:a=[*map(int, a.split("/"))];a=[i//m.gcd(a[0],a[1]) for i in a];print(f"{a[0]}/{a[1]}")
