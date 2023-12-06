with open('inputs.txt', 'r') as f:
    lines = f.readlines()
    num_valid = 0

    for line in lines:
        index, character, password = tuple(line.split(' '))

        first_position, second_position = tuple(map(int, index.split('-')))
        character = character.replace(':', '')
        password = password.strip()

        # a valid password can contain the character in one of the positions 
        if (password[first_position-1] == character) != (password[second_position-1] == character):
            num_valid += 1 

    print(num_valid)