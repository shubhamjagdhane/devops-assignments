CREATE TABLE studentData(name varchar(30) not null, rollno integer unique,  address text not null, mobile_number numeric(10, 0) not null, pan_number varchar(10) not null, PRIMARY KEY(rollno));

INSERT INTO studentData VALUES("Shubham Jagdhane", 18116, "Pune", 1234567890, "ABCDE12345");
