Gomont Tutorial
===================


In this tutorial we will explain the steps to install both the agents and servers of Gomont. Gomont is divided in two parts, an agent, which will execute all monitoring scripts and send them to the server, and a server which will store all of the agent's information and provide an interface to the client's browser.

----------


Agent
--------------------
Using the Go toolkit we can easily install the agent with a simple 
`go get github.com/jybateman/gomont/agent`.
Once installed you will have to edit the `config.json` file which contains the `Username`, `Password` and `Port` all are strings that will be used by the Agent, this information is used so that the server can connect to the agent and start sending information.

    {
        "Port": "4242",
        "Username": "hello",
        "Password": "world"
    }

You are also given an `monitor.json` as a sample. This json file contains an array of  `Name`, `Command`, `Frequence` and `Graph`, which represent the following:

 - `Name` is a string that associates a name to the command that will be displayed on client's interface.
 - `Command` is a string that defines the command to be executed.  *Note: all command must write the result on the standard output which will be the value displayed on the client's interface, and must return an exit value of 0, 1 or 2 which represent the state OK, Warning and Danger respectively*
 - `Frequence`is a string that represents the frequency that the command will be executed, this can be useful for commands whose state does not change frequently like disk size, validtime units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
 - `Graph` is a boolean that defines if a command will be displayed as a graph or a table, this is useful for commands whose previous state is not important like CPU frequency.

    [{
        "Name": "Random",
        "Command": "./test.sh",
        "Frequence": "1s",
        "Graph": false
    
    },{
        "Name": "CPU",
        "Command": "./cpu.sh",
        "Frequence": "1s",
        "Graph": true
    }]

Server
--------------------
Like the Agent the Server can be installed by using the Go toolkit `go get github.com/jybateman/gomont/server`.
Once installed you will have to edit the `config.json` file which will contains `Port`, `Mysql` array containing`Port`, `IP`, `Username` and `Password`.

 - `Port` is a string that defines the port that the webserver will listen on so that the client can access the interface.
 - `Mysql` is an array that containes the following:
  - `Port` is a string that defines the port that MySQL will use to connect to the database.
  - `IP` is a string that defines the address that will be use by MySQL to connect to the database
  - `Username` is a string that MySQL will use as user to access the database. *Note: the user must have read and write privileges*
  - `Password` is a string that defines the password for `Username` so that MySQL can connect to the database.

  

      {
            "Port": "9000",
            "Mysql": {
                "Port": "3306",
                "IP": "127.0.0.1",
                "Username": "root",
                "Password": "helloworld"
            }
        }

Client
--------------------
Once the `Server` is running the can connect to the interface by typing in a web browser on the port defined in the server's `config.json`.
Once on the web interface with the default login page where you can either choose to login using an already existing account or create one.
![enter image description here](https://raw.githubusercontent.com/jybateman/gomont/master/screenshot/login.PNG)![enter image description here](https://raw.githubusercontent.com/jybateman/gomont/master/screenshot/signup.PNG)

Once the client is connected, the client is taken to the home page where all the servers added by the client are visible.
![enter image description here](https://raw.githubusercontent.com/jybateman/gomont/master/screenshot/home+server.PNG)
The client can now view all the servers that are being monitored. And can choose to delete a server, edit one of the monitored servers or add a new server to monitor.
![enter image description here](https://raw.githubusercontent.com/jybateman/gomont/master/screenshot/add_server.PNG)

The client can also get a detailed view of the server being monitored by clicking on the server name.
![enter image description here](https://raw.githubusercontent.com/jybateman/gomont/master/screenshot/server_detail_view.PNG)

Scripts
--------------------
The goal of Gomont is to provide users with an intuitive and simple interface, and the possibility to easily monitor anything in any language without having to use any external library or any binding. To achieve this Gomont simply reads the standard output and the return value of the script. 
In the agent folder the user is provided with a couple of monitoring script examples.

    #!/bin/bash
    echo -n $RANDOM
    exit `expr $RANDOM % 3` 
In this the example a random number is printed on the standard output and a random number between 0 and 2 is return to represent the three states `OK`, `Warning` and `Danger`
