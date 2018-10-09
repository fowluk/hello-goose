# Hello Goose - Java Servlet Page

The simplest and perhaps least modern example possible with Java. This was state of the art on CF installs in 1999. That was a good year! The Java buildpack will supply a tomcat instance and deploy the WAR file into it. 

Note that we explicitly disable session management in the JSP. The gorouter supports session affinity for the JSESSIONID cookie, so we have to turn sessions off to make the loadbalancing fully round-robin and show the id changing. 

Java isn't able to run the UNIX gethostname() function call but handily CF sets an INSTANCE_GUID environment variable.

To use this, run `mvn package ; cf push`. You need maven.
