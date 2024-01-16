# Restructed repo

# Purpose of this Repository
This is close to production ready app for this requirement
Made below changes.
1. The original code was small so it might be an illusion that is readable currently. As and when more code will be added it will become lot less readable as well as there will be issues because of redundant code.

2. Instead of relying on base http for app management, used go-fiber, as it highly efficient and can handle much more apis more efficiently. Also it provides tons of functionality which would have to be added by boiler plate code if required. Also using it will make the app much more structured.

3. MYSQL connection url is loaded from env, instead of being hardcoded as it poses both a security thread as well in case of multiple deployment, we need different version of code.

4. The orinal code had multiple errors which were not handled properly. Eg. err.Error() was invoked without checking if err is nil or not, this could have resulted in returning wrong response to end user as well as unwanted scenarios.

5. Both the apis could be accessed by both POST and GET methods, which is highly discouraged. As the code was meant for POST, all the GET apis should fail ideally without even reaching the handler.

6. fmt.Println was used to log success and error cases : In production grade apps loggers should be used to log to a remote server.

7. In Docker file, entire ubuntu image was pulled which could be replaced by golang image thus making the application light weight.Also the dockerfile was incorrect as app was not a valid directory. so updated it according to latest project structure

8. The code was not testable by default as different components involved to form api are not testable individually. I have refactored the code to make it more testable (individual db calls can be tested, db connection establishment can be tested and api calls can be tested).

9. The code was not maintainable as existing project structure is not testable, not extensible, more complicated requirements will make the code much of a mess.

#BUILD COMMAND : docker build -t app .

#RUN COMMAND : docker run -m 100m -p 8080:8080 -e PORT=8080 -e MONGO_URI=mongodb://host.docker.internal:27017 -e APP_NAME=auth-core -e MONGO_DB_NAME=users -e JWT_SECRET='arSby6PGskbPG0fP89CvCccGyhvd_d9RqzRNin-YWmU=' auth-core:latest

docker build -t app .
docker run -e MYSQL_URI="root:password@tcp(host.docker.internal:3306)/test" -p 8080:8080 app