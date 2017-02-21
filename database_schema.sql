CREATE SCHEMA animals;

CREATE SEQUENCE animals.dog_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE animals.dogs (
    id   integer DEFAULT nextval('animals.dog_seq'::regclass) NOT NULL,
    breed text,
    weight integer,
    height integer,
);

ALTER TABLE ONLY animals.dogs
    ADD CONSTRAINT dog_pkey PRIMARY KEY (id);

