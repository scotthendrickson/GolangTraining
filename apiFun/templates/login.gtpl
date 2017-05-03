<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>This is a test</title>
</head>
<body>
    <h1 style="color:blue">Welcome {{.FirstName}}</h1>
    <form action="/login" method="post">
        First Name:<input type="text" name="firstName">
        Last Name:<input type="text" name="lastName">
        <input type="submit" value="Login">
    </form>
    <br>
    <br>
    <h3>{{.FirstName}}</h3>
    <h3>{{.LastName}}</h3>
</body>
</html>