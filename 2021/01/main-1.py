from load_data import load_data


def main():
    count = 0
    prev_num = 0
    for num in load_data():
        if (num > prev_num) and (prev_num != 0):
            count += 1
        prev_num = num

    print(f'Answer: {count}')


if __name__ == '__main__':
    main()