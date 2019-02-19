package main

import (
	"fmt"
	"github.com/DavidBelicza/TextRank"
)

func main() {
	rawText := "THE fiery rhetoric emitting from the mouth of one of the United for National Development " +
		"(UPND) hotheads Gilbert Liswaniso is a reflection of how guilt is gnawing at their " +
		"conscience in the wake of the atrocities they committed AGAINST THE Patriotic Front (PF) supporters in Sesheke during the recent parliamentary by-election.	Liswaniso, one of UNPD leader Hakainde Hichilema’s attack dogs, was assigned to issue an" +
		" incendiary statement in Lusaka today in which his party is as usual whining " +
		"and sniveling." +
		"The possibility of getting a dose of its own medicine in Bahati" +
		" should it try to export violence to Luapula Province during the forthcoming by-election " +
		"is depriving them of sleep." + "The blood bath that UPND thugs led by Jack Mwiimbu and Gary Nkombo " +
		"visited upon the PF foot soldiers is causing the party to issue incoherent and self-pitying statements " +
		"such as the one Liswaniso issued today." + "Liswaniso said the UPND is ready to defend itself against any " +
		"provocation in Bahati Constituency." + "Who has said they will provoke his violent party?	" +
		"That is the typical language of the guilty. Te said the UPND is ready to die for mother Zambia." +
		"But the opposition party is already dying slowly as it continues to allow a perpetual" +
		" loser to be at its helm while losing members in droves.Liswaniso can inform his fellow tribalists in the" +
		" UPND that the Ushi people of Bahati, like other Bemba-speaking tribes, are traditionally and historically" +
		" warriors who will fiercely and fearlessly fight back should the UPND try to export the violence it perpetrated " +
		"in Sesheke to their constituency.Should the UPND thugs try to terrorise voters and intimidate " +
		"PF foot soldiers during this by-election they will know the difference between professional warriors and farmers." +
		"The PF has never said it will attack UPND zealots in Bahati but" +
		" simply warned that the people of Luapula are ready to eject them from their area should they try to cause violence."
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases order by weight.
	rankedPhrases := textrank.FindPhrases(tr)
	// Most important phrase.
	fmt.Println(rankedPhrases[0])

	// Get all words order by weight.
	words := textrank.FindSingleWords(tr)
	// Most important word.
	fmt.Println(words[0])

	// Get the most important 10 sentences. Importance by phrase weights.
	sentences := textrank.FindSentencesByRelationWeight(tr, 10)
	// Found sentences
	fmt.Println("SSentences By Weight:  ", sentences)

	// Get the most important 10 sentences. Importance by word occurrence.
	sentences = textrank.FindSentencesByWordQtyWeight(tr, 10)
	// Found sentences
	fmt.Println("Sentences By Wieght :  ", sentences)

}
