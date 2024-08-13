<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<hr>
<table border="1" cellspacing="0" width="800">
    <tr>
        <th>姓名</th>
        <th>性别</th>
        <th>手机号</th>

    </tr>
    <c:forEach items="${Users}" var="User">
        <tr>
            <td>${User.userName}</td>
            <td>${User.sex}</td>
            <td>${User.phoneNumber}</td>
        </tr>
    </c:forEach>
</table>
</body>
</html>