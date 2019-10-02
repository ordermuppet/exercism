package strand

func ToRNA(dna string) string {
	var rna string
	complements := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
	for _, nucleotide := range dna {
		rna += complements[nucleotide]
	}
	return rna
}
