# Hello Goose - Java Servlet Page

The simplest and perhaps least optimal example possible with Java. Some might say it's all a bit 1999, but then again, that was a good year! The Java buildpack will supply a tomcat instance and deploy the WAR file into it. 

Note that we explicitly disable session management in the JSP. The gorouter supports session affinity for the JSESSIONID cookie, so we have to turn sessions off to make the loadbalancing work and show the id changing. 

Java isn't able to run the UNIX gethostname() function call but handily CF sets an INSTANCE_GUID environment variable.

To use this, run `mvn package ; cf push`. You need maven.
