# Provider Upjet-for-DNSimple

> **Community Project** — This is an unofficial, community-maintained Crossplane provider. It is not affiliated with or endorsed by DNSimple Corporation.

`provider-upjet-for-dnsimple` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) code generation tools. It exposes XRM-conformant managed resources for the [DNSimple API](https://developer.dnsimple.com/), enabling you to manage DNSimple DNS infrastructure declaratively via Kubernetes.

This provider is generated from the official [DNSimple Terraform provider](https://github.com/dnsimple/terraform-provider-dnsimple) (MPL-2.0), which is the upstream source of truth for all resource schemas.

## Supported Resources

| Group                                               | Kind     | Description                                     |
| --------------------------------------------------- | -------- | ----------------------------------------------- |
| `zonerecord.upjet-for-dnsimple.crossplane.nvst.cloud` | `Record` | DNS zone record (A, AAAA, CNAME, MX, TXT, etc.) |

> **Note:** This provider currently covers **zone records**. Contributions to expose additional DNSimple resources (domains, contacts, email forwards, certificates, etc.) as Crossplane managed resources are welcome.

## Getting Started

### Prerequisites

- A running [Crossplane](https://docs.crossplane.io/latest/software/install/) installation (v1.14+)
- A DNSimple API token and account ID — obtain them from your [DNSimple account settings](https://dnsimple.com/user)

### Install the Provider

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-upjet-for-dnsimple
spec:
  package: ghcr.io/investifytech/provider-upjet-for-dnsimple:v0.1.0
```

### Configure Credentials

Create a Kubernetes secret with your DNSimple token and account ID:

```console
kubectl create secret generic dnsimple-creds \
  --from-literal=credentials='{"dnsimple_token":"<your-api-token>","dnsimple_account":"<your-account-id>"}' \
  -n crossplane-system
```

The credentials JSON supports the following fields:

| Field              | Required | Description                         |
| ------------------ | -------- | ----------------------------------- |
| `dnsimple_token`   | Yes      | DNSimple API access token           |
| `dnsimple_account` | Yes      | DNSimple account ID                 |

Then create a `ProviderConfig`:

```yaml
apiVersion: upjet-for-dnsimple.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: dnsimple
spec:
  credentials:
    source: Secret
    secretRef:
      name: dnsimple-creds
      namespace: crossplane-system
      key: credentials
```

### Example: Create a Zone Record

```yaml
apiVersion: zonerecord.upjet-for-dnsimple.crossplane.nvst.cloud/v1beta1
kind: Record
metadata:
  name: www-example-com
spec:
  forProvider:
    zoneName: example.com
    name: www
    type: A
    value: 192.0.2.10
    ttl: 3600
  providerConfigRef:
    name: dnsimple
```

## Developing

This provider is generated from the [DNSimple Terraform provider](https://github.com/dnsimple/terraform-provider-dnsimple) using [Upjet](https://github.com/crossplane/upjet). The upstream Terraform provider is the source of truth for resource schemas.

### Prerequisites

- Go 1.24+
- A running Kubernetes cluster (for `make run`)

### Run the code-generation pipeline

Regenerates all `zz_*` files from the upstream Terraform provider schema:

```console
go run cmd/generator/main.go "$PWD"
```

### Adding a new resource

1. Add or update the resource configuration under `config/`
2. Re-run the generator: `go run cmd/generator/main.go "$PWD"`
3. Verify the generated types under `apis/`

For a detailed walkthrough, see the [Upjet provider generation guide](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md).

### Run against a Kubernetes cluster

```console
make run
```

### Build, push, and install

```console
make all
```

### Build binary only

```console
make build
```

## License

This project is licensed under the **Apache License 2.0** — see the [LICENSE](LICENSE) file for details.

It incorporates resource schemas derived from the [DNSimple Terraform provider](https://github.com/dnsimple/terraform-provider-dnsimple), which is licensed under the **Mozilla Public License 2.0 (MPL-2.0)**. These two licenses are compatible: MPL-2.0 is a file-level copyleft license that does not propagate to the Larger Work. Attribution is provided in the [NOTICE](NOTICE) file.

This is a **community project** and is not officially affiliated with or endorsed by DNSimple Corporation.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/InvestifyTECH/provider-upjet-for-dnsimple/issues).
