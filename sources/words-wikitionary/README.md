# Wikitionary

Wikitionary is interesting since it's not a dictionary exactly.
The English version is a cataloging of _all words_ in _all languages_
with definitions in _English_.

Using the complete word list (from all languages), and removing
non-ASCII words, still contains Old English, Middle English, Latin,
obsolete forms, is well has many European-based languages.

There isn't a clean way of just getting "English".    Even the
[English Index](https://en.wiktionary.org/wiki/Index:English) says it was
"The 271073 terms on this page were extracted from the 2012-Apr-28 database
dump." which implies it's not trivial.


However, there are many tags for each word:

* Most of English words are in the
  [English_](https://en.wiktionary.org/wiki/Category:English_language)
  "catagory"
* Additional technical or specialized words are in "topics" starting with
  [en:](https://en.wiktionary.org/wiki/Category:en:All_topics)
* Obsolete forms start with `English_obsolete_` in them. There are two
  * [English_obsolete_forms](https://en.wiktionary.org/wiki/Category:English_obsolete_forms) 
  * [English_obsolete_terms](https://en.wiktionary.org/wiki/Category:English_obsolete_terms)
* The category
  [English_terms_with_obsolete_senses](https://en.wiktionary.org/wiki/Category:English_terms_with_obsolete_senses)
  are well things are words that _have_ obsolete variations?  I can't tell.
  [able](https://en.wiktionary.org/wiki/able#English) is an example of one.

### Internet Archive Dump Files

The Internet Archive contains dumps from Wikitionary on a regular basis.

Unfortunately it's not XML (which would be streamable).  It's in Mysql Dump
format of their mediawiki database.

An sample dump is https://archive.org/details/enwiktionary-20160820

The `name` files are words that have entries.  As mentioned above, this is not
useful.

The `categorylinks` file maps words to categories.  Uncompressed it is 
which is a 1 line, 2 gigabyte `INSERT` statement.  :-(

The [download.sh](sources/words-wikitionary/download.sh) script will
attempt to download a dump and insert it into mysql.

### Amazon Linux

Using [AWS Linux](https://aws.amazon.com/amazon-linux-ami/faqs/)  since it's easy and throwaway.

```bash
sudo yum update
sudo yum install mysql56-server mysql56
sudo /etc/init.d/mysqld start
```

The root user is used and passwords are not set since this is never exposed to the
Internet.

### Scripts

* download.sh  -- download and imports data into mysql
* dropindex.sh -- drops the crazy indexes that aren't needed
* enwords.sh -- copies English words into a more-sane table format
* clean.sh -- removed bogus English words
* query.sh -- emits a word file, sorted.
