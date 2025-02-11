package generator

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	@todo :: add more personalities.
	@todo :: add formatting (camelCase, under_score, spaces).
	@todo :: improve people's comments.
*/

var (
	adjectives = [...]string{
		"able", "absorbed", "abundant", "accurate", "admiring", "adorable", "adventurous", "affectionate", "agreeable", "alert", "amazing", "amusing", "animated", "aspiring", "astonishing", "attractive", "auspicious", "automatic", "available", "awake", "aware", "awesome", "beautiful", "befitting", "beneficial", "best", "better", "big", "blissful", "bold", "boundless", "brainy", "brave", "bright", "broad", "bustling", "calm", "capable", "careful", "caring", "certain", "charming", "cheerful", "chivalrous", "classy", "clean", "clear", "clever", "cocky", "coherent", "colorful", "colossal", "comfortable", "compassionate", "competent", "complete", "confident", "conscious", "cool", "cooperative", "coordinated", "courageous", "cuddly", "cultured", "curious", "curly", "cute", "daffy", "dashing", "dazzling", "dear", "debonair", "decisive", "decorous", "delicate", "delicious", "delightful", "determined", "diligent", "discreet", "dreamy", "eager", "easy", "ecstatic", "educated", "efficacious", "efficient", "elastic", "elated", "electric", "elegant", "elfin", "eloquent", "eminent", "enchanting", "encouraging", "energetic", "entertaining", "enthusiastic", "epic", "equable", "ethereal", "excellent", "excited", "exciting", "exuberant", "exultant", "fabulous", "fair", "faithful", "familiar", "famous", "fancy", "fantastic", "fascinated", "fast", "fearless", "fervent", "festive", "fine", "fixed", "flamboyant", "flashy", "flawless", "flowery", "fluffy", "focused", "fortunate", "free", "frequent", "fresh", "friendly", "full", "funny", "furry", "future", "futuristic", "gainful", "gallant", "gentle", "giant", "gifted", "gigantic", "glamorous", "gleaming", "glistening", "glorious", "glossy", "godly", "good", "gorgeous", "graceful", "gracious", "grandiose", "grateful", "gratis", "great", "groovy", "grouchy", "handsome", "handsomely", "handy", "happy", "harmonious", "healthy", "heavenly", "helpful", "heuristic", "hilarious", "holistic", "homely", "honorable", "hopeful", "hospitable", "hot", "huge", "humorous", "hushed", "illustrious", "immense", "important", "incredible", "industrious", "infallible", "innocent", "inquisitive", "inspiring", "intelligent", "interesting", "invincible", "jolly", "jovial", "joyous", "judicious", "juicy", "keen", "kind", "kindhearted", "knowledgeable", "large", "laughing", "learned", "legal", "like", "likeable", "literate", "lively", "living", "long", "lovely", "loving", "lucid", "lucky", "luxuriant", "magical", "magnificent", "majestic", "marvelous", "massive", "mellow", "melodic", "merciful", "mighty", "modern", "modest", "momentous", "mysterious", "mystifying", "nappy", "natural", "neat", "necessary", "neighborly", "new", "nice", "nifty", "nimble", "nippy", "noiseless", "nonchalant", "nonstop", "normal", "nosy", "numerous", "nutritious", "objective", "observant", "obtainable", "oceanic", "omniscient", "open", "optimal", "optimistic", "organic", "outgoing", "outstanding", "overjoyed", "panoramic", "parallel", "pastoral", "peaceful", "pensive", "perfect", "periodic", "permissible", "perpetual", "physical", "piquant", "plausible", "pleasant", "plucky", "poised", "polite", "possible", "powerful", "practical", "precious", "premium", "present", "pretty", "priceless", "prodigious", "productive", "profuse", "protective", "proud", "psychedelic", "public", "pumped", "quaint", "quick", "quickest", "quiet", "quirky", "quizzical", "rampant", "rapid", "rare", "ready", "real", "rebel", "receptive", "reflective", "regular", "relaxed", "relieved", "remarkable", "reminiscent", "resolute", "resonant", "responsible", "reverent", "rich", "right", "righteous", "rightful", "ripe", "ritzy", "robust", "romantic", "roomy", "round", "royal", "ruddy", "rural", "rustic", "safe", "salty", "sassy", "satisfying", "savory", "scientific", "scintillating", "seemly", "selective", "serene", "serious", "sharp", "shiny", "silky", "silly", "simple", "sincere", "skillful", "skinny", "sleepy", "slim", "smart", "smiling", "smooth", "soft", "solid", "sophisticated", "sparkling", "special", "spectacular", "spicy", "spiffy", "spiritual", "splendid", "spotless", "spotted", "spotty", "standing", "statuesque", "steadfast", "steady", "stimulating", "stoic", "straight", "striped", "strong", "stupendous", "sturdy", "substantial", "successful", "succinct", "super", "superb", "supreme", "swanky", "sweet", "swift", "tacit", "talented", "tall", "tame", "tangible", "tangy", "tasteful", "tasty", "telling", "tender", "terrific", "tested", "thankful", "therapeutic", "thick", "thin", "thinkable", "thoughtful", "thundering", "tidy", "tight", "tiny", "toothsome", "tough", "towering", "tranquil", "tremendous", "true", "trusting", "truthful", "ubiquitous", "ultra", "unarmed", "unbiased", "uncovered", "understood", "unique", "unruffled", "unused", "unusual", "upbeat", "useful", "utopian", "utter", "uttermost", "valuable", "various", "vast", "verdant", "versed", "vibrant", "victorious", "vigilant", "vigorous", "vivacious", "wacky", "waggish", "wakeful", "warm", "watery", "wealthy", "whole", "wide", "wild", "willing", "wise", "witty", "wonderful", "workable", "xenodochial", "yielding", "young", "youthful", "yummy", "zany", "zealous", "zen", "zesty", "zippy",
	}

	colors = [...]string{
		"almond", "amber", "amethyst", "apricot", "aqua", "aquamarine", "avocado", "azure", "beige", "black", "blond", "blue", "brandy", "bronze", "brown", "burgundy", "cadet", "carmine", "celeste", "cerulean", "charcoal", "citron", "coral", "cyan", "denim", "emerald", "fuchsia", "gold", "gray", "green", "indigo", "iris", "ivory", "khaki", "lavender", "lemon", "lilac", "lime", "magenta", "mango", "mint", "mustard", "navy", "neon", "ochre", "olive", "onyx", "opal", "orange", "orchid", "peach", "pink", "plum", "purple", "red", "rose", "ruby", "salmon", "sand", "sapphire", "scarlet", "sepia", "sienna", "silver", "sinopia", "teal", "turquoise", "ultramarine", "vanilla", "vermilion", "violet", "viridian", "white", "yellow", "zaffre",
	}

	people = [...]string{
		"albattani",
		"archimedes",
		"aslan",
		"bell",
		"benz",
		"bohr",
		"bose",
		"coanda",
		"curie",
		"darwin",
		"davinci",
		"dijkstra",
		"dirac",
		"edison",
		"einstein",
		"euclid",
		"euler",
		"faraday",
		"fermat",
		"fermi",
		"feynman",
		"franklin",
		"galileo",
		"gauss",
		"hawking",
		"heisenberg",
		"hertz",
		"hopper",
		"hypatia",
		"kepler",
		"lovelace",
		"lumiere",
		"maxwell",
		"mendel",
		"mendeleev",
		"merkle",
		"moore",
		"morse",
		"nash",
		"newton",
		"nobel",
		"pascal",
		"pasteur",
		"ptolemy",
		"roentgen",
		"swartz",
		"tesla",
		"torvalds",
		"turing",
		"wozniak",
		"wright",
		"singh",
	}
)

func Name() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s %s %s", adjectives[rand.Intn(len(adjectives))], colors[rand.Intn(len(colors))], people[rand.Intn(len(people))])
}
