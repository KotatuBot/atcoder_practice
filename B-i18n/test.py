
test = raw_input()

head = test[0]
tail = test[-1]
middle_len = str(len(test[1:-1]))

result = head +  middle_len + tail 

print(result)
