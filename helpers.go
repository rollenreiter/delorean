package main

// Basic Bubble Sort for token types
func Sort(u []token, alphabetical bool) {
	swap := func(u []token, p1 int, p2 int) {
		temp := u[p2]
		u[p2] = u[p1]
		u[p1] = temp
	}
	swapped := true
	for swapped {
		swapped = false
		for i := range u {
			switch {
			case i == 0:
				continue
			case !alphabetical:
				{
					if u[i-1].order > u[i].order {
						swap(u, i-1, i)
						swapped = true
					}
				}
			default:
				{
					if u[i-1].content[0] > u[i].content[0] {
						swap(u, i-1, i)
						swapped = true

					}
				}
			}
		}
	}
}
