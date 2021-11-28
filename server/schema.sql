DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS algos;
DROP TABLE IF EXISTS conds;

CREATE TABLE algos(
    algo_id integer PRIMARY KEY,
    algo_name text NOT NULL
);

CREATE TABLE conds(
    cond_id integer PRIMARY KEY,
    cond_name text NOT NULL,
    example_data text[]
);

--todo add tasks count
CREATE TABLE courses(
    course_id SERIAL PRIMARY KEY,
    course_name text NOT NULL,
    university_group integer NOT NULL,
    algo integer REFERENCES algos (algo_id),
    cond_id integer REFERENCES conds (cond_id)
);

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    email text NOT NULL,
    user_name text NOT NULL,
    is_admin boolean DEFAULT(FALSE)
);

--role == STUDENT/TEACHER
CREATE TABLE groups(
    group_number integer,
    user_number integer REFERENCES users (user_id),
    user_role text,
    PRIMARY KEY(group_number, user_number)
);

INSERT INTO algos (algo_id, algo_name) 
    VALUES(1, 'alphabet-descending');
INSERT INTO algos (algo_id, algo_name) 
    VALUES(2, 'datetime-filo');
INSERT INTO algos (algo_id, algo_name) 
    VALUES(3, 'persentage-ascending');

INSERT INTO conds (cond_id, cond_name, example_data) 
    VALUES(1, 'deadlines', 
    '{"21.09:1", "15.10:2", "29.10:3", "27.11:4", "22.11:5"}');
INSERT INTO conds (cond_id, cond_name, example_data) 
    VALUES(2, 'task_amounts', 
    '{"6:3", "11:4", "23:5"}');

INSERT INTO users (user_id, email, user_name, is_admin)
    VALUES ('a.rudometov@g.nsu.ru', 'Andrey Rudometov', TRUE);
INSERT INTO users (user_id, email, user_name)
    VALUES ('a.ogneva@g.nsu.ru', 'Anastasia Ogneva');

INSERT INTO groups (group_number, user_number, user_role)
    VALUES (19201, 1, 'STUDENT');
INSERT INTO groups (group_number, user_number, user_role)
    VALUES (19201, 2, 'STUDENT');

INSERT INTO courses(course_name, university_group)
    VALUES ('WackoCourse', 19201);
INSERT INTO courses(course_name, university_group)
    VALUES ('OkcawCourse', 19212);







