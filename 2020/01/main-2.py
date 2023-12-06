import sys


with open('inputs.txt', 'r') as f:
    lines = f.readlines()
    lines = list(map(int, lines))
    for first in lines:
        for second in lines:
            for third in lines:
                if first + second + third == 2020:
                    print(first * second * third)
                    sys.exit(0)
