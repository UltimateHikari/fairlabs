CREATE DATABASE fairlabs;
CREATE USER fairlabs WITH PASSWORD 'fairlabs';

\c fairlabs

DROP TABLE IF EXISTS submits;
DROP TABLE IF EXISTS participants;
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
    course_id SERIAL UNIQUE,
    course_name text,
    university_group integer,
    tasks_amount integer NOT NULL, 
    algo integer REFERENCES algos (algo_id),
    cond_id integer REFERENCES conds (cond_id),
    cond_data text[],
    queue integer[] DEFAULT('{}'), 
    PRIMARY KEY (course_id, course_name, university_group)
);

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    email text NOT NULL,
    user_name text NOT NULL,
    group_number integer NOT NULL,
    is_admin boolean DEFAULT(FALSE)
);

--role == STUDENT/TEACHER
CREATE TABLE participants(
    course_id integer REFERENCES courses (course_id),
    user_id integer REFERENCES users (user_id),
    user_role text,
    user_goal integer DEFAULT(5),
    user_priority boolean DEFAULT(FALSE),
    PRIMARY KEY(course_id, user_id)
);

--mb a limit on task_id?
--status == PLANNED/FINISHED/CLEARED
CREATE TABLE submits(
    submit_id SERIAL PRIMARY KEY,
    course_id integer REFERENCES courses (course_id),
    user_id integer REFERENCES users (user_id),
    task_id integer,
    task_status text
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

INSERT INTO users (email, user_name, group_number, is_admin)
    VALUES ('a.rudometov@g.nsu.ru', 'Andrey Rudometov',19201, TRUE);
INSERT INTO users (email, user_name, group_number)
    VALUES ('a.ogneva@g.nsu.ru', 'Anastasia Ogneva', 19201);
INSERT INTO users (email, user_name, group_number)
    VALUES ('test@g.nsu.ru', 'test', 19212);

INSERT INTO courses(course_name, university_group, tasks_amount, queue)
    VALUES ('WackoCourse', 19201, 33, '{1,2}');
INSERT INTO courses(course_name, university_group, tasks_amount, queue)
    VALUES ('OkcawCourse', 19212, 14, '{3}');

INSERT INTO participants (course_id, user_id, user_role)
    VALUES (1, 1, 'STUDENT');
INSERT INTO participants (course_id, user_id, user_role)
    VALUES (1, 2, 'STUDENT');
INSERT INTO participants (course_id, user_id, user_role, user_priority)
    VALUES (2, 3, 'STUDENT', TRUE);

INSERT INTO submits(course_id, user_id, task_id, task_status)
    VALUES (2, 3, 1, 'FINISHED');
INSERT INTO submits(course_id, user_id, task_id, task_status)
    VALUES (2, 3, 2, 'FINISHED');
INSERT INTO submits(course_id, user_id, task_id, task_status)
    VALUES (2, 3, 3, 'PLANNED');
INSERT INTO submits(course_id, user_id, task_id, task_status)
    VALUES (2, 3, 4, 'PLANNED');

GRANT USAGE ON SCHEMA public TO fairlabs;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO fairlabs;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO fairlabs;




