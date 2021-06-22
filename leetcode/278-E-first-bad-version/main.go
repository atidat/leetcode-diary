package main

func main() {

}


/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	// 二分法
	for {
		f, l := 1, n
		t := (f+l)/2

		if isBadVersion(t)  {
			if !isBadVersion(t-1) {
				return t
			}
		}
		l = t
	}
	return -1
}

func isBadVersion(version int) bool {
	return true
}