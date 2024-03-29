package generate_document

/*
! Generate Document
https://www.algoexpert.io/questions/generate-document

You’re given a string of available characters and a string representing a document that you need to generate.
Write a function that determines if you can generate the document using the available characters.
If you can generate the document, your function should return true; otherwise, it should return false.

You’re only able to generate the document if the frequency of unique characters in the characters string
is greater than or equal to the frequency of unique characters in the document string.

For example, if you’re given characters = “abcabc” and document = “aabbccc” you cannot generate the document because you’re missing one c.

The document that you need to create may contain any characters, including special characters, capital letters, numbers, and spaces.

Note: you can always generate the empty string ("").


! Example:
	characters = "Bste!hetsi ogEAxpelrt x "
	document = "AlgoExpert is the Best!"

	output: true
*/

func GenerateDocument(characters string, document string) (output bool) {
	lookupTable := map[rune]int{}

	if document == "" {
		return true
	}

	for _, char := range characters {
		lookupTable[char]++
	}

	for _, char := range document {
		if lookupTable[char] == 0 {
			return false
		}

		lookupTable[char]--
		output = true
	}

	return
}
