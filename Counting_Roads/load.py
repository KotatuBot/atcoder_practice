
lister = []
backet = {}

first_input = input().split(" ")

N = int(first_input[0])
M = int(first_input[1])

# get input
for j in range(M):
    two_input = input().split(" ")
    lister.append(int(two_input[0]))
    lister.append(int(two_input[1]))

# create key
order = range(1,N+1,1)


# create dict
for j in order:
    backet[j]=0

# setting 
for i in lister:
    value = backet[i]
    value += 1
    # update
    backet[i] = value


# printing
for m in order:
    print(backet[m])

