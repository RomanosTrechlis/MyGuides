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



### *more is comming*



[printscreen]: servicemix_camel.png "servicemix console"