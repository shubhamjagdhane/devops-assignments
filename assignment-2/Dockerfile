FROM mysql

ENV MYSQL_ROOT_PASSWORD=root 
ENV MYSQL_DATABASE=devops-assignment-1 
ENV MYSQL_USER=shubham 
ENV MYSQL_PASSWORD=jagdhane

ADD students.sql /docker-entrypoint-initdb.d

EXPOSE 3306
