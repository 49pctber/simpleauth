# SimpleAuth

This is a simple HTTP authentication library that uses JSON Web Tokens.

Use at your own risk!

## How to Use

### Getting started

First, install the command line utility:

``` bash
go install github.com/49pctber/simpleauth/cmd/simpleauth@latest
```

Then in your project folder, initialize the user file with `simpleauth`.
This will create a `simpleauth.json` file that will store the users, their password hashes, and a secret used for signing JSON Web Tokens.
Keep this file secret!

Next, create an admin user using the following command:

``` bash
simpleauth user add -a
```

and enter the username and password when prompted. You can create a regular user by omitting the `-a` flag.

### Using simpleauth

In your Go project, import `github.com/49pctber/simpleauth`.
Let your program know where to find the `simpleauth.json` file you created above using `simpleauth.Configure(path_to_file)`.

Then you can configure your endpoints that require authentication by using `simpleauth.RequireAuthentication` as shown below.

``` go
mux := http.NewServeMux()
mux.Handle("/", simpleauth.RequireAuthentication(http.HandlerFunc(homeHandleFunc), false))
mux.Handle("/admin", simpleauth.RequireAuthentication(http.HandlerFunc(homeHandleFunc), true))

fmt.Println("Serving on :8080")
fmt.Println(http.ListenAndServe(":8080", mux))
```

The boolean argument of `simpleauth.RequireAuthentication` indicates whether the user needs to be an administrator to access the endpoint.

If a user visits an endpoint and is not logged in, they will be presented with a log in screen.
If they successfully authenticate (and have administrator permissions, if applicable), they will be taken to their desired endpoint.
Otherwise, they will be presented with a log in screen.
