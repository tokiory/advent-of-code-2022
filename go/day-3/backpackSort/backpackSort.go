package backpackSort

type RuneBackpackPart []rune

type RuneBackpack [2]RuneBackpackPart

func (b *RuneBackpackPart) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b RuneBackpackPart) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b RuneBackpackPart) Greater(i, j int) bool {
	return b[i] > b[j]
}

func (b RuneBackpackPart) Len() int {
	return len(b)
}
