package leetcoode

func canCompleteCircuit(gas []int, cost []int) int {
	var n = len(gas)
	for i := 0; i < n; {
		var totalGas, totalCost = 0, 0
		var j = i
		var cnt = 0
		for cnt < n {
			j = (i + cnt) % n
			totalGas += gas[j]
			totalCost += cost[j]
			if totalGas < totalCost {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		}
		i = i + cnt + 1
	}
	return -1
}
