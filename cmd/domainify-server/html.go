package main

const (
	htmlHead = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="description" content="Domainify turns phrases into domains">
	<meta name="author" content="Domainify Conglomerate">

	<!-- TODO: make favicons and put in s3

	<link rel="apple-touch-icon" sizes="57x57" href="/favicons/apple-icon-57x57.png">
	<link rel="apple-touch-icon" sizes="60x60" href="/favicons/apple-icon-60x60.png">
	<link rel="apple-touch-icon" sizes="72x72" href="/favicons/apple-icon-72x72.png">
	<link rel="apple-touch-icon" sizes="76x76" href="/favicons/apple-icon-76x76.png">
	<link rel="apple-touch-icon" sizes="114x114" href="/favicons/apple-icon-114x114.png">
	<link rel="apple-touch-icon" sizes="120x120" href="/favicons/apple-icon-120x120.png">
	<link rel="apple-touch-icon" sizes="144x144" href="/favicons/apple-icon-144x144.png">
	<link rel="apple-touch-icon" sizes="152x152" href="/favicons/apple-icon-152x152.png">
	<link rel="apple-touch-icon" sizes="180x180" href="/favicons/apple-icon-180x180.png">
	<link rel="icon" type="image/png" sizes="192x192"  href="/favicons/android-icon-192x192.png">
	<link rel="icon" type="image/png" sizes="32x32" href="/favicons/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="96x96" href="/favicons/favicon-96x96.png">
	<link rel="icon" type="image/png" sizes="16x16" href="/favicons/favicon-16x16.png">
	<link rel="manifest" href="/manifest.json">

	-->

	<meta name="theme-color" content="#ffffff">
	<meta name="application-name" content="Domainify"/>
	<meta property="og:title" content="Domainify" />
	<meta property="og:description" content="Domainify turns phrases into domains" />	

	<title>Domainify</title>

	<link rel="stylesheet" type="text/css" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
`

	htmlHome = htmlHead +
		`</body>

<div class="container">
	<h1>Domainify</h1>
	<p>Enter your phrase here:</p>
	<input id="searchterm"/>
	<a id="reflectedlink" class="btn btn-default" href="" role="button">Go</a>
</div>

<script type="text/javascript">
    var link= document.getElementById('reflectedlink');
    var input= document.getElementById('searchterm');
    input.onchange=input.onkeyup= function() {
        link.pathname = encodeURIComponent(input.value);
    };
</script>

</body>
</html>
`

	htmlListDomains = htmlHead +
		`</body>

<div class="container">
	<h1>Domainify</h1>
	<p>Enter your phrase here:</p>
	<input id="searchterm"/>
	<a id="reflectedlink" class="btn btn-default" href="" role="button">Go</a>
	<br><br>
	<p>Domains:</p>
	{{range .}}
	<p>{{.}}</p>
	{{end}}
</div>

<script type="text/javascript">
    var link= document.getElementById('reflectedlink');
    var input= document.getElementById('searchterm');
    input.onchange=input.onkeyup= function() {
        link.pathname = encodeURIComponent(input.value);
    };
</script>

</body>
</html>
`
)
