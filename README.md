# Restructed repo

# Purpose of this Repository
This is close to production ready app for this requirement
Made below changes.
1. The original code was small so it might be an illusion that is readable currently. As and when more code will be added it will become lot less readable as well as there will be issues because of redundant code.

2. Instead of relying on base http for app management, used go-fiber, as it highly efficient and can handle much more apis more efficiently. Also it provides tons of functionality which would have to be added by boiler plate code if required. Also using it will make the app much more structured.

3. MYSQL connection url is loaded from env, instead of being hardcoded as it poses both a security thread as well in case of multiple deployment, we need different version of code.

4. The orinal code had multiple errors which were not handled properly. Eg. err.Error() was invoked without checking if err is nil or not, this could have resulted in returning wrong response to end user as well as unwanted scenarios.

5. Both the apis could be accessed by both POST and GET methods, which is highly discouraged. As the code was meant for POST, all the GET apis should fail ideally without even reaching the handler.

6. Ensure in db, the table is already created. If not it will throw error. I manually created table for testing, though this process can be automated by including .sql scripts in the go project.

7. fmt.Println was used to log success and error cases : In production grade apps loggers should be used to log to a remote server.

8. In Docker file, entire ubuntu image was pulled which could be replaced by golang image thus making the application light weight.Also the dockerfile was incorrect. so updated it according to latest project structure and its working.

9. The code was not testable by default as different components involved to form api are not testable individually. I have refactored the code to make it more testable (individual db calls can be tested, db connection establishment can be tested and api calls can be tested).

10. The code was not maintainable as existing project structure is not testable, not extensible, more complicated requirements will make the code much of a mess.

11. Use below command to create a docker image of mysql
    a. docker pull mysql:latest
    b. docker run -d --mysql-container -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -p 3306:3306 mysql:latest

12.  Use below command to create and run a docker image of this app
    a. docker build -t app .
    b. docker run -e MYSQL_URI="root:password@tcp(host.docker.internal:3306)/test" -p 8080:8080 app

