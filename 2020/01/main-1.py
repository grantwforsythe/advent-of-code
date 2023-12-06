import sys


with open('inputs.txt', 'r') as f:
    lines = f.readlines()
    lines = list(map(int, lines))
    for first in lines:
        for second in lines:
            if first + second == 2020:
                print(first * second)
                sys.exit(0)                
