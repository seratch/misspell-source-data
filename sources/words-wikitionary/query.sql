SELECT DISTINCT(word) FROM
    (
        SELECT IF(cl_sortkey_prefix = "", cl_sortkey, cl_sortkey_prefix) AS word 
        FROM categorylinks 
        WHERE (cl_sortkey not like "%'%")
	  AND (cl_to like "English_%" OR cl_to like "en:%")
	  AND (cl_to <> "English_misspellings")
    ) AS TMP
    WHERE word REGEXP "^[A-Z ]+$";
