# If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

# Find the sum of all the multiples of 3 or 5 below 1000.

def main(multiple1, multiple2, upTo):
	# i can think of 2 ways to implement, run through each number upto 1000, checking modulus. O(upTo) time; 
	# OR while number less than upTo, add to dynamicArr. this is O(UpTO/multiple1 + UpTO/multiple2) time, which is quicker by a fraction of time.
	# I won't really be bothered to show its faster. less mem efficient

	sum = 0

	temp = multiple1

	while temp < upTo:
		sum += temp
		temp += multiple1

	temp = multiple2

	while temp < upTo:
		if not temp%multiple1 == 0:
			sum += temp
		temp += multiple2

	print (sum)


if __name__ == '__main__':
	main(3,5,1000)