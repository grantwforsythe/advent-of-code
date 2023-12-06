from load_data import load_data


def main():
    # starting positions
    horizontal = 0
    depth = 0

    for line in load_data():
        val = int(line[-1])

        match line[0]:
            # forward
            case 'f':
                horizontal += val
            # down
            case 'd':
                depth += val
            # up
            case 'u':
                depth -= val
            case _:
                raise ValueError("There is more than 3 positions!")

    print(f'Answer: {horizontal * depth}')


if __name__ == '__main__':
    main()