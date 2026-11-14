# Personal SSG in Go

## IDEA
- NOW should list what I am doing now, but older entries list what I did in the past. I think this gives a cool overview of what I did over the years
## Project overview

I want to build a small static site generator for my own website. It will turn Markdown files in a GitHub repository into a complete static website.

This is a personal tool, not a general-purpose framework. There is no need for plugins, a configuration language, theme support, or compatibility with every possible Markdown feature. I should optimize for the way I write and publish content.

The basic flow is:

```text
Markdown and assets in GitHub
            |
            v
       Personal SSG
  scan -> parse -> render
            |
            v
          dist/
            |
            v
       GitHub Pages
```

## Design ideas
Design I like the most:
[[constellation-pisces-monochrome-mockup.png]]

### Mockups

#### Earlier directions

[[butterfly-inspired-blog-mockup.png]]
[[typographic-grid-blog-mockup.png]]
[[experimental-color-grid-blog-mockup.png]]

#### Archive and editorial directions

[[living-index-blog-mockup.png]]
[[annotated-reading-room-mockup.png]]
[[fast-minimal-blog-mockup.png]]

#### Experimental layout directions

[[constellation-index-mockup.png]]
[[terminal-ledger-mockup.png]]
[[editorial-strip-blog-mockup.png]]
[[workbench-blog-mockup.png]]

#### Constellation variations

[[constellation-light-index-mockup.png]]
[[constellation-orbit-mockup.png]]
[[constellation-axis-mockup.png]]
[[constellation-timeline-mockup.png]]

#### Constellation index variations

These are small deviations from the original **Constellation Index** direction. The structure stays the same while the palette, node shapes, spacing, and supporting details change.

[[constellation-variation-01-cream-cobalt.png]]
[[constellation-variation-02-sage-coral.png]]
[[constellation-variation-03-lavender-blue.png]]
[[constellation-variation-04-peach-navy.png]]
[[constellation-variation-05-offset-nodes.png]]
[[constellation-variation-06-square-nodes.png]]
[[constellation-variation-07-archive-rail.png]]
[[constellation-variation-08-dotted-field.png]]
[[constellation-variation-09-centered-reading.png]]
[[constellation-variation-10-color-legend.png]]

#### Constellation and grid experiments

These variations change the underlying graph and grid arrangement rather than only changing colors or node shapes.

[[constellation-variation-11-triangle-grid.png]]
[[constellation-variation-12-lattice.png]]
[[constellation-variation-13-clusters.png]]
[[constellation-variation-14-grid-rail.png]]
[[constellation-variation-15-radial-grid.png]]
[[constellation-variation-16-diagonal-grid.png]]

#### Pisces-inspired variations

April 15 falls under **Aries in Western astrology**, while the Sun is physically in the **Pisces constellation** around that date. Since the site's visual identity is inspired by being a fish, these variations use an abstract Pisces structure: two connected chains of nodes joined by a thin cord. The daughter's date can be represented by a small, separate accent node rather than a literal zodiac illustration.

[[constellation-pisces-paper-mockup.png]]
[[constellation-pisces-monochrome-mockup.png]]
[[constellation-pisces-dark-mockup.png]]
[[constellation-aries-dark-mockup.png]]

The current design goal is a fast, mostly static constellation-style site: typography, CSS Grid, Flexbox, thin rules, nodes, dots, and small inline SVG details should do most of the work. Avoid large image assets, heavy animation libraries, canvas effects, and unnecessary JavaScript.

The base constellation experience should work without JavaScript. CSS can handle the layout, nodes, connecting rules, borders, colors, responsive stacking, and hover states. JavaScript is optional and should only add enhancements such as cursor trails, filtering, animated connections, or client-side search.

## Implementation guide

The CSS and layout concepts for building the constellation design live in [[constellation CSS implementation.md]]. It covers semantic HTML, CSS Grid, Flexbox, nodes, connection paths, inline SVG, responsive layouts, themes, accessibility, and the recommended learning order.

## Design and template layer

The visual design is part of this website, not a separate generic theme or framework. The mockups are references for the site's presentation layer:

```text
Markdown content
      |
      v
Go content and page data
      |
      v
HTML templates + CSS
      |
      v
Generated static pages
```

The generator should own the content pipeline:

- Scan and parse Markdown.
- Build typed page and collection data.
- Resolve URLs and page types.
- Pass data to templates.
- Write the final HTML.

The templates and stylesheet should own the design:

- Typography and spacing.
- CSS Grid and Flexbox layouts.
- Navigation and page structure.
- Post, project, and collection presentation.
- Responsive behavior.
- Small inline SVG or CSS details.

There is no need for theme switching or a generic theme API. I only need one design that makes my own writing enjoyable. The initial design direction should combine the **Constellation Index** and **Fast Minimal** mockups: mostly text, thin rules, a small accent palette, nodes and connections, and unusual structure without heavy assets.

## Hosting and publishing

GitHub Pages is a good fit:

- It hosts static HTML, CSS, JavaScript, and images.
- GitHub Actions can build and deploy the site whenever I push to `main`.
- It is free for a public repository, making it a very cheap option for a personal blog.
- A custom domain can be added later if I want one.

The repository remains the source of truth. I write Markdown, commit it, and push it. GitHub Actions runs the Go generator and publishes the generated `dist/` directory to GitHub Pages.

The first deployment should use GitHub Pages' **GitHub Actions** source rather than committing generated files to the repository. The workflow will:

1. Check out the repository.
2. Install or use the required Go version.
3. Run `go run ./cmd/ssg` (or a future `go build` binary).
4. Upload `dist/` as the Pages artifact.
5. Deploy that artifact to GitHub Pages.

## Repository structure

```text
personal-website/
├── content/
│   ├── index.md
│   ├── about.md
│   ├── blog/
│   │   ├── first-post.md
│   │   └── second-post.md
│   └── projects/
│       ├── project-one.md
│       └── project-two.md
├── layouts/
│   ├── base.html
│   ├── home.html
│   ├── page.html
│   ├── post.html
│   ├── project.html
│   └── partials/
│   │   ├── header.html
│   │   ├── footer.html
│   │   └── metadata.html
├── public/
│   ├── css/
│   │   └── site.css
│   ├── js/              # Keep empty unless interaction is genuinely useful
│   └── images/          # Prefer no images; use inline SVG when practical
├── cmd/
│   └── ssg/
│       └── main.go
├── internal/
│   ├── content/
│   ├── markdown/
│   ├── render/
│   └── site/
├── .github/
│   └── workflows/
│       └── deploy.yml
├── go.mod
└── dist/                 # Generated locally; do not edit by hand
```

The exact package split can stay small at first. I should not create abstractions until the generator needs them.

## Content format

Each Markdown file can use simple YAML frontmatter:

```markdown
---
title: My first post
date: 2026-11-07
description: A short introduction to my blog.
type: post
draft: false
---

The post content goes here.
```

The fields should be limited to what the website actually uses. A sensible starting set is:

- `title`
- `date`
- `description`
- `type` (`page`, `post`, or `project`)
- `draft`

The URL should normally be derived from the file path. For example, `content/blog/first-post.md` becomes `/blog/first-post/`.

## First version: required functionality

Only build the functionality needed by this site:

1. Recursively scan `content/` for Markdown files.
2. Read and parse frontmatter.
3. Convert Markdown to HTML.
4. Render the HTML through Go `html/template` layouts.
5. Generate clean URLs and `index.html` files.
6. Generate collection pages for blog posts and projects.
7. Copy everything in `public/` to `dist/`.
8. Fail with a useful error when frontmatter or a template is invalid.

I should use existing libraries for Markdown and frontmatter rather than writing either parser from scratch. A reasonable Go stack is:

| Concern | Choice |
| --- | --- |
| Language | Go |
| CLI | A small `main.go` using the standard library |
| Markdown | Goldmark |
| Frontmatter | Goldmark frontmatter extension or a small YAML library |
| Templates | `html/template` |
| File handling | Go standard library |
| Hosting | GitHub Pages |
| Automation | GitHub Actions |

The renderer should pass typed data to templates rather than exposing an overly clever template language.

For example, a page can be represented as:

```go
type Page struct {
    Title       string
    Description string
    Date        time.Time
    Type        string
    URL         string
    HTML        template.HTML
}
```

The template decides how that data looks. The generator does not need to know whether a post is displayed as a row, a strip, a card, or part of an index grid.

## Build pipeline

```text
Load site settings
       |
Scan content/
       |
Read frontmatter and Markdown
       |
Convert Markdown to HTML
       |
Resolve URL and page type
       |
Render layouts
       |
Generate collection pages
       |
Copy public/ to dist/
       |
Write dist/
```

The local command should be:

```bash
go run ./cmd/ssg
```

It should produce the same output locally and in GitHub Actions.

## Implementation phases

### Phase 1: one page

- Initialize the Go module.
- Add the CLI entry point.
- Read one Markdown file.
- Parse its frontmatter.
- Convert its body to HTML.
- Render one template to `dist/index.html`.

### Phase 2: real pages and assets

- Walk the entire `content/` directory.
- Generate paths such as `about/index.html`.
- Add the base, page, post, and project layouts.
- Add shared header, footer, and metadata partials.
- Copy `public/` into `dist/`.
- Add the initial design system in `public/css/site.css`.
- Implement the selected Fast Minimal/Living Index/Editorial Strips direction.
- Make the layout responsive with CSS Grid and Flexbox.
- Keep the homepage JavaScript-free.

### Phase 3: collections

- Treat files under `content/blog/` as posts.
- Treat files under `content/projects/` as projects.
- Generate blog and project index pages.
- Sort posts by date.
- Ignore drafts in production builds.
- Generate the homepage index from recent writing, projects, and optional `now` content.

### Phase 4: deployment

- Add the GitHub Actions workflow.
- Build on pushes to `main`.
- Deploy the `dist/` artifact to GitHub Pages.
- Add a local preview workflow later if useful.

### Phase 5: only if needed

- RSS feed.
- `sitemap.xml`.
- Better metadata and canonical URLs.
- Draft/preview mode.
- Custom domain.

RSS, sitemap, and SEO should not delay the first working website.

## Definition of done

The first useful version is done when this command:

```bash
go run ./cmd/ssg
```

produces something like:

```text
dist/
├── index.html
├── about/
│   └── index.html
├── blog/
│   ├── index.html
│   ├── first-post/
│   │   └── index.html
│   └── second-post/
│       └── index.html
├── projects/
│   ├── index.html
│   └── project-one/
│       └── index.html
├── css/
└── images/
```

The complete workflow is successful when I can write a Markdown file, push it to GitHub, wait for the Action to finish, and see the updated page on my GitHub Pages site.

## Guiding principle

Build the smallest tool that makes publishing my own writing enjoyable. Keep the content, generator, templates, and deployment workflow in one repository. Avoid building a framework; spend the effort on the site and on learning Go.
