
-- delete empty and small words
DELETE FROM enwords where LENGTH(word) < 4;

-- delete this tag.  If word has other tags, it survives.
-- DELETE FROM enwords WHERE tag = "English_terms_with_obsolete_senses";

-- delete all words completely if tag matches
CREATE TEMPORARY TABLE IF NOT EXISTS junk1 AS (
  SELECT DISTINCT(id) from enwords 
  WHERE tag IN (
    "English_archaic_forms",
    "English_abbreviations",
    "English_misspellings",
    "English_dated_forms",
    "English_obsolete_terms",
    "English_obsolete_forms",
    "English_initialisms",
    "English_terms_with_obsolete_senses"
  )
);
DELETE FROM enwords
USING enwords, junk1
WHERE enwords.id = junk1.id;

-- delete words that are errors in wikitionary, or 
-- likely to be a misspelling anyways
CREATE TEMPORARY TABLE IF NOT EXISTS junk2 AS (
    SELECT DISTINCT(id) from enwords 
    WHERE word IN (
	"TATH",
	"AGIN",
	"ACCIDENTLY",
        "ABUTTS",
	"ADN",
	"ALOT",
	"ANNOINT",
	"ANNOINTED",
	"ANNOINTING",
	"ANNOINTS",
	"ANSWERES",
	"BELIEFES",
	"BELIVE",
	"BELIVES",
	"CANCELLS",
	"CHINEES",
	"CANNISTERS",
	"CRACKES",
	"NINTY",
	"DEVASTED",
	"DISPENCED",
	"DISPENCING",
	"DRINKES",
	"CIGARETS",
	"OFFRED",
	"EUROPIAN",
	"PROOVED",
	"INSTRUCTER",
	"VESSELLS",
	"INSTRUCTERS",
	"INSTALS",
	"WICH",
	"EXPENCES",
	"EVERYTING",
	"EXPELLS",
	"VITAMINES",
	"STRAT",
	"STONG",
	"HUNDERD",
	"ECT",
	"TEH",
	"DOUB",
	"CHAMBRE",
	"CHAMBRES",
	"CLAS",
	"HALP",
	"CONTINUOS",
	"FACIST",
"THRID",
"TALLENTS",
"TALLETS",
"TEACHED",
"SURVIVERS",
"SURVIVOURS",
"SURPRIZE",
"SURPRIZED",
"STREAMES",
"STUPIDY",
"STREAMES",
"STILUS",
"SUCCEDED",
"SUCCEDES",
"RISED",
"SAUGHT",
"QUITTED",
"RAELISM",
"PYRAMIDES",
"PROPPER",
"PERSUED",
"PERSUING",
"PAYED",
"PATTENED",
"NUCULAR",
"NESTIN",
"MEDICINS",
"MAKED",
"MACHINS",
"LENGTHLY",
"FLACONS",
"FOUTH",
"FROME",
"GALATIC",
"HEARED",
"HARRASES",
"HARRASSED",
"INMIGRANT",
"INMIGRANTS",
"RABBITTS"
	)
);

DELETE FROM enwords 
USING enwords, junk2
WHERE enwords.id = junk2.id;
