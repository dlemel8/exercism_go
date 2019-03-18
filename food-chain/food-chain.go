package foodchain

func Verse(i int) string {
	if i == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	res := "I know an old lady who swallowed a "
	firstPart := map[int]string{
		1: "fly.\n",
		2: "spider.\nIt wriggled and jiggled and tickled inside her.\n",
		3: "bird.\nHow absurd to swallow a bird!\n",
		4: "cat.\nImagine that, to swallow a cat!\n",
		5: "dog.\nWhat a hog, to swallow a dog!\n",
		6: "goat.\nJust opened her throat and swallowed a goat!\n",
		7: "cow.\nI don't know how she swallowed a cow!\n",
	}
	res += firstPart[i]

	switch i {
	case 7:
		res += "She swallowed the cow to catch the goat.\n"
		fallthrough
	case 6:
		res += "She swallowed the goat to catch the dog.\n"
		fallthrough
	case 5:
		res += "She swallowed the dog to catch the cat.\n"
		fallthrough
	case 4:
		res += "She swallowed the cat to catch the bird.\n"
		fallthrough
	case 3:
		res += "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n"
		fallthrough
	case 2:
		res += "She swallowed the spider to catch the fly.\n"
	}
	return res + "I don't know why she swallowed the fly. Perhaps she'll die."
}

func Verses(start int, end int) string {
	res := ""
	for i := start; i < end; i++ {
		res += Verse(i) + "\n\n"
	}
	return res + Verse(end)
}

func Song() string {
	return Verses(1, 8)
}
