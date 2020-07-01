<!DOCTYPE html>

<html>
<head>
  <title>用户列表</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">

  </style>
  <script src="/static/js/reload.min.js"></script>
</head>

<body>
  <header>
    <h2>报名参与</h2>

    <form action="/user" enctype="application/x-www-form-urlencoded" method="post">
      <label for="text">手机号:</label><br>
      <input type="input" id="phone" name="phone"><br>
      <label for="desc">征文:</label><br>
      <input type="text" id="desc" name="desc"><br><br>
      <label for="captcha">验证码:</label><br>
      <input type="text" id="captcha" name="captcha"><span>获取手机验证码</span><br>
      <input type="submit" value="报名">
    </form>
  </header>
</body>
</html>
