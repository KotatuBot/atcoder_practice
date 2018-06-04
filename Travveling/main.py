N = raw_input()
norm = raw_input()
distance = [int(i) for i in norm.split(" ")]

distance.sort()
distance.reverse()

dist = distance[0] - distance[-1]

print(dist)

