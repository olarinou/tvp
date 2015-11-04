
## Tarkvaraprojekt
### [Piletimüük](https://piletimyyk.herokuapp.com/)  

###  Build process and deployment

#### Deploying 
1. [Create a free Heroku account](https://signup.heroku.com/dc)
2. [Install Go Programming Language](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
3. [Add  $GOPATH/bin to your $PATH](http://www.computerhope.com/issues/ch000549.htm)
4. [Install Heroku Toolbelt](https://devcenter.heroku.com/articles/getting-started-with-go#set-up)
5. Now you can use the heroku command on your command shell.
6. Open your command shell and type `heroku login`, use the credentials you used in creating your Heroku account.
7. Next type `git clone https://github.com/projektsionist/piletimyyk.git` to clone the repo and `cd piletimyyk` to go to that directory.
8. Next step is to prepare Heroku. Type `heroku create` this creates a random app on Heroku, ready to receive your source code. 
9. `git push heroku master` to deploy your source code.
10. `heroku addons:create heroku-postgresql:hobby-dev` to create a database.
11. To visit the URL the app was created on, simply type `heroku open`.
