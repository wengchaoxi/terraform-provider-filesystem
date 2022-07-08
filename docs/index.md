# filesystem Provider

This provider is used to manage filesystem.

## Example Usage

Do not keep your `secret_key` in HCL for production environments, use Terraform environment variables.

```hcl
provider "filesystem" {
  secret_key = "h3110_w0r1d"
}
```

## Schema

### Optional

- `secret_key` (String, Sensitive) The secret key will be used to encrypt the file.
