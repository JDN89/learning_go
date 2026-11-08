# NOTES

## Redirects

Unfortunately, to achieve this you can't just send a regular 3xx response. The crux of the problem is that that browsers will automatically intercept and follow 3xx responses before HTMX has access to them — so HTMX never gets to see the 3xx response, only the final response after any redirects. It doesn't know that a redirect happened behind the scenes, and will just swap in the returned content like normal.

Instead, if you want something that behaves more like a regular redirect, you need to send a 2xx response along with the HX-Redirect header. 

## Checking if a request is coming from HTMX

This obviously isn't ideal. A much better approach would be to change the response that our searchUsers handler sends, depending on whether the request is coming from HTMX or not. Specifically:

    If the request is coming from HTMX, we should return an HTML partial that it can swap into the table, just like we already are.
    If the request is not coming from HTMX, we should return a full HTML page that contains the matching user details.

But you also need to handle the situation where the original request might not be coming from HTMX — especially if you are using progressive enhancement so that your application still works if JavaScript is disabled or HTMX doesn't load correctly. In that case, it's important to fallback to sending a regular 3xx response from your Go handler, rather than the 2xx response and HX-Redirect header.

Putting this together, I normally create a redirect() helper which leverages the isHTMXRequest() function we made earlier and looks like this:

## Aditional config

```html
<meta
    name="htmx-config"
    content='{
        "includeIndicatorStyles": false,
        "historyCacheSize": 0,
        "historyRestoreAsHxRequest": false,
        "responseHandling":[
            {"code":"204", "swap": false},
            {"code":"422", "swap": true},
            {"code":"[45]..", "swap": true, "target": "body"},
            {"code":"...", "swap": true}
        ],
        "timeout": 5000
    }'
>
```
