*This is CLI application*

--main.go--

3 imported packages, fmt, sync and time,

4 pacakage level variables

1 custom data struct

Thread implementation

5 function implemented out of main scope to dont repeat yourself principle and to make code clean

user input is from console

user validation implemented

user information stored in custom struct and appended to slice

program will end then all threads will be executed

--helper.go--

1 import of strings

has 1 function, shares sames package variables,

the main pupose of helper.go is that it validates user input

it mean that its check is it VALID NAME it has to have more then 2 char, checks is Valid Email,
it just checks if it contains in that string "@" CHAR. and check is ticket more then left or equal, that
user couldnt order more tickets that exist for this event.

and it return these validations as boolean type true or false.


*In progress*
