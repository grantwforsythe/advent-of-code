from load_data import load_data


def main():
    count = 0
    prev_sum = 0
    data = load_data()

    for idx, _ in enumerate(data):
        try:
            # get the three following values (the last value is exclusive)
            val = sum(data[idx:idx+3])
            if (val > prev_sum) and (prev_sum != 0):
                count +=1
        except IndexError:
            # index is out of range, there isn't anymore windows
            break
        prev_sum = val

    print(f'Answer: {count}')


if __name__ == '__main__':
    main()