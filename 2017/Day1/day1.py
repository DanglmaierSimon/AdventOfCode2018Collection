#!/usr/bin/env python3

import sys

input = sys.argv[1]

sum = 0

for idx in range(len(input)):
	if input[idx] == input[(idx+1)%len(input)]:
		sum += input[idx]-48

print(sum)