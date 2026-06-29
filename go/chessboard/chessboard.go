package chessboard

const fileLen = 8

type File [fileLen]bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cbFile, ok := cb[file]
	if !ok {
		return 0
	}

	count := 0
	for i := range fileLen {
		if cbFile[i] {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	index := rank - 1
	if index < 0 || index > fileLen-1 {
		return 0
	}

	count := 0
	for _, file := range cb {
		if file[index] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return fileLen * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for i := range fileLen {
			if file[i] {
				count++
			}
		}
	}

	return count
}
