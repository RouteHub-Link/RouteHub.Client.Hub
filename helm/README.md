# How to deploy the hub?

## Create a namespace

```bash
kubectl create namespace routehub-clients
```

## Create a secret for the registry

```bash
kubectl create secret docker-registry your-secret-name \
  --namespace routehub-clients \
  --docker-server=your.registry \
  --docker-username=your-username \
  --docker-password=your-password
```

## Deploy the hub

```bash
helm install your-hub-name ./helm \
  --namespace routehub-clients \
  --create-namespace \
  --set global.customRegistry=your.registry/ \
  --set global.imagePullSecrets[0].name=your-secret-name \
  --set routehubClientRest.environment.Name="Test Hub" \
  --set routehubClientRest.environment.ORGANIZATION_ID=test-org \
  --set routehubClientRest.environment.OWNER_ID=test-owner \
  --set routehubClientRest.environment.PLATFORM_ID=12058bdf-8940-43b3-bd90-13487e4c8fc4 \
  --set routehubClientRest.environment.SEED=TRUE 
```

## For Debug you could add the following flags

```bash
--dry-run --debug
```
