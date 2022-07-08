# filesystem_file (Data Source)

## Example Usage

```hcl
resource "filesystem_file" "writer" {
  path    = "./helloworld.txt"
  content = "Hello World"
}

data "filesystem_file" "reader" {
  path = filesystem_file.writer.path
}
```

## Schema

### Required

- `path` (String) The path to save the file.

### Read-Only

- `content` (String) The content of the file.
- `filename` (String) The name of the file.
- `id` (String) The ID of this resource.
- `size` (String) The size of The file.
- `update_time` (String) The update time of the file.
