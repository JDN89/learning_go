# Resources

[go/template package](https://pkg.go.dev/html/template)
[web examples](https://gowebexamples.com/?utm_source=chatgpt.com)
[htmx documentation](https://htmx.org/docs/)
[[Let's go further expansion distributed system]]
## Recommended project setup

For HTMX with Go's `net/http` and `html/template`, use a lightweight MVC-style
architecture. Keep HTTP concerns in handlers, business rules in services, and
database access in repositories. This gives the project clear boundaries
without introducing a large framework.

### Request flow

```text
HTTP request
    ↓
Router and middleware
    ↓
Handler (controller)
    ↓
Service (business logic)
    ↓
Repository (database access)
    ↓
HTML response
    ├── Full page for a normal browser request
    └── Partial HTML fragment for an HTMX request
```

### Directory structure

```text
myapp/
├── cmd/
│   └── web/
│       └── main.go              # Application startup
├── internal/
│   ├── config/                  # Environment and application config
│   ├── server/                  # Server, routes, and middleware
│   ├── handlers/                # HTTP controllers
│   ├── services/                # Business logic
│   ├── repositories/            # SQL and database access
│   ├── models/                  # Domain/data models
│   └── views/                   # Template parsing and rendering
│       ├── layouts/
│       │   └── base.html
│       ├── pages/
│       │   └── movies.html
│       └── partials/
│           └── movie-row.html
├── migrations/
├── static/
│   ├── css/
│   └── js/                      # Optional; HTMX can be the main client behavior
└── go.mod
```

Keep handlers thin: parse input, call a service, choose a full page or partial
template, and write the response. Detect an HTMX request with
`r.Header.Get("HX-Request") == "true"`. Use `303 See Other` after successful
form submissions where a redirect is appropriate. Prefer server-rendered HTML
over JSON unless the application also needs a separate API.

### Chapter 2 todos

- [x] Rename the application entry point from `cmd/api` to `cmd/web`.
- [x] Create the basic web server and configure its port and environment.
- [x] Add the main application structure for shared configuration, logging, and dependencies.
- [x] Register the homepage route and a simple health route such as `/healthz`.
- [x] Replace API terminology with web application terminology throughout the project.
- [x] Add the initial template directories for layouts, pages, and HTMX partials.
- [x] Confirm that the server starts and both initial routes respond correctly.

### Chapter 3 todos

- [ ] Replace JSON response thinking with server-rendered HTML responses.
- [ ] Create a base layout for shared document structure and navigation.
- [ ] Add page templates for complete browser requests.
- [ ] Add partial templates for HTMX requests.
- [ ] Decide which routes return full pages and which return fragments.
- [ ] Add consistent rendering and error-handling helpers.
- [ ] Confirm that normal requests and HTMX requests receive the correct HTML.

### Chapter 4 todos

- [ ] Replace JSON request parsing with HTML form parsing.
- [ ] Define the form fields required by each page or feature.
- [ ] Add validation for required fields, formats, and allowed values.
- [ ] Create a consistent way to display validation errors in templates.
- [ ] Return form errors as page or HTMX partial responses where appropriate.
- [ ] Preserve submitted form values when validation fails.
- [ ] Add tests for valid submissions and common invalid inputs.

### Chapter 7 todos

- [ ] Define the domain models needed for the first CRUD feature.
- [ ] Move database queries into repository methods.
- [ ] Add service methods for creating, reading, updating, and deleting records.
- [ ] Add handlers for the CRUD pages and form submissions.
- [ ] Create full-page templates for listing and viewing records.
- [ ] Create HTMX partials for rows, forms, and status messages.
- [ ] Add validation and user-friendly success and error responses.

### Chapter 8 todos

- [ ] Separate reusable CRUD logic from feature-specific business rules.
- [ ] Add edit and delete interactions using HTMX where appropriate.
- [ ] Return only the affected partial after an HTMX update.
- [ ] Handle empty results, missing records, and database errors consistently.
- [ ] Add confirmation steps for destructive actions.
- [ ] Add tests for service behavior and important handler responses.

### Chapter 9 todos

- [ ] Define the filtering and sorting options for each list page.
- [ ] Add repository support for search, filters, sorting, and pagination.
- [ ] Parse and validate query parameters in the handler layer.
- [ ] Add pagination controls to the page templates.
- [ ] Make search and filter controls update the list with HTMX.
- [ ] Preserve active filters when moving between pages.
- [ ] Add an empty-state message when no records match the filters.

## Chapter-by-chapter adaptation guide

Based on the book's official table of contents:

**Legend:** **Keep** = mostly copy · **Adapt** = modify for server-rendered
HTML/HTMX · **Rewrite** = substantially change

| Chapter                       | Keep or adapt     | What changes                                                          |
| ----------------------------- | ----------------- | --------------------------------------------------------------------- |
| 2. Getting started            | **Adapt**         | MVC-style folders, routes, dependency wiring, and template setup      |
| 3. Sending JSON               | **Rewrite**       | Render full HTML pages and HTMX fragments instead                     |
| 4. Parsing JSON               | **Rewrite**       | Parse form data and validate inputs                                   |
| 5. Database setup             | **Keep**          | PostgreSQL and connection pooling remain applicable                   |
| 6. SQL migrations             | **Keep**          | Use the same migration approach                                       |
| 7. CRUD operations            | **Adapt**         | Keep models and SQL ideas; rewrite handlers and responses             |
| 8. Advanced CRUD              | **Adapt**         | Preserve repository logic and use HTMX partial updates                |
| 9. Filtering and pagination   | **Adapt**         | Add HTMX search, filtering, pagination, and table fragments           |
| 10. Rate limiting             | **Keep**          | Apply the same limits; tune them for browser interactions if needed   |
| 11. Graceful shutdown         | **Keep**          | No meaningful architectural change                                    |
| 12. User registration         | **Adapt**         | Use HTML forms, validation messages, and redirects/fragments          |
| 13. Sending emails            | **Keep**          | Email infrastructure remains applicable                               |
| 14. User activation           | **Adapt**         | Render HTML confirmation and error pages                              |
| 15. Authentication            | **Rewrite**       | Prefer secure cookie-based sessions over bearer tokens for the web UI |
| 16. Permissions               | **Adapt**         | Keep permission logic; return HTML errors or redirects                |
| 17. CORS                      | **Usually skip**  | Same-origin HTMX does not normally need CORS                          |
| 18. Metrics                   | **Keep**          | Use the same instrumentation principles                               |
| 19. Build and quality control | **Keep**          | Include templates, static assets, and template tests                  |
| 20. Deployment                | **Minor changes** | Serve HTML, CSS, and HTMX assets alongside the Go application         |
