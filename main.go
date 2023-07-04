package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site!<h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, 
	email me at <a href=\"mailto:jon@calhoun.io\>jon@calhoun.io</a>.`)
}

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	fmt.Fprint(w, ` <h1>FAQ Page</h1>
// 				    <p>Q: Is there a free version?
// 					<p>A: Yes! We offer a free trial for 30 days on any paid plans.
// 					<p>Q: What are your support hours?
// 					<p>A: We have support sraff answering emails 24/7, though response times
// 						  may be a bit slower on weekends.
// 					<p>Q: How do i contact support?
// 					<p>A: Email us - <a href=\"mailto:support@lenslocked.com\>support@lenslocked.com</a>.`)
// }

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, ` <h1>FAQ Page</h1>
<ul>
	<li><b>Is there a free version?</b>Yes! We offer a free trial for 30 days on any paid plans.</li>
	<li><b>What are your support hours?</b>We have support sraff answering emails 24/7, though response times 
	may be a bit slower on weekends.</li>
	<li><b>How do i contact support?</b>Email us - 
	<a href="mailto:support@lenslocked.com">support@lenslocked.com</a>.</li>
</ul>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	http.ListenAndServe(":3000", router)
}
