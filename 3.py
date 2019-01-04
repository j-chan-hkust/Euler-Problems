# The prime factors of 13195 are 5, 7, 13 and 29.

# What is the largest prime factor of the number 600851475143 ?


def main(n):
	# slowly work up the primes to find the factors of n
	# begin with base knowledge. if a number is not divisible by smaller primes, it is prime.
	# from this we can assemble larger and larger primes
	# n will continuously be divisible by these larger and larger primes
	# eventually we will factor out all primes.

	primes = []
	primes.append(2)

	currNumber = 3

	temp = n

	#first iteration

	while temp%primes[-1] == 0:
		temp /= primes[-1]


	k = 0
	while not temp == 1: # lets remember to add an exit case so i dont fry this computer
		if k>n:
			break
		#find the next prime
		isPrime = True

		for prime in primes:
			if currNumber%prime == 0:
				isPrime = False

		if isPrime:
			while temp%currNumber == 0:
				temp /= currNumber
			primes.append(currNumber)
		currNumber +=1

	print(primes[-1])



if __name__ == '__main__':
	main(600851475143)