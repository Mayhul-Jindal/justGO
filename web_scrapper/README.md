# **web_scrapper**

## About
This is a basic web-scrapper used to scrape [devdungeon](https://www.devdungeon.com/).
> Note:  this website allows scrapping, not all websites permits such scrapping

## Functionality

1. Login to the website, source code in login.go 
   
![](./../img/8.jpg)

2. Checking the speed simulateously of multiple websites, for example sake:
    - https://www.google.com
    - https://www.youtube.com
    - https://www.facebook.com
    - https://www.amazon.com

![](./../img/7.jpg)

3. Custom cookie we made in the code following with the custom http header with our own user-agent because some sites may not allow the default one or they will ban it causing trouble
  
![](./../img/1.jpg)

4. Saving the http response body in an output.html file

![](./../img/5.jpg)

5. Title is scrapped from the output.html file using sub-string matching to find the element

![](./../img/2.jpg)

6. Comments are scrapped from output.html file using regular expressions

![](./../img/3.jpg)

7. Links from the page are scrapped using goquery

![](./../img/4.jpg)

8. Parsing complex url to retrieve any one from protocol, authority, host name, port number, path, query, filename, and reference

![](./../img/6.jpg)