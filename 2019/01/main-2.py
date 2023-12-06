import math


def calculate_fuel(fuel):
    new_fuel = math.floor(fuel / 3) - 2 
    if new_fuel > 0:
        return new_fuel + calculate_fuel(new_fuel)
    else:
        return 0

with open('inputs.txt', 'r') as f:
    modules = list(map(int, f.readlines()))
    total_fuel = 0

    for module in modules: 
        total_fuel += calculate_fuel(module)
    
    print(total_fuel)