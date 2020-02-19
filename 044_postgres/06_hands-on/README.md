# hands-on exercise
1. delete all of your current tables.
1. READ ALL OF THIS: create a new table called employees with these fields ```id, name, score, salary``` AND give ```score``` a default value of 10 AND have the ```id``` field automatically increment.
1. add these records and then show all of the records
```
 id |  name  | score | salary 
----+--------+-------+--------
  1 | Daniel |    23 |  55000
  2 | Arin   |    25 |  65000
  3 | Juan   |    24 |  72000
  4 | Shen   |    26 |  64000
  5 | Myke   |    27 |  58000
  6 | McLeod |    26 |  72000
  7 | James  |    32 |  35000
```

CREATE TABLE employees (
   ID SERIAL PRIMARY KEY,
   NAME           TEXT    NOT NULL,
   SCORE          INT DEFAULT 10,
   SALARY         REAL
   
  );
  
INSERT INTO employees (NAME,SCORE,SALARY) VALUES ('Daniel',23,55000),('Arin',25,65000),('Juan',24,72000);
INSERT INTO employees (NAME,SCORE,SALARY) VALUES ('Shen',26,64000),('Myke',27,58000),('McLeod',26,72000),('James',32,35000);
(4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');