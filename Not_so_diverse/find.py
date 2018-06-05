backet = {}
# get
first_input = input().split(" ")
N = int(first_input[0])
K = int(first_input[1])

data_input = input().split(" ")

# sort
data_input.sort()

# unique
unique = set(data_input)


# create backet
for n in unique:
    backet[n] = 0

# backet setting
for m in data_input:
    value = backet[m]
    value += 1
    # update
    backet[m] = value

backet_list = list(backet.values())
backet_list.sort()
backet_list.reverse()
if K < len(unique):
    ball_value = backet_list[K:]
    goukei = sum(ball_value)
    print(goukei)
else:
    print("0")

