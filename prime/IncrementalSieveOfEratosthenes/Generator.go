package prime

// Generator is an excellent implementation using the Incremental Sieve of Eratosthenes method
// http://lukelafountaine.com/programming/math/2016/12/10/Incremental-Sieve.html
func Generator() <-chan int {

	out := make(chan int, 1)

	go func() {

		out <- 2
		num := 3
		composites := make(map[int][]int)

		for {
			if _, ok := composites[num]; !ok {
				out <- num
				composites[num*num] = []int{num}
			} else {
				for _, prime := range composites[num] {
					next := num + prime
					for next%2 == 0 {
						next += prime
					}
					if _, ok := composites[next]; ok {
						composites[next] = append(composites[next], prime)
					} else {
						composites[next] = []int{prime}
					}
				}
				delete(composites, num)
			}
			num += 2
		}
	}()

	return out

}
