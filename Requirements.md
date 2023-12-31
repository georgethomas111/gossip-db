# Who is this for?

When you build systems in a serverless envvironment, you need a data-base that corresponds to it. With serverless, you bring code closer to the user. 

# What is a typical severless app?

    --> Server near user 
Req                       ---> Get user data, Get app specific data --> Call a bunch of APIs based on users actions/update them
    --> Server near user


# With this, we are converting the get userdata and App specific data to be API's that call the data base. 

# Hypothetical example

eg blackboard available for people to share and interact with. 
User goes to the link, writes something This can be saved to a data base. Each user in different parts of the world sees the same information

In traditional databases the data is stored in a bunch of systems that might be in a cluster somewhere around the work. They could be consistent(one update to a row at a time) or eventually consitent(may have multiple updates to a row) but mostly require a leader to make sure the data is synchronizing. 

If the data is stored in a system like that, the advantages of using serveless is lost as the data-base is going to add latency. Building serverless applications that are backed up by a database that itself is serverless helps to reduce latency. The writes to one location always get replicated around the world and the reads always return almost real time information of what is happening around the world.

It should be possible to have persistence.

Some of the metrics to think of for a data base

* data that can be stored - 10 Mb max to start with for one database
* Number of writes allowed per minute globally. 
* Max age of data to be saved - 30minutes
* Abitity to specify the behavior when conflict happens.

Proof of Concept

Send requests to different locations around the world and draw a graph of the requests per second using the data that gets written.
