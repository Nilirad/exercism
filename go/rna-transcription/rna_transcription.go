package rnatranscription

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))

	for i := range len(dna) {
		switch dna[i] {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		default:
			return ""
		}
	}

	return string(rna)
}
