package problem0207

func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish1(numCourses, prerequisites)
}

//解法1:拓扑排序,广度优先遍历
func canFinish1(numCourses int, prerequisites [][]int) bool {
	//入度表
	in := make([]int, numCourses)
	for _, pre := range prerequisites {
		in[pre[0]]++
	}
	// 队列
	queue := make([]int, 0)
	// 进入队列
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		//出队列
		node := queue[0]
		queue = queue[1:]
		//numCourses--
		//更新入度表
		for _, pre := range prerequisites {
			if pre[1] == node {
				in[pre[0]]--
				if in[pre[0]] == 0 {
					queue = append(queue, pre[0])
				}
			}
		}
	}
	return numCourses == 0
}

//解法二
func canFinish2(numCourses int, prerequisites [][]int) bool {
	req := make([]int, numCourses)
	taken := make([]bool, numCourses)
	postreq := make([][]int, numCourses)
	count := 0
	// 初始化 req postreq
	//req 先修课的数目
	//postreq 先修课的序号
	for _, p := range prerequisites {
		req[p[1]]++
		postreq[p[0]] = append(postreq[p[0]], p[1])
	}
	// 初始化req,cur
	cur := []int{}
	for k := range req {
		if req[k] == 0 {
			cur = append(cur, k)
			taken[k] = true
			count++
		}
	}

	for len(cur) > 0 {
		next := []int{}
		for _, c := range cur {
			for _, p := range postreq[c] {
				if !taken[p] {
					req[p]--
					if req[p] == 0 {
						next = append(next, p)
						taken[p] = true
						count++
					}
				}
			}
		}
		cur = next
	}

	return count == numCourses
}
