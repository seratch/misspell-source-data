
-- simplified schema
DROP TABLE IF EXISTS enwords;
CREATE TABLE `enwords` (
  `id` int(8) unsigned NOT NULL DEFAULT '0',
  `tag` varbinary(255) NOT NULL DEFAULT '',
  `word` varbinary(230) NOT NULL DEFAULT '',
  UNIQUE KEY `cl_from` (`id`,`tag`),
  KEY `tag` (`tag`),
  KEY `word` (`word`)
) ENGINE=InnoDB DEFAULT CHARSET=binary ROW_FORMAT=DYNAMIC;

-- sometimes the English word is in cl_sortkey and other times it's in cl_sortkey_prefix
-- this tries to pick the right one
-- Eliminate all hyphanted words and words with apostrophe.
--    misspell doesn't work with those.
-- Eliminate all non-ASCII words -- 
--   * all sorts of emojis and symbols
--   * braille forms
--   * don't care about accented words
-- NOTE: sortkey regexp includes a newline.  For unknown reasons
--    sortkey is often in the form of "FOO\nFOO"
--
INSERT INTO enwords (id, tag, word) 
        SELECT cl_from AS id, cl_to as tag, IF(cl_sortkey_prefix = "", cl_sortkey, cl_sortkey_prefix) AS word 
        FROM categorylinks 
	WHERE (cl_to LIKE "English_%" OR cl_to LIKE "en:%")
	  AND cl_sortkey REGEXP "^[A-Z\n]*$"
          AND cl_sortkey_prefix REGEXP "^[A-Z]*$";
