###

INSERT INTO :table (field1, field2) VALUES (:field1, :field2)

DELETE FROM :table WHERE :where LIMIT 1

UPDATE :table SET field1=value1,field2=value2 WHERE :where

SELECT field1,field2 FROM :table WHERE :where GROUP BY :groupBy HAVING :having ORDER BY :orderBy LIMIT :limit OFFSET :offset