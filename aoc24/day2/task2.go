package main

func canBecomeSafeByRemovingOne(report []int) bool {
	for i := 0; i < len(report); i++ {

		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		if isSafe(newReport) {
			return true
		}
	}
	return false
}
