## PL/pgSQL

### №1. Выведите на экран любое сообщение

```SQL
DO $$
BEGIN
    RAISE NOTICE 'Random message';
END; $$;
```

### №2. Выведите на экран текущую дату

```SQL
SELECT CURRENT_DATE;
```

### №3. Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран.

```SQL
DO $$
DECLARE
   a integer := 10;
   b integer := 20;
BEGIN
    RAISE NOTICE'a + b =  %', a+b;
	RAISE NOTICE'a - b =  %', a-b;
	RAISE NOTICE'a * b =  %', a*b;
END $$;
```

### №4. Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.

```SQL
DO $$
DECLARE
   a integer := 4;
BEGIN
	IF a = 5 THEN
		RAISE NOTICE 'Отлично';
	END IF;
	IF a = 4 THEN
	RAISE NOTICE 'Хорошо';
	END IF;
	IF a = 3 THEN
	RAISE NOTICE 'Удовлетворительно';
	END IF;
	IF a = 2 THEN
	RAISE NOTICE 'Неуд';
	END IF;
END $$;
```

```SQL
DO $$
DECLARE
   a integer := 3;
BEGIN
	CASE a
		WHEN 5 THEN RAISE NOTICE 'Отлично';
		WHEN 4 THEN RAISE NOTICE 'Хорошо';
		WHEN 3 THEN RAISE NOTICE 'Удовлетворительно';
		WHEN 2 THEN RAISE NOTICE 'Неуд';
	END CASE;
END $$;
```

### №5. Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).

```SQL
DO $$
DECLARE
   k integer := 20;
BEGIN
	LOOP
    	RAISE NOTICE '%', k*k;
		k := k + 1;
		IF k > 30 THEN
        	EXIT;
    	END IF;
END LOOP;
END $$;
```

```SQL
DO $$
DECLARE
   k integer := 20;
BEGIN
	WHILE k <= 30 LOOP
		RAISE NOTICE '%', k*k;
		k := k + 1;
	END LOOP;
END $$;
```

```SQL
DO $$
BEGIN
	FOR i IN 20..30 LOOP
		RAISE NOTICE '%', i*i;
	END LOOP;
END $$;
```

### №6. Написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.

```SQL
CREATE OR REPLACE FUNCTION Collatz(num int)
	RETURNS integer AS
	$$
        DECLARE
			chislo int := num;
			counter int := 0;
        BEGIN
			WHILE chislo != 1 LOOP
				IF chislo % 2 = 0 THEN
					chislo := chislo/2;
				ELSE
					chislo := chislo*3 + 1;
				END IF;
				counter := counter + 1;
			END LOOP;
        	RETURN counter;
   		END;
		$$ LANGUAGE plpgsql;
```

```SQL
CREATE OR REPLACE PROCEDURE CollatzProcedure(num int)
	AS $$
        DECLARE
			chislo int := num;
			counter int := 0;
        BEGIN
			WHILE chislo != 1 LOOP
				RAISE NOTICE '%', chislo;
				IF chislo % 2 = 0 THEN
					chislo := chislo/2;
				ELSE
					chislo := chislo*3 + 1;
				END IF;
			END LOOP;
   		END;
		$$ LANGUAGE plpgsql;
```

### №7. Написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел.

```SQL
CREATE OR REPLACE FUNCTION Luke(num int)
	RETURNS integer AS
	$$
        DECLARE
			chislo int := num;
        BEGIN
			IF chislo = 1 THEN
				RETURN 2;
			END IF;
			IF chislo = 2 THEN
				RETURN 1;
			END IF;
			RETURN Luke(chislo-1) + Luke(chislo-2);
   		END;
	$$ LANGUAGE plpgsql;
```

```SQL
CREATE OR REPLACE PROCEDURE LukeProcedure(num int)
	AS $$
        DECLARE
			chislo int := num;
			L0 int := 2;
			L1 int := 1;
			tmp int;
        BEGIN
		RAISE NOTICE '%', L0;
		RAISE NOTICE '%', L1;
		FOR i IN 0..num-3 LOOP
			tmp := L0;
			L0 := L1;
			L1 := tmp + L1;
			RAISE NOTICE '%', L1;
		END LOOP;
   		END;
	$$ LANGUAGE plpgsql;
```

### №8. Напишите функцию, которая возвращает количество человек родившихся в заданном году.

```SQL
CREATE OR REPLACE FUNCTION countBorn(yr int)
RETURNS int AS $$
DECLARE
	cnt int;
BEGIN
	SELECT count(*) INTO cnt
	FROM people
	WHERE EXTRACT(YEAR FROM people.birth_date) = yr;
	RETURN cnt;
END
$$ LANGUAGE plpgsql;
```

### №9. Напишите функцию, которая возвращает количество человек с заданным цветом глаз.

```SQL
CREATE OR REPLACE FUNCTION countEyes(color varchar)
RETURNS int AS $$
DECLARE
	cnt int;
BEGIN
	SELECT count(*) INTO cnt
	FROM people
	WHERE people.eyes = color;
	RETURN cnt;
END
$$ LANGUAGE plpgsql;
```

### №10. Напишите функцию, которая возвращает ID самого молодого человека в таблице.

```SQL
CREATE OR REPLACE FUNCTION youngest()
RETURNS int AS $$
DECLARE
	person int;
BEGIN
	SELECT people.id INTO person
	FROM people
	WHERE birth_date = (SELECT max(birth_date) FROM people);
	RETURN person;
END
$$ LANGUAGE plpgsql;
```

### №11. Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного. ИМТ = масса в кг / (рост в м)^2.

```SQL
CREATE OR REPLACE PROCEDURE BMI(imt int)
AS $$
DECLARE
	pRT people%ROWTYPE;
BEGIN
	FOR pRT IN SELECT * FROM people
		LOOP
			IF pRT.weight / (pRT.growth/100)^2 > imt THEN
				RAISE NOTICE 'id: %, name: %, surname: %', pRT.id, pRT.name, pRT.surname;
			END IF;
		END LOOP;
END
$$ LANGUAGE plpgsql;
```

### №12. Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми. Код должен быть представлен в виде транзакции (Например (добавление атрибута): BEGIN; ALTER TABLE people ADD COLUMN leg_size REAL; COMMIT;). Дополните БД данными.

```SQL
BEGIN;

CREATE TABLE parent_child(
	people_id int REFERENCES people(id),
	child_id int REFERENCES people(id)
);

INSERT INTO parent_child (people_id, child_id)
VALUES (1, 3),
(2, 3),
(4, 1);

COMMIT;
```

### №13. Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством.

```SQL
CREATE OR REPLACE PROCEDURE addPeopleAndKinship
(name varchar, surname varchar, birth_date date, growth real, weight real, eyes varchar, hair varchar, child_id int, parent1_id int, parent2_id int)
AS $$
DECLARE
	person_id int;
BEGIN
	INSERT INTO people (name, surname, birth_date, growth, weight, eyes, hair)
	VALUES (name, surname, birth_date, growth, weight, eyes, hair) RETURNING id INTO person_id;
	INSERT INTO parent_child (people_id, child_id)
	VALUES (person_id, child_id);
	INSERT INTO parent_child (people_id, child_id)
	VALUES (parent1_id, person_id);
	INSERT INTO parent_child (people_id, child_id)
	VALUES (parent2_id, person_id);
END
$$ LANGUAGE plpgsql;

CALL addPeopleAndKinship('Vladimir', 'Leonov', '09.04.1975', 178.3, 78.3, 'blue', 'black', 4, 1, 2)
```

### №14. Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека (выполнить также, как п.12).

```SQL
BEGIN;
ALTER TABLE people
ADD time_of_relevance timestamp NOT NULL DEFAULT NOW();
COMMIT;
```

### №15. Напишите процедуру, которая позволяет актуализировать рост и вес человека.

```SQL
CREATE OR REPLACE PROCEDURE updateGrowthAndWeight(person_id int, newGrowth real, newWeight real)
LANGUAGE plpgsql
AS $$
BEGIN
	UPDATE people
	SET growth = newGrowth, weight = newWeight, time_of_relevance = NOW()
	WHERE people.id = person_id;
END
$$;
```
