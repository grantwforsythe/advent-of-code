from typing import List


def load_data() -> List:
    """ 
    load the input data
    """
    with open('input.txt', 'r') as f:
        lines = [line.strip() for line in f.readlines()]
    return lines