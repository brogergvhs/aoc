package main

func compressLayout1(origDisk []InfoPart) []InfoPart {
	n := len(origDisk)

	disk := make([]InfoPart, len(origDisk))
	copy(disk, origDisk)

	for {
		lastFileIndex := -1
		for i := n - 1; i >= 0; i-- {
			if !disk[i].empty {
				lastFileIndex = i
				break
			}
		}

		if lastFileIndex == -1 {
			break
		}

		firstFreeIndex := -1
		for i := 0; i < n; i++ {
			if disk[i].empty {
				firstFreeIndex = i
				break
			}
		}

		if firstFreeIndex == -1 || firstFreeIndex >= lastFileIndex {
			break
		}

		disk[firstFreeIndex] = disk[lastFileIndex]
		disk[lastFileIndex] = InfoPart{empty: true}
	}

	return disk
}
