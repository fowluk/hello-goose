<%@ page session="false" %><html>
  <head>
    <title>Hello Goose</title>
  </head>
  <body>
    <h1>Hello Goose</h1>
    <img src="goose.jpg" width="50%"/><p/>
    <p>Here is my <strong>goose</strong></p>
    <p>Run it on the cloud for me</p>
    <p>I do not care how</p>
    <small>Instance id: <%= System.getenv("INSTANCE_GUID") %><br/><a href="https://github.com/jgjeffrey/hello-goose">github.com/jgjeffrey/hello-goose</a></small>
  </body>
</html>
