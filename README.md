<h1 align="center">Wells Go Framework</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
</p>

<p align="center">
  <strong>Wells Go</strong> is a lightweight scaffolding tool for generating modular Go backend projects with <b>Clean Architecture</b> principles.  
  Save time, enforce structure, and start coding your business logic right away!
</p>

<hr>

<h2>📌 Table of Contents</h2>
<ul>
  <li><a href="#about-the-project">About The Project</a></li>
  <li><a href="#features">Features</a></li>
  <li><a href="#project-structure">Generated Project Structure</a></li>
  <li><a href="#installation">Installation</a></li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#cli-commands">CLI Commands</a></li>
  <li><a href="#example-crud">Example CRUD</a></li>
  <li><a href="#testing">Testing</a></li>
  <li><a href="#development-workflow">Development Workflow</a></li>
  <li><a href="#contributing">Contributing</a></li>
  <li><a href="#license">License</a></li>
</ul>

<h2 id="about-the-project">ℹ️ About The Project</h2>
<p>
  <b>Wells Go</b> helps Go developers kickstart projects with a clean and modular foundation.  
  Instead of manually setting up boilerplate, Wells Go generates a ready-to-use structure following Domain-Driven Design (DDD) and Clean Architecture.
</p>

<h2 id="features">🚀 Features</h2>
<ul>
  <li>Generate Go project structure instantly</li>
  <li>Opinionated <b>Clean Architecture + DDD</b> folder layout</li>
  <li>Built-in support for DTOs, Mappers, Services, and Repositories</li>
  <li>Includes working example <b>CRUD API</b></li>
  <li>Test-ready scaffolding (unit + integration)</li>
  <li>CLI commands for code generation and runtime execution</li>
</ul>

<h2 id="project-structure">🗂️ Generated Project Structure</h2>
<pre>
myapp/
├── application/       # Usecases, services, mappers, DTOs
├── domain/            # Entities, repository interfaces
├── infrastructure/    # Database, Redis, external services
├── interfaces/http/   # HTTP handlers, routing
├── response/          # Response helpers (JSON, paging, error)
├── util/              # Utility functions
├── main.go            # Application entry point
└── go.mod             # Go modules
</pre>

<h2 id="installation">🛠️ Installation</h2>
<h3>Prerequisites</h3>
<ul>
  <li>Go 1.18 or higher</li>
  <li>Git</li>
</ul>

<h3>Install Wells CLI</h3>
<pre><code>go install github.com/welliardiansyah/wells-cli@latest
</code></pre>
<p>Make sure your <code>$GOPATH/bin</code> is in your <code>PATH</code>:</p>
<pre><code>export PATH=$PATH:$(go env GOPATH)/bin
</code></pre>

<h2 id="usage">⚡ Usage</h2>
<p>To scaffold a new Go project:</p>
<pre><code>wells new myapp
cd myapp
go mod tidy
wells run
</code></pre>

<h2 id="cli-commands">💻 CLI Commands</h2>
<h3>Code Generation Commands</h3>
<ul>
  <li><code>wells model NAME</code> → Creates a model/entity</li>
  <li><code>wells migration-postgres NAME</code> → Creates a PostgreSQL migration file</li>
  <li><code>wells inbound-http-fiber NAME</code> → Generates HTTP handlers using Fiber</li>
  <li><code>wells inbound-message-rabbitmq NAME</code> → Creates RabbitMQ message consumers</li>
  <li><code>wells inbound-command NAME</code> → Creates CLI command handlers</li>
  <li><code>wells outbound-database-postgres NAME</code> → Creates PostgreSQL database adapter</li>
  <li><code>wells outbound-http NAME</code> → Creates HTTP client adapter</li>
  <li><code>wells outbound-message-rabbitmq NAME</code> → Creates RabbitMQ message adapter</li>
  <li><code>wells outbound-cache-redis NAME</code> → Creates Redis cache adapter</li>
  <li><code>wells generate-mocks</code> → Generates mock implementations from go:generate directives</li>
</ul>

<h3>Runtime Commands</h3>
<ul>
  <li><code>wells run</code> → Runs the project (HTTP server by default)</li>
  <li><code>wells http</code> → Runs project in HTTP server mode</li>
  <li><code>wells message SUBSCRIBER</code> → Runs message consumer</li>
  <li><code>wells command CMD NAME</code> → Executes specific CLI command</li>
</ul>

<h2 id="example-crud">📝 Example CRUD</h2>
<p>Wells Go generates a working CRUD for <code>User</code>:</p>
<ul>
  <li><code>GET /users</code> → List users</li>
  <li><code>POST /users</code> → Create new user</li>
  <li><code>GET /users/:id</code> → Get user by ID</li>
  <li><code>PUT /users/:id</code> → Update user</li>
  <li><code>DELETE /users/:id</code> → Delete user</li>
</ul>

<p>Example request:</p>
<pre><code>curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
</code></pre>

<h2 id="testing">🧪 Testing</h2>
<p>Run all tests:</p>
<pre><code>go test ./...
</code></pre>
<p>Generate coverage report:</p>
<pre><code>go test -coverprofile=coverage.profile -cover ./...
go tool cover -html=coverage.profile -o coverage.html
</code></pre>

<h2 id="development-workflow">🛠 Development Workflow</h2>
<p>Step-by-step workflow with Wells CLI:</p>
<ol>
  <li><strong>Create Models:</strong> <code>wells model product</code></li>
  <li><strong>Create Database Migrations:</strong> <code>wells migration-postgres product</code></li>
  <li><strong>Implement Domain Logic:</strong> Add business logic in <code>domain/</code> and <code>application/usecases/</code></li>
  <li><strong>Create Inbound Adapters:</strong>
    <ul>
      <li><code>wells inbound-http-fiber product</code></li>
      <li><code>wells inbound-message-rabbitmq product</code></li>
      <li><code>wells inbound-command product</code></li>
    </ul>
  </li>
  <li><strong>Create Outbound Adapters:</strong>
    <ul>
      <li><code>wells outbound-database-postgres product</code></li>
      <li><code>wells outbound-http product_api</code></li>
      <li><code>wells outbound-cache-redis product</code></li>
      <li><code>wells outbound-message-rabbitmq product</code></li>
    </ul>
  </li>
  <li><strong>Write Tests and Generate Mocks:</strong> <code>wells generate-mocks</code></li>
  <li><strong>Run the Application:</strong>
    <ul>
      <li><code>wells http</code> → Run HTTP server</li>
      <li><code>wells message consumer</code> → Run message consumer</li>
      <li><code>wells command CMD name</code> → Run specific command</li>
    </ul>
  </li>
</ol>

<h2 id="contributing">🤝 Contributing</h2>
<p>Contributions are welcome! Fork this repo, create a feature branch, and submit a pull request.</p>

<h2 id="license">📄 License</h2>
<p>MIT License - see <a href="https://github.com/welliardiansyah/wells-cli/blob/master/LICENSE.md">LICENSE</a> for details.</p>