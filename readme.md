# Routehub.Client.HUB

- go install github.com/a-h/templ/cmd/templ@latest
- go get github.com/a-h/templ

# Client HUB

## Depolyment

### ComposeEase CEase

Creat another micro application to manage Deployments.

#### Features

Closed to the world. Only accessible by server.
GRPC Server to manage deployments.

- Has Composes. (Id,  Version, Yaml, Name, Description, CreatedAt, UpdatedAt)
- Has Deployments. (Id, ComposeJson, Version, Name, Description, CreatedAt, UpdatedAt)
- Has DeploymentsLogs. (Id, DeploymentId, Log, CreatedAt)
- Has HealthCheck. (Id, DeploymentId, Status, CreatedAt, UpdatedAt)

- Use Docker SDK to deploy compose with Traefik Labels
- Simple & fast & Reliable Storage for Composes, Deployments, DeploymentLogs, HealthCheck
- Simple Dashboard to manage Composes, Deployments, DeploymentLogs, HealthCheck

## Data Synchronization Strategies for Client Configuration Data

### Option 2: **Decentralized Data Store (Per-Client Database)**
In this approach, each client application has its own isolated data store (e.g., Redis, SQLite, or even Clickhouse) to store configuration data locally.

#### How It Works:
1. **Data Storage**: The configuration data is stored locally in each client’s isolated database (e.g., Redis or SQLite) as part of the client’s Docker stack.
2. **Data Syncing**: When the dashboard updates data via the GraphQL API, the API pushes the new data directly to the client’s local database (e.g., via an HTTP or gRPC call).
3. **Client Data Access**: The client app reads the configuration from its local database for fast access.

#### Pros:
- **Low Latency**: The client reads directly from a local database, ensuring faster access to configuration data.
- **Offline Capability**: The client app is more resilient, even if the API is temporarily unavailable.
- **Isolated Data**: Each client stack has its own independent data store, reducing the risk of cross-client issues.

#### Cons:
- **Data Synchronization Complexity**: Ensuring all clients are updated promptly with the latest changes adds complexity.
- **Redundancy**: Configuration data is replicated across all clients, leading to potential data duplication.

#### Best Practice:
- Use **webhooks** or **gRPC** from the GraphQL API to push updates to the client’s local database as soon as a change is made. Alternatively, you could have a small background service in each client stack that periodically pulls updates from the API and stores them locally.
- Ensure data consistency with versioning or timestamps to avoid conflicts.

### Final Considerations:
- **Consistency and Conflict Resolution**: Use timestamps or versioning in your data schema to manage potential update conflicts across clients.
- **Security and Access Control**: If exposing Redis or a similar service across clients, ensure proper access control, ideally with client-specific namespaces.
- **Scalability**: Consider the load on your centralized API if you have many clients polling or fetching updates frequently. Caching layers and selective polling intervals can help manage this.

These are a few best practice options. The final choice depends on your priority between real-time updates, performance, security, and architectural complexity. Let me know if you need more details on any of these approaches!