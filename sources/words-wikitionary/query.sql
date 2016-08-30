SELECT DISTINCT(word) FROM
    (
        SELECT IF(cl_sortkey_prefix = "", cl_sortkey, cl_sortkey_prefix) AS word 
        FROM categorylinks 
        WHERE (cl_to like "English_%" OR cl_to like "en:%")
    ) AS TMP
    WHERE word REGEXP "^[A-Z '-]+$";
