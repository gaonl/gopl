<html>
<head>
	<title>上传文件</title>
</head>
<body>
	<form enctype="multipart/form-data" action="http://127.0.0.1:8181/upload" method="post">
		<input type="file" name="uploadfile"/>
		<input type="text" name="msg"/>
		<input type="submit" value="upload"/>
	</form>
</body>
</html>