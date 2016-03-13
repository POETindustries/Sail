package fallback

// NOTFOUND404 is a very basic web page signaling a 404 error.
// It contails the bare minimum necessary for a syntactically correct html web
// page and is used in those cases when not even basic database connections
// and templates work. The cms cannot be considered functional should that
// happen, and this markup at least tells the user as much. The markup is as
// generic as possible while still being somewhat good looking.
const NOTFOUND404 = `
<!doctype html>
<html>
	<head><title>Sorry About That</title><meta charset="utf-8"></head>
	<body style="background:black;text-align:center;color:white;padding:72px;font-size:1.5em;">
		<p style="font-size:2em;">Sorry About That!</p>
		<p>PAGE NOT FOUND</p>
	</body>
</html>`
