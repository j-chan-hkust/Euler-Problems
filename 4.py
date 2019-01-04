# A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

# Find the largest palindrome made from the product of two 3-digit numbers.

# Notes - ok i know that this is very very rough, but the simplicity of the problem means i barely need to optimize


def main():
	# to begin with, max length of the palindrome is 6, min length is 5
	# if we brute force from the direction of the palindrome, we will need to check 9*9*9*2 = 1458 options
	# for each palindrome, we will need to bruteforce check divisibility by at most 100-999 options
	# time complexity does not seem that bad.

	# greedy search

	k = 999
	while k>99:
		palindrome = str(k)+str(k)[::-1]

		palindromeNum = int(palindrome)

		for number in range(100,999):
			if palindromeNum%number == 0:
				if palindromeNum/number>99 and palindromeNum/number<1000:
					print(palindromeNum)
					return
		k -= 1

	k = 999
	while k>99:
		palindrome = str(k)+str(k)[0:1:-1]

		palindromeNum = int(palindrome)

		for number in range(100,999):
			if palindromeNum%number == 0:
				if palindromeNum/number>99 and palindromeNum/number<1000:
					print(palindromeNum)
					return
		k -= 1

	



if __name__ == '__main__':
	main()