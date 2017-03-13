CREATE SCHEMA hospital;

CREATE SEQUENCE hospital.patients_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE hospital.patients (
    id   integer DEFAULT nextval('hospital.patients_seq'::regclass) NOT NULL,
    name text,
    weight integer,
    height integer,
    age integer,
    gender text,
    bustSize integer,
    waistSize integer,
    hipSize integer
);

ALTER TABLE ONLY hospital.patients
    ADD CONSTRAINT patient_pkey PRIMARY KEY (id);

