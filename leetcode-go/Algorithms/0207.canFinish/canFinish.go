package _207_canFinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	next := map[int][]int{}
	for _, v := range prerequisites {
		_, ok := next[v[1]]
		if ok {
			next[v[1]] = append(next[v[1]], v[0])
		} else {
			next[v[1]] = []int{v[0]}
		}
	}

	vis := make([]int, numCourses)
	for v := range next {
		if !checkLoop(vis, v, next) {
			return false
		}
	}
	return true
}

func checkLoop(visit []int, v int, mp map[int][]int) bool {
	if visit[v] == 1 {
		return false
	}
	if visit[v] == -1 {
		return true
	}
	for _, next := range mp[v] {
		if !checkLoop(visit, next, mp) {
			return false
		}
	}
	visit[v] = -1
	return true
}
