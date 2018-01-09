package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/johanbrandhorst/protobuf/grpcweb/status"
	"honnef.co/go/js/dom"

	"github.com/johanbrandhorst/demo/proto/client"
)

// Build this snippet with GopherJS, minimize the output and
// write it to html/frontend.js
//go:generate gopherjs build frontend.go -m -o html/frontend.js

// Integrate generated JS into a Go file for static loading.
//go:generate bash -c "go run assets_generate.go"

// This constant is very useful for interacting with
// the DOM dynamically
var document = dom.GetWindow().Document().(dom.HTMLDocument)

// Define no-op main since it doesn't run when we want it to
func main() {}

// Ensure our setup() gets called as soon as the DOM has loaded
func init() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

// Setup is where we do the real work.
func setup() {
	// This is the address to the server, and should be used
	// when creating clients.
	serverAddr := strings.TrimSuffix(document.BaseURI(), "/")

	// TODO: Use functions exposed by generated interface
	cli := client.NewBackendClient(serverAddr)

	target := document.GetElementByID("target").(*dom.HTMLDivElement)
	populated := false

	button := document.GetElementByID("button").(*dom.HTMLButtonElement)
	button.AddEventListener("click", false, func(_ dom.Event) {
		if populated {
			target.SetInnerHTML("")
			button.SetTextContent("Get User")
			populated = false
			return
		}

		go func() {
			user, err := cli.GetUser(context.Background(), &client.UserRequest{
				Id: "abcd",
			})
			if err != nil {
				st := status.FromError(err)
				target.SetInnerHTML("<b>" + st.Error() + "</b>")
				return
			}

			target.SetInnerHTML(`<p>User ` + user.GetName() + ` is ` + strconv.Itoa(int(user.GetAge())) + ` years old</p>`)
			button.SetTextContent("Clear user")
			populated = true
		}()
	})
}
