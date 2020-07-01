<!DOCTYPE html>

<html>
<head>
  <title>获奖名单</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">

  </style>
  <script src="/static/js/reload.min.js"></script>
</head>

<body>
  <header>
    <a href="/lottery">去抽奖</a> <a href="/export/lottery">下载中奖名单</a>
    <h2>名单：共{{.count}}份</h2>
    <div>
    <table>
        <tr>
            <td>用户</td>
            <td>奖品</td>
            <td>时间</td>
        </tr>
      <tbody>
          {{range $index, $row := .lotterys}}
          <tr>
              <td>{{$row.Phone}}</td>
              <td>{{$row.PrizeName}}</td>
              <td>{{$row.CreatedAt}}</td>
      	  </tr>
          {{end}}
      	</tbody>
      </table>
    </div>
  </header>
</body>
</html>
