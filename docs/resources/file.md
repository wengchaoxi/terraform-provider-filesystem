# filesystem_file (Resource)

## Example Usage

```hcl
resource "filesystem_file" "writer" {
  path    = "./helloworld.txt"
  content = "Hello World"
}
```

## Schema

### Required

- `path` (String) The path to save the file.

### Optional

- `content` (String) The content of the file.

### Read-Only

- `id` (String) The ID of this resource.
