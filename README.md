# Initial Set-Up
  To start the project on your own machine you'll need Golang instlled, along with mongoDB.
  Once you've insured these requirements are met, download or clone this repo on your own machine
  and execute the following commands in order: 
  
  ```
  go install 
  ```
  
  One this command is successfully executed, you're ready to run the API.
  
  # Running the API
  To start the process, input the following command in terminal:
  
  ```
  go run . 
  ```
  
  This runs all the components in the go package (packgake main case), make sure the terminal 
  output shows the following: 
  
  ```
 Starting the application...
  ```
  
  Now the API is running on your machine on the local port 12345, next lets move onto testing out newly
  implemented API!
  
 # Testing
  We recommend you use Postman to input data into the dataset and compass to monitor your input data.
  Input data in the  formats of the sturctures provided in posts.go and user.go in the repo for posts and 
  users respectively.  
  To input data in user collection the link is: http://localhost/12345/user  
  To input data in post collection the link is: http://localhost/12345/post  
  To get user data out of user collection the link is: http://localhost/12345/user/{id} replace id by the user's ID you want.  
  To all data in user collection the link is: http://localhost/12345/users  
  To get post data out of user collection the link is: http://localhost/12345/posts/users/{id} replace id by the user's ID whose   post you want to see.  

  However, we have also done the unit testing outselves! The results we recieved from doing so are given as follows:
  
![image](https://user-images.githubusercontent.com/53595554/136667926-cc732aec-d4e9-4cb1-803f-d497b2781e56.png)

<hr>

![image](https://user-images.githubusercontent.com/53595554/136668130-3061836c-e53d-4aa9-9ac8-6210e1846a90.png)
