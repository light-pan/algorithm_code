package employeedispatch

func getMin(x, y, cntx, cnty int) int {

	// 至少有这么多编号的员工才能满足
	k := 1

	// 既不是x也不是y的倍数
	others := 0

	for cntx+cnty > others {
		modx := k % x
		mody := k % y

		if modx != 0 && mody != 0 {
			others++
		} else if modx != 0 {
			if cntx > 0 {
				cntx--
			}

		} else if mody != 0 {
			if cnty > 0 {
				cnty--
			}
		}
		k++
	}

	return k - 1
}
