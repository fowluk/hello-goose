<%@ page session="false" %>
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Hello Goose !</title>
  <link rel="stylesheet" href="${pageContext.request.contextPath}/css/style.css">
</head>
<body>
<nav>
  <h1>Hello Goose ! </h1>
</nav>
<div>
  <img src="${pageContext.request.contextPath}/images/goose.jpg" width="100%"/>
</div>

<div>
  <blockquote>
    Here is my goose <br>
    Run it on the cloud for me <br>
    I do not care how
    <p>- James Jeffrey</p>
  </blockquote>
</div>
<div>
  <h4>Server sent information :</h4>
  <p>Instance id: <%= System.getenv("CF_INSTANCE_GUID") %></p>
  <p>Instance index: <%= System.getenv("CF_INSTANCE_INDEX") %></p>
</div>
<footer>
  <div>
    <a href="https://github.com/jgjeffrey/hello-goose">Contribute here and claim your Goose contributor badge</a>
  </div>
</footer>
</body>
</html>
