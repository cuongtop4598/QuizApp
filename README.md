Build a website that supports Quiz with Go Gin & HTTP2: 

1. Support user registration function, login is required to perform the survey. 
	
2. Once logged in, the first page displays the first question and the total number of questions to be answered.
	
3. Click save and next to save the results and go to the next page, back to return to the previous question (Optional: Support skip question).
	
4. Questions can choose 1 or more answers.
	
5. Show a list of selected questions and answers along with the correct answer and total score after completing the survey. 
	
6. Use Token Based Authentication with access token, refresh token. 
	
7. Write Middleware for Authentication/Authorization actions. 
	
8. Implement audit log for all actions. 
	
9. Use MongoDB as database, can use any ORM for query.


Mongodb error. 

Just do those two commands for temporary solution:

$ sudo rm -rf /tmp/mongodb-27017.sock

$ sudo service mongod start

For details:

That shall be fault due to user permissions in .sock file, You may have to change the owner to monogdb user.

chown -R mongodb:mongodb /var/lib/mongodb
chown mongodb:mongodb /tmp/mongodb-27017.sock