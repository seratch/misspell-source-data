
-- delete this tag.  If word has other tags, it survives.
DELETE FROM categorylinks WHERE cl_to = "English_terms_with_obsolete_senses";

-- delete all words completely if tag matches
DELETE tmp.*
FROM categorylinks tmp
WHERE cl_from IN (
	SELECT cl_from FROM (
		  SELECT DISTINCT(cl_from) from categorylinks
		  WHERE (cl_to = "English_archaic_forms")
		     OR (cl_to = "English_abbreviations")
		     OR (cl_to = "English_misspellings")
		     OR (cl_to LIKE "English_obsolete%")
	) x );
