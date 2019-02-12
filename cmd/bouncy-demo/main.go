package main

import (
	"encoding/json"
	"log"

	"github.com/blevesearch/bleve"

	// to trigger the init func
	_ "github.com/danmux/bouncy/foundationdb"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()

	// you will need to `rm -rf example.bleve`
	index, err := bleve.NewUsing("example.bleve", mapping, "upside_down", "foundationdb", nil)
	if err != nil {
		log.Fatalf("new failed <%v>", err)
	}

	// index some data
	err = index.Index("wikipedia-1", yourData)
	if err != nil {
		log.Fatal(err)
	}

	d := yourData + "widget"

	err = index.Index("wikipedia-2", d)
	if err != nil {
		log.Fatal(err)
	}

	jj := `{"widget": {
		"debug": "on",
		"window": {
			"title": "Sample Konfabulator Widget",
			"name": "main_window",
			"width": 500,
			"height": 500
		},
		"image": { 
			"src": "Images/Sun.png",
			"name": "sun1",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		},
		"text": {
			"data": "Click Here",
			"size": 36,
			"style": "bold",
			"name": "text1",
			"hOffset": 250,
			"vOffset": 100,
			"alignment": "center",
			"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
		}
	}}  `

	v := map[string]interface{}{}
	err = json.Unmarshal([]byte(jj), &v)
	if err != nil {
		log.Fatal(err)
	}

	err = index.Index("obj1", v)
	if err != nil {
		log.Fatal(err)
	}

	err = index.Index("obj2", v)
	if err != nil {
		log.Fatal(err)
	}

	err = index.Index("obj1", v)
	if err != nil {
		log.Fatal(err)
	}

	// search for some text
	query := bleve.NewMatchQuery("widget")
	search := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(searchResult.String())
}

var yourData = `From Wikipedia, the free encyclopedia
Jump to navigationJump to search
For other uses, see Regret (disambiguation).

John Greenleaf Whittier's fictional heroine Maud Muller gazes into the distance, regretting her inaction and thinking about what might have been.
Part of a series on
Emotions
Plutchik dyads.png
Acceptance Affection Anger Angst Anguish Annoyance Anticipation Anxiety Apathy Arousal Awe Boredom Confidence Contempt Contentment Courage Curiosity Depression Desire Despair Disappointment Disgust Distrust Ecstasy Embarrassment Empathy Enthusiasm Envy Euphoria Fear Frustration Gratitude Grief Guilt Happiness Hatred Hope Horror Hostility Humiliation Interest Jealousy Joy Loneliness Love Lust Outrage Panic Passion Pity Pleasure Pride Rage Regret Social connection Rejection Remorse Resentment Sadness Saudade Schadenfreude Self-confidence Shame Shock Shyness Sorrow Suffering Surprise Trust Wonder Worry
vte
Regret is a negative conscious and emotional reaction to one's personal decision-making, a choice resulting in action or inaction. Regret is related to perceived opportunity. Its intensity varies over time after the decision, in regard to action versus inaction, and in regard to self-control at a particular age. The self-recrimination which comes with regret is thought to spur corrective action and adaptation. In Western societies adults have the highest regrets regarding choices of their education.[citation needed]


Contents
1	Definition
2	Models
3	Life domains
4	Determinants of intensity
4.1	Action versus inaction
4.2	Age
4.3	Opportunity
4.4	Lost opportunity principle
4.5	In health care decisions
5	Neuroscience
6	In other species
7	See also
8	References
Definition
Regret has been defined by psychologists in the late 1990s as a "negative emotion predicated on an upward, self-focused, counterfactual inference".[1] Regret lingers where opportunity existed, with self-blame being a core element to ultimately spur corrective action in decision-making.[1] Another definition is "an aversive emotional state elicited by a discrepancy in the outcome values of chosen vs. unchosen actions".[2]

Regret is distinct from disappointment; Both are negative emotional experiences relating to a loss outcome, and both have similar neuronal correlates. However, they differ in regard to feedback about the outcome, comparing the difference between outcomes for the chosen vs. unchosen action; In regret, full feedback occurs and with disappointment partial feedback. They also differ in regard to agency (self in regret versus external in disappointment).[3]

Models
There are conceptual models of regret in regret (decision theory) mostly in theoretical economics and finance under a field called behavioral economics. The term buyer's remorse, also called buyer's regret which is assumed to be the cause for regret aversion.

Anticipated regret or how much regret one thinks one will feel in the future, appears to be overestimated for actions and choices.[4][5] This appears to be, in part, due to a tendency to underestimate the extent to which people attribute bad outcomes to external factors rather than to internal factors (i.e., themselves).[4] It can lead to inaction or inertia and omission bias.[6]

Existential regret has been specifically defined as "a profound desire to go back and change a past experience in which one has failed to choose consciously or has made a choice that did not follow one’s beliefs, values, or growth needs".[7]

Instruments to measure regret in people having to make medical decisions have failed to address current concepts of regret and failed to differentiate regret from disappointment. They have also not looked for positive impacts of regret.[8] Process regret may occur, if a person does not consider information about all available choices before making a decision.[8]

Life domains
A 2005 meta-analysis of 9 studies (7 US, one Germany, one Finland) about what adults regret most concluded, that overall adults regret choices regarding their education the most. Subsequent rankings included decisions about career, romance, and parenting. Education has been the forerunner of regret in the U.S. per Gallup surveys in 1949, 1953, and 1965. Education was the forerunner of regret because it is seen as something where circumstances could be changed: "In contemporary society, education is open to continual modification throughout life. With the rise of community colleges and student aid programs in recent decades, education of some sort is accessible to nearly all socioeconomic groups."This finding can be attributed to the principle of perceived opportunity. "People´s biggest regrets are a reflection of where in life they see their largest opportunities; that is, where they see tangible prospects for change, growth, and renewal.[1]

In other cultures, regrets may be ranked differently depending on the perceived opportunity in a particular society.[9]

Determinants of intensity
Action versus inaction
There is an interplay between action versus inaction and time. Regrets of an action are more intense in the short term, whereas regrets of inaction are more intense over the long term.[10]

Age
In a 2001 study, high intensity of regret and intrusive thoughts in older adults was related to self-control, and low internal control was expected to be self-protective and help to decrease regret. In younger adults, internal-control facilitated active change and was associated with low intensity of regret.[11]

Opportunity
People's biggest regrets occur where they perceive the greatest and most important opportunity for corrective action.[1] When no opportunity exists to improve conditions, thought processes mitigate the cognitive dissonance caused by regret, e.g. by rationalization, and reconstrual.[1] Regret pushes people toward revised decision making and corrective action as part of learning that may bring improvement in life circumstances. A 1999 study measured regret in accordance to negative reviews with service providers. Regret was an accurate predictor of who switched providers. As more intense regret is experienced, the likelihood of initiating change is increased. Consequently, the more opportunity of corrective action available, the larger the regret felt and the more likely corrective action is achieved. Feeling regret spurs future action to make sure other opportunities are taken so that regret will not be experienced again. People learn from their mistakes.[12]

Lost opportunity principle
With a lost opportunity regret should intensify, not diminish, when people feel that they could have made better choices in the past but now perceive limited opportunities to take corrective action in the future. "People who habitually consider future consequences (and how they may avoid future negative outcomes) experience less, rather than more, intense regret after a negative outcome." [13] This principle offers another reason as to why education is the most regretted aspect in life. Education becomes a more limited opportunity as time passes. Aspects such as making friends, becoming more spiritual, and community involvement tend to be less regrettable which makes sense because these are also aspects in life that do not become limited opportunities. As the opportunity to remedy a situation passes, feelings of hopelessness may increase.[14] An explanation of the lost opportunity principle can be seen as a lack of closure: Low closure makes past occurrences feel unresolved. Low closure is associated with "reductions in self-esteem and persistent negative affect over time" and with the realization and regret of lost opportunity. High closure is associated with acceptance of lost opportunity.[15]

The lost opportunity principle suggests, that regret does not serve as a corrective motive (which the opportunity principle suggests). Instead, regret serves as a more general reminder to seize the day.[citation needed]

In health care decisions
A 2016 review of past studies found risk factors for people to develop "decision regret" regarding their health care were: higher decisional conflict, lower satisfaction with the decision, adverse outcomes in physical health, and greater anxiety levels.[16]

Neuroscience
Research upon brain injury and fMRI have linked the orbitofrontal cortex to the processing of regret.[17][18]

Completeness of feedback about the outcomes after making a decision determined whether persons experienced regret (outcomes from both the choice and the alternative) vs. disappointment (partial-feedback, seeing only the outcome from the choice) in a magnetoencephalography study. Another factor was the type of agency: With personal decision making the neural correlates of regret could be seen, with external agency (computer choice) those of disappointment. Feedback regret showed greater brain activity in the right anterior and posterior regions, with agency regret producing greater activity in the left anterior region.[3] Both regret and disappointment activated anterior insula and dorsomedial prefrontal cortex but only with regret the lateral orbitofrontal cortex was activated.[19]

Psychopathic individuals do not show regret and remorse. This was thought to be due to an inability to generate this emotion in response to negative outcomes. However, in 2016, people with antisocial personality disorder and dissocial personality disorder were found to experience regret, but did not use the regret to guide their choice in behavior. There was no lack of regret but a problem to think through a range of potential actions and estimating the outcome values.[20]

In other species
A study published in 2014 by neuroscientists based at the University of Minnesota suggested that rats are capable of feeling regret about their actions. This emotion had never previously been found in any other mammals apart from humans. Researchers set up situations to induce regret, and rats expressed regret through both their behavior and specific neural patterns in brain activity.[21]

See also
Regret (decision theory)
Apology
References
	Wikimedia Commons has media related to Regret.
	Wikiquote has quotations related to: Regret
	Look up regret in Wiktionary, the free dictionary.
 Roese, N.J. (2005). "What We Regret Most...and Why". Personality & Social Psychology Bulletin. 31 (9): 1273–85. doi:10.1177/0146167205274693. PMC 2394712. PMID 16055646.
 Zeelenberg M, Pieters R. A theory of regret regulation 1.0. J Consum Psychol. 2007;17(1):3–18.
 Giorgetta, C; Grecucci, A; Bonini, N; Coricelli, G; Demarchi, G; Braun, C; Sanfey, AG (Jan 2013). "Waves of regret: a meg study of emotion and decision-making". Neuropsychologia. 51 (1): 38–51. doi:10.1016/j.neuropsychologia.2012.10.015. PMID 23137945.
 Gilbert, Daniel T.; Morewedge, Carey K.; Risen, Jane L.; Wilson, Timothy D. (2004-05-01). "Looking Forward to Looking Backward The Misprediction of Regret". Psychological Science. 15 (5): 346–350. CiteSeerX 10.1.1.492.9980. doi:10.1111/j.0956-7976.2004.00681.x. ISSN 0956-7976. PMID 15102146.
 Sevdalis, Nick; Harvey, Nigel (2007-08-01). "Biased Forecasting of Postdecisional Affect". Psychological Science. 18 (8): 678–681. doi:10.1111/j.1467-9280.2007.01958.x. ISSN 0956-7976. PMID 17680936.
 Dibonaventura, M; Chapman, GB (2008). "Do decision biases predict bad decisions? Omission bias, naturalness bias, and inf luenza vaccination". Med Decis Making. 28 (4): 532–9. CiteSeerX 10.1.1.670.2689. doi:10.1177/0272989x08315250. PMID 18319507.
 Lucas, Marijo (January 2004). "Existential Regret: A Crossroads of Existential Anxiety and Existential Guilt". Journal of Humanistic Psychology. 44 (1): 58–70. doi:10.1177/0022167803259752. Archived from the original on 3 July 2012. Retrieved 15 March 2011.
 Joseph-Williams N, Edwards A, Elwyn G. The importance and complexity of regret in the measurement of ‘good’ decisions: a systematic review and a content analysis of existing assessment instruments. Health Expect 2011; 14: 59-83 doi: 10.1111/j.1369-7625.2010.00621.x PMID 20860776, PMCID:PMC5060557
 Gilovich, T; Wang, RF; Regan, D; Nishina, S (2003). "Regrets of action and inaction across cultures". Journal of Cross-Cultural Psychology. 34: 61–71. doi:10.1177/0022022102239155.
 Gilovich, T; Medvec, VH (1995). "The experience of regret: What, when, and why". Psychological Review. 102 (2): 379–395. doi:10.1037/0033-295x.102.2.379.
 Wrosch, C; Heckhausen, J (2002). "Perceived control of life regrets: Good for young and bad for old adults". Psychology and Aging. 17 (2): 340–350. doi:10.1037/0882-7974.17.2.340. PMID 12061416.
 Zeelenberg, M (1999). "The use of crying over spilled milk: A note on the rationality and functionality of regret" (PDF). Philosophical Psychology. 13 (951–5089): 326–340. doi:10.1080/095150899105800.
 Roese, Neal J. (Jan 1997). "Counterfactual Thinking". Psychological Bulletin. 121 (1): 133–148. doi:10.1037/0033-2909.121.1.133. PMID 9000895.
 Beike, Denise (December 19, 2008). "What We Regret Most Are Lost Opportunities: A Theory of Regret Intensity". Personality and Social Psychology Bulletin. 35 (3): 385–397. doi:10.1177/0146167208328329. PMID 19098259. Archived from the original on 15 November 2014. Retrieved 11 May 2015.
 Beike, Denise; Wirth-Beaumont, Erin (2005). "Psychological closure as a memory phenomenon". Memory. 13 (6): 574–593. doi:10.1080/09658210444000241. PMID 16076673.
 Becerra Pérez MM, Menear M, Brehaut JC, Légaré F.Med Decis Making. Extent and Predictors of Decision Regret about Health Care Decisions: A Systematic Review. 2016 Aug;36(6):777-90. doi: 10.1177/0272989X16636113.
 Coricelli, HD; Joffily, M; O'Doherty, JP; Sirigu, A; Dolan, RJ (2007). "Regret and its avoidance: a neuroimaging study of choice behavior". Nat Neurosci. 8 (9): 1255–62. doi:10.1038/nn1514. hdl:21.11116/0000-0001-A327-B. PMID 16116457.
 Coricelli, G; Dolan, RJ; Sirigu, A (2007). "Brain, emotion and decision making: the paradigmatic example of regret". Trends Cogn Sci. 11 (6): 258–65. doi:10.1016/j.tics.2007.04.003. hdl:21.11116/0000-0001-A325-D. PMID 17475537.
 Chua HF1, Gonzalez R, Taylor SF, Welsh RC, Liberzon I.Decision-related loss: regret and disappointment. Neuroimage. 2009 Oct 1;47(4):2031-40. doi: 10.1016/j.neuroimage.2009.06.006.
 Baskin-Sommers, A; Stuppy-Sullivan, AM; Buckholtz, JW (2016). "Psychopathic individuals exhibit but do not avoid regret during counterfactual decision making". Proc Natl Acad Sci U S A. 113 (50): 14438–14443. doi:10.1073/pnas.1609985113. PMC 5167137. PMID 27911790.
 Steiner, Adam P; Redish, A David (2014-06-08). "Behavioral and neurophysiological correlates of regret in rat decision-making on a neuroeconomic task". Nature Neuroscience. 17 (7): 995–1002. doi:10.1038/nn.3740. ISSN 1546-1726. PMC 4113023. PMID 24908102.
Hein, David. "Regrets Only: A Theology of Remorse." The Anglican 33, no. 4 (October 2004): 5-6.`
