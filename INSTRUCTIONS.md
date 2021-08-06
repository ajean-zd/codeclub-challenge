Zendesk Melbourne - Coding Challenge 
OVERVIEW 
Using the provided data (tickets.json and users.json and organization.json) write a simple command line application to search the data and return the results in a human readable format. 
● Feel free to use libraries or roll your own code as you see fit. However, please do not use a database or full text search product as we are interested to see how you write the solution. 
● Where the data exists, values from any related entities should be included in the results, i.e. searching organization by id should return its tickets and users. 
● The user should be able to search on any field, full value matching is fine (e.g. "mar" won't return "mary"). 
● The user should also be able to search for empty values, e.g. where description is empty. 
Search can get pretty complicated pretty easily, we just want to see that you can code a basic but efficient search application. Ideally, search response times should not increase linearly as the number of documents grows. You can assume all data can fit into memory on a single machine. 
EVALUATION CRITERIA 
We will look at your project and assess it for: 
1. Extensibility - separation of concerns. 
2. Simplicity - aim for the simplest solution that gets the job done whilst remaining readable, extensible and testable. 
3. Test Coverage - breaking changes should break your tests. 
4. Performance - should gracefully handle a significant increase in the amount of data provided. 
5. Robustness - should handle and report errors. 
6. Usability - Should provide installation instructions and how easy it is to use the
application 
7. General technical skills - Demonstrate proficiency in the chosen language and strong attention to details 
If you have any questions about these criteria please ask. 
SPECIFICATIONS 
● Use the language in which you are strongest. 
● Include a README with (accurate) usage instructions. 
● Document the assumptions and tradeoffs you’ve made. 
SUBMISSION 
GitHub is the preferred option (a public repository is fine) but we will also accept a .zip file if necessary. 
HINT 
One of the assessors wrote How To Pass A Coding Test.
SAMPLE OUTPUT 
Note: This output is not prescriptive, in fact we encourage you to do better. CLI - Display Results


meeting 1 notes.
how do we close the loop?
walking skeleton
tracer bullet - enough code to pretend it's working https://wiki.c2.com/?WalkingSkeleton

testing

step one. 
Load data.
take an existing file
read the file
parse the file

2 things we can do: 

1 load into memory and poke around
2 only load some of the data

start
take users
search by id
ask a question of each record (DO YOU MATCH?????)
if they match output the record as it is


