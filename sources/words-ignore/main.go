package main

import (
	"fmt"
	"sort"
)

// these are words that are
//  Not Typos (do not correct to anything)
//  Not Real Words (do not use in clustering)
//
// They are words we just want to ignore
// Frequently they are proper nounds such as
//  * Names (given or family)
//  * Brands
//  * Words from  popular works of fiction
//  * Technical words
//  * Words from other langauges
//  * swear words in various types
//  * Intentional misspellings and slang (e.g. wannabe, automagical)
//
var words = []string{
	"replicaset",
	"klingon",
	"sennheiser", // brand, name
	"demarco",
	"chillin",
	"palpatine",
	"alexandria",
	"alexander",
	"donovan",
	"kakashi", // Japanese for "scarecrow"/
	"welbeck",
	"mcmahon",
	"ganondorf",
	"jeffrey",
	"forsett",
	"shyvana",
	"froakie",
	"jenkins",
	"walgreens",
	"buzzfeed",
	"kershaw",
	"rebecca",
	"gandalf",
	"apeshit",
	"cazorla",
	"maximus",
	"reginald",
	"andrews",
	"abdullah",
	"ableton",
	"abrahamic",
	"accutane", //brand (drug)
	"acuerdo",
	"aguero",
	"aldrig",  // ??
	"alguien", // unknown
	"alienware",
	"alistar",  // name
	"ambrose",  // name
	"amendola", // unknown
	"amiibos",
	"amirite",  // slang
	"anderson", // name
	"andreas",
	"antonio",
	"aquaman",
	"arduino",
	"atlantis",
	"atletico", // ?
	"baldwin",  // name
	"balotelli",
	"baratheon", // ??
	"barbara",
	"baretta",
	"batshit",
	"bayonetta",
	"beckham",
	"benjamin",
	"bernard",
	"bioshock",
	"bitchin",
	"bjergsen",
	"bledsoe",    // name
	"blitzcrank", // ??
	"boohoooooo",
	"bortles", // ??
	"bradley", // name
	"bradshaw",
	"bridgewater", // brand
	"brienne",
	"britney",
	"bronies",
	"bulbasaur",
	"bullshitting",
	"bundesliga",
	"caitlyn",
	"cartoony", // FP
	"castlevania",
	"catelyn",
	"celestia",
	"charizard",
	"charles",
	"charleston", //place name with possible variations
	"charlie",
	"charlotte", // name
	"charmander",
	"chipotle",
	"chomsky",
	"christie",
	"christina", // name
	"christine", // name
	"christopher",
	"chromebook",
	"chromecast",
	"chronos",
	"cinderhulk",
	"circlejerk",
	"circlejerking",
	"circlejerks",
	"ciudadanos", // Spanish
	"clusterfuck",
	"collins",
	"connery",
	"constantine",
	"copypasta",
	"courtney",
	"craigslist",
	"crawford",
	"chrysler", // brand
	"cortana",
	"coutinho",
	"crossfit", // brand
	"cualquier",
	"dawkins",
	"dratini",
	"dempsey",
	"daenerys",
	"dailymotion", // brand
	"daniels",
	"darkrai",
	"deandre",
	"dennis",
	"dipshit",
	"doritos", // brand
	"douchebags",
	"dragonball",
	"dragonborn",
	"dragonite",
	"einstein",
	"evangelion", // brand, name
	"expandjs",   // javascript something?
	"edwards",    // name
	"distros",    // too small,
	"deviantart", // brand
	"dignitas",
	"dogecoin",
	"disneyland", // brand
	"fatebringer",
	"ferguson",
	"fernando", // name
	"fittit",
	"fitzgerald", //name
	"floats",
	"francis",
	"freakin",
	"frickin",
	"friendzone", // causual word
	"fuckton",
	"gabriel", //name
	"galactica",
	"gilbert",
	"gjallarhorn",
	"goddamned",
	"goddamnit",
	"godzilla",
	"golems",
	"googled",
	"gracias",
	"greatsword",
	"gregory",
	"greninja",
	"griffin",
	"grinds",
	"harbaugh",
	"hawkmoon",
	"hecarim",
	"hermione",
	"herrera",
	"historia", //??
	"horseshit",
	"illidan",
	"importante", // Spanish
	"iniesta",    // ??
	"iphones",    // brand
	"jessica",
	"jigglypuff",
	"johnson",
	"judgment", // misspelling
	"jungler",
	"junglers",
	"kalista", // name
	"kanthal",
	"karambit",
	"kassadin",
	"katarina",
	"katowice",
	"kendrick", // name, multiple spellings
	"killionaires",
	"kingston",
	"krugman",     // name
	"lamborghini", // name, brand
	"leblanc",
	"lebowski", // name
	"leonardo", // name
	"lightsaber",
	"linkedin",
	"lucario",
	"machina",
	"magicka",
	"magneto",
	"manziel",
	"mariota",
	"marshawn",
	"martinez",
	"mayweather",
	"mcdonalds",
	"megaman",
	"metroid",
	"michael",
	"michaels",
	"minecraft",
	"mitchell",
	"mohammed",
	"monedero",
	"monsanto",
	"morgana",
	"motherfucker",
	"motherfuckers",
	"motherfucking",
	"mourinho",
	"muchos", // spanish?
	"muhammad",
	"murican", // slang
	"murphy",
	"myfitnesspal",
	"mythbusters", // brand
	"natalie",     // name
	"netanyahu",
	"netflix",
	"neville", // name
	"nicholas",
	"nickelback", // brand
	"nicolas",    // name
	"nintendo",
	"nosotros",
	"okcupid",
	"outside",
	"overheat",
	"pacquiao",
	"partido",
	"patreon", // misspelling of patron
	"pcpartpicker",
	"phones",
	"pikachu",
	"planetside", // brand
	"porsche",
	"presser",
	"problema",
	"programa",
	"radiohead", // name
	"redditor",
	"redditors",
	"reddits",
	"retards",
	"rngesus",
	"roberts", // name
	"rodriguez",
	"ronaldo",
	"rosalina",
	"runescape",
	"samsung",
	"sanchez",
	"santorin",
	"sarkeesian",
	"scientology",
	"seinfeld",
	"sejuani",
	"sephora",
	"shithead",
	"shitlord",
	"shitlords",
	"shitpost",
	"shitposting",
	"shitposts",
	"shitter",
	"shittier",
	"shittiest",
	"shoulda",
	"siempre",
	"sigelei",
	"skeltal", // misspelling of skeletal?
	"snapchat",
	"soundcloud", // company
	"spaces",
	"spongebob",
	"starcraft",
	"stephanie",
	"stephen",     // name
	"stomps",      // lots of FP
	"superheroes", // FP
	"superstars",  // FP
	"sverige",
	"swordbearer",
	"tarantino",
	"targaryen",
	"templatejs",
	"terran",
	"terrans", // (in science fiction) an inhabitant of the planet Earth.
	"terraria",
	"thunderlord",
	"tolkien",
	"tristana",
	"trudeau",
	"ventura",
	"veronica", // name
	"villionaires",
	"vladimir",
	"voldemort",
	"vonnegut", // name
	"douglas",  // name
	"walmart",
	"warlock",
	"dreamcast",
	"dreamhack",
	"dumbasses",
	"dumbledore",
	"fabregas",
	"fucktard",
	"fleshlight",
	"futurama",
	"gangsta",
	"gambino",
	"genesect",
	"ghostbusters",
	"grimoire",
	"gobierno",
	"gonzalez",
	"heisenberg",
	"hendricks",
	"hitchens",
	"hogwarts",
	"honedge",
	"hyundai",
	"instagram",
	"jirachi",
	"warpig",
	"westeros",
	"kaepernick",
	"kardashian",
	"karthus",
	"katrina",
	"whatcha", // slang
	"affleck",
	"limbaugh",
	"whatsapp",
	"william",
	"witcher",
	"wrestlemania",
	"youtuber",
	"zelnite",
	"lololol",
	"macklemore",
	"magikarp",
	"malphite",
	"materia",
	"mcdonald",
	"melissa",
	"metacritic",
	"micheal",
	"microsoft",
	"nietzsche",
	"nordstrom",
	"orianna",
	"optimus",
	"phillips",
	"playstation",
	"archers",
	"rhaegar",
	"reddiquette",
	"richardson",
	"richards",
	"roosevelt",
	"scarlet",
	"shadowbanned",
	"shadowrun",
	"schneider",
	"shitload",
	"skrillex",
	"sturridge",
	"scarlett",
	"sakurai",
	"arsehole",
	"sylvanas",
	"thalmor",
	"tmobile",
	"triforce",
	"tryndamere",
	"ubisoft",
	"victoria",
	"virginia",
	"warlocks",
	"wannabe",
	"westboro",
	"wikipedia",
	"winterfell",
	"wolfenstein",
	"xpecial",
	"yogscast",
	"administrador", // Spanish
	"officiella",
	"undernet",      // name of computer network
	"transitionend", // JS event
	"invalide",      // French
	"releasers",
	"supplementals", // should be corrected to supplemental (adjective)
	"officiella",
	"authorid",
	"characterid",
	"contractid",
	"customid",
	"positionid",
	"discountid",
	"employeeid",
	"multicasting",
	"questionid",
	"randomid",
	"summonerid",
	"installare",    // IT
	"installato",    // IT
	"dokumentation", // DE
	"superpowder",
	"proletara",
	"pediction",
	"motorola",
	"moranian",
	"exerciseid",
	"mingleplayer",
	"melbournite",
	"mathological",
	"carnagie", //name
	"granda",
	"sourceid",
	"messageid",
	"certificacion",    // Spanish
	"progresse",        // progesses, progressives
	"generalizaciones", //spanish
	"sithlord",
	"administratie", // Dutch spelling
	"killionaires",
	"villionaires",
	"zillionaires",
	"codifications",
	"administracion",
	"potentiella",
	"pistos",
	"corpe",
	"reactjs", // name of product
	"chroo",   // not clear what this is a typo of
	"cleane",  // not clear if "cleaner" or "cleanser", causes FP
	"mongos",
	"mongod", // mongo database daemon
	"parens", // common for parenthesis
	"thru",   // informal, style
	"warpig",
	"intereating",
	"interdating",
	"dogspeed",
	"laventine",
	"administrador", // Spanish
	"expresso",      // wrong guess
	"administra",    // Spanish?
	"launchd",       // name of program
	"systemd",
	"collectd",
	"octobre",
	"novembre",
	"septembre",
	"revolucion",
	"tortilleras",
	"automagically",
	"benchmarkers",
	"englisch",
}

func main() {
	sort.Strings(words)
	for _, word := range words {
		fmt.Println(word)
	}
}
