# ServiceMix 7.x Getting Started

## Starting ServiceMix
To start ServiceMix we navigate to %KARAF_HOME% and execute the `bin\servicemix` command.

## Deploy a new camel blueprint
In order to deploy a camel route blueprint we need to save the blueprint in an xml format inside the `%KARAF_HOME%\deploy` folder.
We save the following xml.
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<route customId="false" id="route2" xmlns="http://camel.apache.org/schema/spring">
    <from uri="file:C:\\camel\\in"/>
    <log message="Moving ${file:name} to the output directory" id="log2"/>
    <to uri="file:C:\\camel\\out" id="to2"/>
</route>
```
Inside the servicemix console type `route-list` to see if the blueprint was correctly loaded. Then type `route-info <route_name>` to see statistcs about the route and `route-show <route_name>` to see the content of the route.

![servicemix console][printscreen]


## Logging
Set the logging level using the command `log:set <LOGGING_LEVEL>`. To see the logs type `log:display`. Because the logs are too many to effectively handle use the pipe the grep command as well using it like this: `log:display | grep <route_name>`.


## Starting and Stoping bundles
To examine the bundles execute the `bundle:list` command. The first column represents the bundle id. You can stop bundles by their id using the `bundle:stop <bundle_id>` command. Similarly, we can start a bundle using the `bundle:start <bundle_id>` command.


## Webconsole
Install the webconsole by running the `feature:install webconsole`. Then we navigate to http://localhost:8181/system/console  and we enter smx/smx as username and password.
We can uninstall the webconsole using the `feature:uninstall webconsole` command.


## Hawtio
We now will install hawtio console. First we need to add the hawt repository to service mix. We do that by typing the `feature:repo-add hawtio 1.4.65` command. Then we can install it using the `feature:install hawtio` command. The hawtio web console.


## Creating a simple bundle
This bundle is quite simple. It contains a route that download a webpage and write it inside a folder. This project is build with maven in NetBeans IDE and it produces a .jar file that is then deployed to ServiceMix.

### The project structure
```
* ServiceMixExample1 OSGi Bundle
    * src
        * main
            * java
                * com.romanostrechlis.servicemix.example1
                    * StockProcessor.java
            * resources
                * OSGI-INF
                    * blueprint
                        * blueprint.xml
    * nbactions.xml
    * pom.xml
```

### The project
The pom contains the dependencies and the build rules to create the .jar file. The `StockProcessor.java` is a processor that for now does nothing. And then there is the `blueprint.xml` which contains the processor creation as a spring bean the camel context and the route in spring DSL. We can explore all these files in `examples\ServiceMixExample1` folder.

### Running the project
First we need to install camel-http feature in servicemix in order to get the webpage. We do that by executing the `feature:install camel-http` command. This dependency is part of the servicemix. If it wasn't, we would heve to deploy the .jar dependency.
Then we are ready to clean and build our example and deploy the produced .jar file in the deploy folder of servicemix. Running the `bundle:list | grep Example` command we can see the status of our deployment. If by chance the status is "GracePeriod" then execute the `bundle:diag` command to see what is wrong.
If everything is ok then a text file must be created inside the c:\camel\out folder containg the webpage.


## Getting Twitter Messages to an ActiveMQ queue
For this route we will connect to twitter API an get some tweets. First we need to install the corresponding camel dependency by executing the command `feature:install camel-twitter`. Then we create the blueprint in spring DSL.
```xml
<blueprint
    xmlns="http://www.osgi.org/xmlns/blueprint/v1.0.0"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="
      http://www.osgi.org/xmlns/blueprint/v1.0.0
      http://www.osgi.org/xmlns/blueprint/v1.0.0/blueprint.xsd">
 
    <camelContext xmlns="http://camel.apache.org/schema/blueprint">
        <route id="get_tweets">
            <from uri="twitter://timeline/home?type=polling&amp;delay=900000&amp;consumerKey=KEY&amp;consumerSecret=KEY&amp;accessToken=KEY&amp;accessTokenSecret=KEY" />
            <convertBodyTo type="java.lang.String"/>
            <to uri="activemq:queue:twitter" />
        </route>
    </camelContext>
</blueprint>
``` 
This blueprint sends tweets to an ActiveMQ queue called twitter which is automatically created if it doesn't exists.
The convertBodyTo element is required for getting the body's format recognized.






[printscreen]: img/servicemix_camel.png "servicemix console"