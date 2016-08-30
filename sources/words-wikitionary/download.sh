#!/bin/sh
set -ex

#
# Wikitionary is unsual since it's "all words, from all languages"
# with definitions in English.  There is no way of getting "just English words".
# (Well, the SimpleEnglish dump is maybe the top 10,000 English words but
#  that is too small for our purposes)
#
# Just downloads the the categorylinks file.  This maps
# a word to a category.  English words are categories in
# one of a thousand "English_XXX" categories, with suppliments
# in the "en:XXX" category.
#
# Expanded this is about 2G
DATE=20160820
TARGET=enwiktionary-${DATE}-categorylinks.sql.gz
if [ ! -f ${TARGET} ] ; then
   curl -L -o ${TARGET} \
	https://archive.org/download/enwiktionary-${DATE}/${TARGET}
fi

echo "DROP IF EXISTS DATABASE words; CREATE DATABASE WORDS" | mysql
gzcat ${TARGET} | mysql words
