---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nginxproxymanager_proxy_hosts Data Source - nginxproxymanager"
subcategory: ""
description: |-
  Proxy Hosts data source
---

# nginxproxymanager_proxy_hosts (Data Source)

Proxy Hosts data source

## Example Usage

```terraform
# Fetch all proxy hosts
data "nginxproxymanager_proxy_hosts" "all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `proxy_hosts` (Attributes List) (see [below for nested schema](#nestedatt--proxy_hosts))

<a id="nestedatt--proxy_hosts"></a>
### Nested Schema for `proxy_hosts`

Read-Only:

- `access_list_id` (Number) The ID of the access list used by the proxy host.
- `advanced_config` (String) The advanced configuration used by the proxy host.
- `allow_websocket_upgrade` (Boolean) Whether websocket upgrades are allowed for the proxy host.
- `block_exploits` (Boolean) Whether exploits are blocked for the proxy host.
- `caching_enabled` (Boolean) Whether caching is enabled for the proxy host.
- `certificate_id` (String) The ID of the certificate used by the proxy host.
- `created_on` (String) The date and time the proxy host was created.
- `domain_names` (List of String) The domain names associated with the proxy host.
- `enabled` (Boolean) Whether the proxy host is enabled.
- `forward_host` (String) The host used to forward requests to the proxy host.
- `forward_port` (Number) The port used to forward requests to the proxy host.
- `forward_scheme` (String) The scheme used to forward requests to the proxy host. Can be either `http` or `https`.
- `hsts_enabled` (Boolean) Whether HSTS is enabled for the proxy host.
- `hsts_subdomains` (Boolean) Whether HSTS is enabled for subdomains of the proxy host.
- `http2_support` (Boolean) Whether HTTP/2 is supported for the proxy host.
- `id` (Number) The ID of the proxy host.
- `locations` (Attributes List) The locations associated with the proxy host. (see [below for nested schema](#nestedatt--proxy_hosts--locations))
- `meta` (Map of String) The meta data associated with the proxy host.
- `modified_on` (String) The date and time the proxy host was last modified.
- `ssl_forced` (Boolean) Whether SSL is forced for the proxy host.

<a id="nestedatt--proxy_hosts--locations"></a>
### Nested Schema for `proxy_hosts.locations`

Read-Only:

- `advanced_config` (String) The advanced configuration used by the location.
- `forward_host` (String) The host used to forward requests to the location.
- `forward_port` (Number) The port used to forward requests to the location.
- `forward_scheme` (String) The scheme used to forward requests to the location. Can be either `http` or `https`.
- `path` (String) The path associated with the location.

