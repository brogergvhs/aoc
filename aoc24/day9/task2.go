package main

func compressLayout2(origDisk []InfoPart) []InfoPart {
	var empties []struct {
		free  int
		index int
	}

	disk := make([]InfoPart, len(origDisk))
	copy(disk, origDisk)

	for i := 0; i < len(disk); i++ {
		if !disk[i].empty {
			continue
		}

		j := i + 1
		for ; j < len(disk); j++ {
			if !disk[j].empty {
				break
			}
		}

		empties = append(
			empties, struct {
				free  int
				index int
			}{free: j - i, index: i})
		i = j
	}

	r := len(disk) - 1
	for r >= 0 {
		if disk[r].empty || disk[r].isMoved {
			r--
			continue
		}

		j := r - 1
		for ; j >= 0; j-- {
			if disk[j].empty {
				break
			}
			if disk[j].id != disk[r].id {
				break
			}
		}
		length := r - j

		for pos, empty := range empties {
			if empty.index > r {
				break
			}

			if length <= empty.free {
				for i := 0; i < length; i++ {
					disk[empty.index+i] = disk[r-i]
					disk[empty.index+i].isMoved = true
					disk[r-i] = InfoPart{empty: true}
				}

				if length == empty.free {
					empties = append(empties[:pos], empties[pos+1:]...)
				} else {
					empties[pos].free = empty.free - length
					empties[pos].index += length
				}
				break
			}
		}

		r = r - length
	}

	return disk
}
