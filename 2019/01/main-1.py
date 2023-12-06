import math

with open('inputs.txt', 'r') as f:
    modules = list(map(int,f.readlines()))
    total_fuel = 0

    for module in modules:
        total_fuel += math.floor(module / 3) - 2 
    
    print(total_fuel)