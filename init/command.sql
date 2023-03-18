CREATE DATABASE school_management_db;

USE school_management_db;

CREATE TABLE students (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)  NOT NULL,
    sex VARCHAR(10) NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP
);

CREATE TABLE courses (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    unit INT NOT NULL,
    semester INT NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    UNIQUE KEY courseInfo (title, unit, semester),
    CHECK (unit BETWEEN 0 AND 5 AND semester BETWEEN 1 AND 2)
);

CREATE TABLE student_courses (
    studentId INT NOT NULL,
    courseId INT NOT NULL,
    FOREIGN KEY (studentId) REFERENCES students(id) ON DELETE CASCADE,
    FOREIGN KEY (courseId) REFERENCES courses(id) ON DELETE CASCADE,
    PRIMARY KEY(studentId, courseId)
);

SHOW TABLES;
DESC students;
DESC courses;
DESC student_courses;
