build:
	docker build -t shubham1010/devops-assignment-2 .
	
run: build
	docker run -itd --name simple-mysql-example -p 3306:3306 shubham1010/devops-assignment-2

exec:
	docker exec -it simple-mysql-example /bin/bash	

clean:
	docker stop simple-mysql-example; docker rm simple-mysql-example; docker rmi shubham1010/devops-assignment-2

.PHONY: build run clean exec
