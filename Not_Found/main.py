alphabet = list("abcdefghijklmnopqrstuvwxyz")
strings = raw_input()

input_data = list(strings)

input_data.sort()
input_dict = {}

for j in range(len(alphabet)):
    input_dict[alphabet[j]]=0


for i in range(len(input_data)):
    count = input_dict[input_data[i]]
    input_dict[input_data[i]] = count + 1


hit=0
for m in alphabet:
    data = input_dict[m]
    if data == 0:
        hit = m
        break

if hit==0:
    print("None")
else:
    print(hit)
