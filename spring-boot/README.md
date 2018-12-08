# Hello Goose - Java - Spring Boot

This is a modern java version. To use it, run `mvn clean package ; cf push` - you need Maven.

Java isn't able to run the UNIX gethostname() function call but handily CF sets an INSTANCE_GUID environment variable.

N.B. Due to RAM requirements, this won't work on the free tier of PWS unless you reduce the number of instances.
