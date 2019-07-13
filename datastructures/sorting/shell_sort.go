package sorting


func ShellSort(a []int) {
	n := len(a)

	h := 1
	for h<n/3 {
		h = 3*h +1 
	}

	for h >=1 {
		for i:=h; i<n; i++ {
			for j :=i ;j>=h && a[j] < a[j-h] ;j-=h {
				swap(a, j, j-h)
			}
		}
		h = h/3
	}
	

}