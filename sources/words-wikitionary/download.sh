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
# Expanded this mostly a 2G, on one line, INSERT statement :-(
#

DATE=20160820
TARGET=enwiktionary-${DATE}-categorylinks.sql.gz
MYSQLUSER="-u root"
DBNAME="words"

if [ ! -f ${TARGET} ] ; then
   curl -L -o ${TARGET} \
	https://archive.org/download/enwiktionary-${DATE}/${TARGET}
fi

echo "DROP DATABASE IF EXISTS ${DBNAME}; CREATE DATABASE ${DBNAME}" | mysql ${MYSQLUSER}
zcat ${TARGET} | mysql ${MYSQLUSER} ${DBNAME} 
