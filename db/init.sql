-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS expenses_id_seq;

-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL NOT NULL DEFAULT nextval('expenses_id_seq'::regclass),
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[],
        PRIMARY KEY ("id")
	);

INSERT INTO "expenses" ("title","amount","note","tags") VALUES ('strawberry smoothie', 79, 'night market promotion discount 10 bath', '{"food", "beverage"}');