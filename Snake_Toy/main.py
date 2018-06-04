import sys

def main():
    sums = 0
    raw_inputer = raw_input()
    raw_inputer_test = raw_inputer.split(" ")

    N = int(raw_inputer_test[0])
    K = int(raw_inputer_test[1])


    lenter = raw_input()
    lenter_test = lenter.split(" ")

    if N!=len(lenter_test):
        sys.exit()

    if N < K:
       sys.exit() 
    elif K < 1:
        sys.exit()
    elif 50 < N:
        sys.exit()
    else:
        pass

    for j in lenter_test:
        if j < 1 and 50 <j:
            sys.exit()
            break

    lenter_test.sort()
    lenter_test.reverse()

    for i in range(K):
        sums += int(lenter_test[i])


    print(sums)

if __name__=="__main__":
    main()
