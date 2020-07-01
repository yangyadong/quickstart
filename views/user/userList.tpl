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
    <a href="/participate">我要参与</a>
    <h2>列表：共征集{{.count}}份</h2>
    <div>
    <table>
        <tr>
            <td>用户</td>
            <td>征文</td>
        </tr>
      <tbody>
          {{range $index, $row := .users}}
          <tr>
              <td>{{$row.Phone}}</td>
              <td>{{$row.Desc}}</td>
      	  </tr>
          {{end}}
      	</tbody>
      </table>
    </div>
  </header>
</body>
</html>
