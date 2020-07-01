<!DOCTYPE html>

<html>
<head>
  <title>抽奖</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">

  </style>
  <script src="/static/js/reload.min.js"></script>
</head>

<body>
  <header>
        <a href="/lottery/list">查看中奖名单</a>
        <a href="/user">查看参与者</a>
        <form action="/lottery" enctype="application/x-www-form-urlencoded" method="post">
          <label for="text">手机号:</label><br>
          <input type="input" id="phone" name="phone"><br>
          <input type="submit" value="抽奖">
        </form>
  </header>
</body>
</html>
